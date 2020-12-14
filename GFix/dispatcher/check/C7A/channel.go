package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/constraint"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/mycallgraph"
	"github.com/system-pclub/GCatch/GFix/dispatcher/path"
	"github.com/system-pclub/GCatch/GFix/dispatcher/search"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"go/types"
	"strings"
)

type channel struct {
	make_inst   ssa.Instruction
	name_make   string
	sends       []*ch_send
	receives    []*ch_receive
	closes      []*ch_close
	escapes     []*ch_escapes
	buffer      int
	parent_node *mycallgraph.Call_node
	main_thread goroutine
	threads     []goroutine
}

type ch_send struct {
	name             string
	inst             ssa.Instruction // this can be *ssa.Send or *ssa.Select
	is_in_case       bool
	is_case_blocking bool
	thread           goroutine
	whole_line       string
	path             []*mycallgraph.Call_node // path represents the call-chain from channel's make func
	// to the parent function of the send inst. call-chain includes
	// head function and tail function
	bb_path                  []*ssa.BasicBlock
	conds                    []path.Cond
	smt                      *constraint.SMT_set
	foundByLineNumberFromSSA bool
}

type ch_receive struct {
	name             string
	inst             ssa.Instruction
	is_in_case       bool
	is_case_blocking bool
	thread           goroutine
	whole_line       string
	path             []*mycallgraph.Call_node
	bb_path          []*ssa.BasicBlock
	conds            []path.Cond
	smt              *constraint.SMT_set
}

type ch_close struct {
	name       string
	inst       ssa.Instruction // inst may be *ssa.Call or *ssa.Defer
	thread     goroutine
	whole_line string
	path       []*mycallgraph.Call_node
	bb_path    []*ssa.BasicBlock
	conds      []path.Cond
	smt        *constraint.SMT_set
}

type ch_escapes struct {
	name       string
	line       int
	inst       ssa.Instruction
	thread     goroutine
	whole_line string
	path       []*mycallgraph.Call_node
}

type goroutine struct {
	go_inst   *ssa.Go // it this is nil, this must be the main goroutine
	head_node *mycallgraph.Call_node
}

func list_local_unbuffer_chan(fn *ssa.Function) (result []*channel) {
	for _, bb := range fn.Blocks {
		for _, inst := range bb.Instrs {
			if sync_check.Is_make_channel(inst) {
				new_channel := channel{}
				new_channel.make_inst = inst
				inst_MakeChan, ok := inst.(*ssa.MakeChan)
				if !ok { // this is not possible
					continue
				}

				// store the buffer size
				bv := inst_MakeChan.Size
				bv_Const, ok := bv.(*ssa.Const)
				if !ok { // according to ssa.MakeChan, this is not possible
					continue
				}
				defer func() {
					if r := recover(); r != nil { // I am concerned that bv_Const.Int64() may panic
						fmt.Println("Recovered in func:", fn.String())
					}
				}()
				b_int := bv_Const.Int64()
				new_channel.buffer = int(b_int)
				if new_channel.buffer != 0 {
					continue
				}

				// store its name
				full_name, part_name := search.Chan_name_make(inst_MakeChan)
				if full_name != part_name || full_name == "" { // Meaning this channel is stored in a field, so it may be used somewhere else
					continue
				}
				new_channel.name_make = full_name

				result = append(result, &new_channel)
			}
		}
	}
	return
}

func (ch *channel) prepare_call_graph(max_layer int) {
	parent := ch.make_inst.Parent()

	// generate call-graph of parent function
	parent_node := mycallgraph.New_Call_node(parent, 0)
	mycallgraph.Initialize(global.C7A_max_recursive_count)
	parent_node.Fill_call_map_after_init(max_layer)

	ch.parent_node = parent_node

	main_goroutine := goroutine{
		go_inst:   nil,
		head_node: parent_node,
	}

	ch.main_thread = main_goroutine

}

// check_early_usage checks for a kind of easy and rare bugs: an unbuffered channel is directly used
// after creation, and there is no "go" between the usage
func (ch *channel) check_early_usage() {
	if ch.buffer != 0 {
		return
	}

	bb_make := ch.make_inst.Block()

	type bb_inst struct {
		bb      *ssa.BasicBlock
		inst    ssa.Instruction
		is_send bool // if this is false, then it is recv
		send    *ch_send
		recv    *ch_receive
	}

	//list all bbs where ch is used in the same function as bb_make, and not used in select
	BIs := []bb_inst{}
	for _, send := range ch.sends {
		bb := send.inst.Block()
		if bb.Parent().String() != bb_make.Parent().String() {
			continue
		}
		if send.is_in_case {
			continue
		}
		new_BI := bb_inst{
			bb:      bb,
			inst:    send.inst,
			is_send: true,
			send:    send,
			recv:    nil,
		}
		BIs = append(BIs, new_BI)
	}
	for _, recv := range ch.receives {
		bb := recv.inst.Block()
		if bb.Parent().String() != bb_make.Parent().String() {
			continue
		}
		if recv.is_in_case {
			continue
		}
		new_BI := bb_inst{
			bb:      bb,
			inst:    recv.inst,
			is_send: false,
			send:    nil,
			recv:    recv,
		}
		BIs = append(BIs, new_BI)
	}

	for _, BI := range BIs {
		paths, err := path.Find_paths_locally(bb_make, BI.bb)
		if err != nil {
			fmt.Println("Error in check_early_usage:", err.Error())
		}

		// Let's report a bug if all insts between inst_make and inst_use
		// are not creating a goroutine that is in ch.goroutines
		// and are not one of ch.escapes's inst

	path_loop:
		for _, a_path := range paths {
			insts_between := []ssa.Instruction{}

			if len(a_path) == 1 {
				// case 1: bb_make and bb_use is the same.
				index_make, index_use := -1, -1
				for i, inst := range bb_make.Instrs {
					if inst == ch.make_inst {
						index_make = i
					}
					if inst == BI.inst {
						index_use = i
					}
				}
				if index_make == -1 || index_use == -1 {
					continue
				}

				// For each inst between the make_inst and use_inst, see if it is creating a goroutine that is in ch.goroutines
				for i := index_make + 1; i < index_use; i++ {
					inst := bb_make.Instrs[i]
					insts_between = append(insts_between, inst)
				}

			} else {
				// case 2: bb_make and bb_use are different.

				// In bb_make, record insts after inst_make
				flag_found_make := false
				for _, inst := range bb_make.Instrs {
					if inst == ch.make_inst {
						flag_found_make = true
					}

					if flag_found_make {
						insts_between = append(insts_between, inst)
					}
				}

				// In bb_use, record insts before inst_use
				flag_found_use := false
				for _, inst := range BI.bb.Instrs {
					if inst == BI.inst {
						flag_found_use = true
					}

					if flag_found_use == false {
						insts_between = append(insts_between, inst)
					}
				}

				// In other bbs, record all insts
				for _, bb := range a_path {
					if bb == bb_make || bb == BI.bb {
						continue
					}

					for _, inst := range bb.Instrs {
						insts_between = append(insts_between, inst)
					}
				}

			}

			for _, inst_between := range insts_between {
				for _, thread := range ch.threads {
					if thread.go_inst == inst_between {
						continue path_loop
					}
				}

				for _, escape := range ch.escapes {
					if escape.inst == inst_between {
						continue path_loop
					}
				}
			}

			//if we reach here, meaning we find a path trigger the bug
			new_report := bug_report{
				ch:   ch,
				send: nil,
				recv: nil,
			}
			if BI.is_send {
				new_report.send = BI.send
			} else {
				new_report.recv = BI.recv
			}
			//report_direct_usage(new_report)
			return
		}
	}
}

// store ch.sends/receives/closes, by recursively scan ch.make_inst.Parent() and its callees
func (ch *channel) store_SRC(max_call_chain_length int) {
	current_path := []*mycallgraph.Call_node{ch.parent_node}
	ch.store_SRC_in_node_and_callees(ch.parent_node, current_path, ch.main_thread, max_call_chain_length)
}

func (ch *channel) store_SRC_in_node(node *mycallgraph.Call_node, current_path []*mycallgraph.Call_node, current_goroutine goroutine) {
	//current_path includes node

	for _, bb := range node.Fn.Blocks {
		for _, inst := range bb.Instrs {
			//p := (global.Prog.Fset).Position(inst.Pos())
			flag_send, flag_receive := sync_check.Is_send_to_channel(inst), sync_check.Is_receive_to_channel(inst)
			flag_close, flag_select := sync_check.Is_chan_close(inst), sync_check.Is_select_to_channel(inst)

			if flag_send == false && flag_receive == false && flag_close == false && flag_select == false {
				continue
			}

			if flag_select {
				inst_select, ok := inst.(*ssa.Select)
				if !ok { //This should never happen
					continue
				}
				for _, state := range inst_select.States {
					ch_name := search.Chan_name_value(state.Chan)
					if ch_name != ch.name_make {
						continue
					}

					if state.Dir == types.SendOnly {
						new_send := ch_send{
							name:             ch_name,
							inst:             inst_select,
							is_in_case:       true,
							is_case_blocking: inst_select.Blocking,
							thread:           current_goroutine,
							path:             current_path,
						}
						if is_ch_send_in_slice(&new_send, ch.sends) == false {
							ch.sends = append(ch.sends, &new_send)
						}
					} else if state.Dir == types.RecvOnly {
						new_receive := ch_receive{
							name:             ch_name,
							inst:             inst_select,
							is_in_case:       true,
							is_case_blocking: inst_select.Blocking,
							thread:           current_goroutine,
							path:             current_path,
						}
						if is_ch_receive_in_slice(&new_receive, ch.receives) == false {
							ch.receives = append(ch.receives, &new_receive)
						}
					}

				}
			} else if flag_send {
				//if ch.parent_node.Fn.Name() == "TestNonblockingDialWithEmptyBalancer" && node.Fn.Name() == "TestNonblockingDialWithEmptyBalancer$1" {
				//	ch.parent_node.Fn.Name()
				//}
				inst_Send, ok := inst.(*ssa.Send)
				if !ok { // this should never happen
					continue
				}
				ch_name := search.Chan_name_value(inst_Send.Chan)
				if ch_name != ch.name_make {
					continue
				}
				new_send := ch_send{
					name:             ch_name,
					inst:             inst,
					is_in_case:       false,
					is_case_blocking: false,
					thread:           current_goroutine,
					path:             current_path,
				}
				if is_ch_send_in_slice(&new_send, ch.sends) == false {
					ch.sends = append(ch.sends, &new_send)
				}
			} else if flag_receive {
				inst_UnOp, ok := inst.(*ssa.UnOp)
				if !ok { // this should never happen
					continue
				}
				if inst_UnOp.Op != token.ARROW {
					continue
				}
				ch_name := search.Chan_name_value(inst_UnOp.X)
				if ch_name != ch.name_make {
					continue
				}
				new_receive := ch_receive{
					name:             ch_name,
					inst:             inst,
					is_in_case:       false,
					is_case_blocking: false,
					thread:           current_goroutine,
					path:             current_path,
				}
				if is_ch_receive_in_slice(&new_receive, ch.receives) == false {
					ch.receives = append(ch.receives, &new_receive)
				}
			} else if flag_close {

				var call *ssa.CallCommon
				inst_Call, ok := inst.(*ssa.Call)
				if ok {
					call = inst_Call.Common()
				}

				deferIns, ok := inst.(*ssa.Defer)
				if ok {
					call = deferIns.Common()
				}

				if call == nil {
					continue
				}
				if call.IsInvoke() {
					continue
				}
				if len(call.Args) != 1 {
					continue
				}

				ch_name := search.Chan_name_value(call.Args[0])
				if ch_name != ch.name_make {
					continue
				}
				new_close := ch_close{
					name:   ch_name,
					inst:   inst,
					thread: current_goroutine,
					path:   current_path,
				}
				if is_ch_close_in_slice(&new_close, ch.closes) == false {
					ch.closes = append(ch.closes, &new_close)
				}
			}
		}
	}
}

func (ch *channel) store_SRC_in_node_and_callees(node *mycallgraph.Call_node, current_path []*mycallgraph.Call_node, current_goroutine goroutine, max_call_chain_length int) {

	if len(current_path) > max_call_chain_length {
		return
	}

	ch.store_SRC_in_node(node, current_path, current_goroutine)

	for call_inst, callees := range node.Call_map {
		for _, callee := range callees {
			new_path := append(current_path, callee)
			new_goroutine := current_goroutine
			if inst_go, is_go := call_inst.(*ssa.Go); is_go {
				new_goroutine = goroutine{
					go_inst:   inst_go,
					head_node: callee,
				}
			}
			ch.store_SRC_in_node_and_callees(callee, new_path, new_goroutine, max_call_chain_length)
		}
	}

}

type ignore_tuple struct {
	line int
	fn   string
}

func (ch *channel) list_ignore_tuples() []ignore_tuple {
	ignore_list := []ignore_tuple{}

	//make
	inst := ch.make_inst
	p := (global.Prog.Fset).Position(inst.Pos())
	t := ignore_tuple{
		line: p.Line,
		fn:   inst.Parent().String(),
	}
	ignore_list = append(ignore_list, t)

	//send
	for _, send := range ch.sends {
		inst = send.inst
		p := (global.Prog.Fset).Position(inst.Pos())
		t := ignore_tuple{
			line: p.Line,
			fn:   send.thread.head_node.Fn.String(),
		}
		ignore_list = append(ignore_list, t)
	}

	//receive
	for _, receive := range ch.receives {
		inst = receive.inst
		p := (global.Prog.Fset).Position(inst.Pos())
		t := ignore_tuple{
			line: p.Line,
			fn:   receive.thread.head_node.Fn.String(),
		}
		ignore_list = append(ignore_list, t)
	}

	//close
	for _, a_close := range ch.closes {
		inst = a_close.inst
		p := (global.Prog.Fset).Position(inst.Pos())
		t := ignore_tuple{
			line: p.Line,
			fn:   a_close.thread.head_node.Fn.String(),
		}
		ignore_list = append(ignore_list, t)
	}

	return ignore_list
}

// fill all other usages of the channel, by recursively scan ch.make_inst.Parent() and its callees
// ignoring insts that are of the same Line as any of ch.make_inst/send/receive/close
// ignoring MakeClosure that makes any goroutine in ch.threads
func (ch *channel) fill_escape(max_call_chain_length int) {
	ignore_list := ch.list_ignore_tuples()
	current_path := []*mycallgraph.Call_node{ch.parent_node}
	ch.fill_escape_in_node_and_callees(ch.parent_node, ignore_list, current_path, ch.main_thread, max_call_chain_length)
}

func (ch *channel) fill_escape_in_node(node *mycallgraph.Call_node, ignore_list []ignore_tuple, current_path []*mycallgraph.Call_node, current_goroutine goroutine) {
	//current_path includes node
	//scanned_values := []*ssa.Value{}
	for _, bb := range node.Fn.Blocks {
	Loop_inst:
		for _, inst := range bb.Instrs {
			// if inst is *ssa.MakeClosure, see if it uses ch as binding. If does, directly append new escape

			inst_MakeClosure, ok := inst.(*ssa.MakeClosure)
			if ok {
				for _, t := range ignore_list {
					if t.fn == inst_MakeClosure.Fn.String() {
						continue Loop_inst
					}
				}
				for _, binding := range inst_MakeClosure.Bindings {
					Alloc, ok := binding.(*ssa.Alloc)
					if ok {
						p := (global.Prog.Fset).Position(inst.Pos())
						t := ignore_tuple{
							line: p.Line,
							fn:   inst.Parent().String(),
						}
						if Alloc.Comment == ch.name_make {
							new_escape := ch_escapes{
								name:       ch.name_make,
								line:       t.line,
								inst:       inst,
								thread:     current_goroutine,
								whole_line: search.Read_inst_line(inst),
								path:       current_path,
							}
							if is_escape_in_slice(&new_escape, ch.escapes) == false {
								for _, thread := range ch.threads {
									if thread.head_node.Fn == inst_MakeClosure.Fn {
										continue Loop_inst
									}
								}
								ch.escapes = append(ch.escapes, &new_escape)
								continue Loop_inst
							}
						}
					}
				}

			}

			operand_list := inst.Operands([]*ssa.Value{})
			for _, operand := range operand_list {
				var Alloc *ssa.Alloc
				operand_as_Alloc, ok := (*operand).(*ssa.Alloc)
				if ok {
					Alloc = operand_as_Alloc
				} else {
					operand_as_UnOp, ok := (*operand).(*ssa.UnOp)
					if ok {
						if operand_as_UnOp.Op == token.MUL {
							Alloc, _ = operand_as_UnOp.X.(*ssa.Alloc)
						}
					}

				}
				if Alloc == nil {
					continue
				}

				if Alloc.Comment != ch.name_make {
					continue
				}

				// now we are sure that ch is used here, and this is not scanned before
				// we want to know if it is at the same line as any in ignore_list
				p := (global.Prog.Fset).Position(inst.Pos())
				t := ignore_tuple{
					line: p.Line,
					fn:   inst.Parent().String(),
				}
				if is_ignore_tuple_in_slice(t, ignore_list) == false {
					new_escape := ch_escapes{
						name:       ch.name_make,
						line:       t.line,
						inst:       inst,
						thread:     current_goroutine,
						whole_line: search.Read_inst_line(inst),
						path:       current_path,
					}
					if is_escape_in_slice(&new_escape, ch.escapes) == false {
						// TODO: the strategies here are very ad-hoc
						flag_skip_this_escape := false

						// if line == 0, this may be sythesized inst, ignore this escape;
						if new_escape.line == 0 {
							flag_skip_this_escape = true
						}

						// if the whole_line contains "case", ignore this escape
						// 	because send/receive.inst may be select, and the line of select is not the line of case
						if strings.Contains(new_escape.whole_line, "case") {
							flag_skip_this_escape = true
						}

						// if line number matches goroutine's ssa.Go inst, or the whole_line contains a substring matching any name of goroutines, ignore this escape
						goroutines := []goroutine{}
						for _, send := range ch.sends {
							goroutines = append(goroutines, send.thread)
						}
						for _, recv := range ch.receives {
							goroutines = append(goroutines, recv.thread)
						}
						for _, close := range ch.closes {
							goroutines = append(goroutines, close.thread)
						}
						for _, goroutine := range goroutines {
							if strings.Contains(new_escape.whole_line, goroutine.head_node.Fn.Name()) {
								flag_skip_this_escape = true
							}
							if goroutine.go_inst != nil {
								p := (global.Prog.Fset).Position(goroutine.go_inst.Pos())
								if new_escape.line == p.Line {
									flag_skip_this_escape = true
								}
							}
						}

						if flag_skip_this_escape {
							continue
						}

						ch.escapes = append(ch.escapes, &new_escape)
					}
				}
			}
		}
	}
}

func (ch *channel) fill_escape_in_node_and_callees(node *mycallgraph.Call_node, ignore_list []ignore_tuple, current_path []*mycallgraph.Call_node, current_goroutine goroutine, max_call_chain_length int) {

	if len(current_path) > max_call_chain_length {
		return
	}

	ch.fill_escape_in_node(node, ignore_list, current_path, current_goroutine)

	for call_inst, callees := range node.Call_map {
		for _, callee := range callees {
			new_path := append(current_path, callee)
			new_goroutine := current_goroutine
			if inst_go, is_go := call_inst.(*ssa.Go); is_go {
				new_goroutine = goroutine{
					go_inst:   inst_go,
					head_node: callee,
				}
			}
			ch.fill_escape_in_node_and_callees(callee, ignore_list, new_path, new_goroutine, max_call_chain_length)
		}
	}

}

func (ch *channel) fill_threads() {
	ch.threads = []goroutine{}
	for _, send := range ch.sends {
		if is_goroutine_in_slice(send.thread, ch.threads) == false {
			ch.threads = append(ch.threads, send.thread)
		}
	}
	for _, recv := range ch.receives {
		if is_goroutine_in_slice(recv.thread, ch.threads) == false {
			ch.threads = append(ch.threads, recv.thread)
		}
	}
	for _, a_close := range ch.closes {
		if is_goroutine_in_slice(a_close.thread, ch.threads) == false {
			ch.threads = append(ch.threads, a_close.thread)
		}
	}
}

func (ch *channel) fill_SRC_bb_path_cond_whole_line() {
	for _, send := range ch.sends {
		path_send, err := path.Find_path_by_call_chain(ch.make_inst.Block(), send.inst.Block(), send.path)
		if err != nil {
			send.name = "Unhealthy"
			if global.Print_err_log {
				fmt.Println("Error in C7A.fill_SRC_bb_path_cond:\t", err)
			}
			continue
		}
		send.bb_path = path.Delete_useless_bbs(path_send)
		// list the constraints of the path
		send.conds = path.List_cond_of_path(send.bb_path, send.inst.Block())
		send.whole_line = search.Read_inst_line(send.inst)
	}

	for _, recv := range ch.receives {
		path_recv, err := path.Find_path_by_call_chain(ch.make_inst.Block(), recv.inst.Block(), recv.path)
		if err != nil {
			recv.name = "Unhealthy"
			if global.Print_err_log {
				fmt.Println("Error in C7A.fill_SRC_bb_path_cond:\t", err)
			}
			continue
		}
		recv.bb_path = path.Delete_useless_bbs(path_recv)
		// list the constraints of the path
		recv.conds = path.List_cond_of_path(recv.bb_path, recv.inst.Block())
		recv.whole_line = search.Read_inst_line(recv.inst)
	}

	for _, a_close := range ch.closes {
		path_close, err := path.Find_path_by_call_chain(ch.make_inst.Block(), a_close.inst.Block(), a_close.path)
		if err != nil {
			a_close.name = "Unhealthy"
			if global.Print_err_log {
				fmt.Println("Error in C7A.fill_SRC_bb_path_cond:\t", err)
			}
			continue
		}
		a_close.bb_path = path.Delete_useless_bbs(path_close)
		// list the constraints of the path
		a_close.conds = path.List_cond_of_path(a_close.bb_path, a_close.inst.Block())
		a_close.whole_line = search.Read_inst_line(a_close.inst)
	}

	return
}

// for one send, see if its precondition can => definitely reach one receive in another goroutine
func (ch *channel) check_one_send_liveness(send *ch_send, recv_list []*ch_receive) {

	// use some strategies to exclude some easy cases that can directly report bug

	// strategy 1: no recv and send is not in select
	if len(recv_list) == 0 && send.is_in_case == false {
		new_bug := bug_report{
			ch:   ch,
			send: send,
			recv: nil,
		}
		report(new_bug)
		return
	}

	// strategy 2: all recvs are in select, but send is not
	flag_send_in_select := send.is_in_case
	flag_all_recv_in_select := true
	for _, recv := range recv_list {
		if recv.is_in_case == false {
			flag_all_recv_in_select = false
		}
	}
	if flag_send_in_select == false && flag_all_recv_in_select == true {
		new_bug := bug_report{
			ch:   ch,
			send: send,
			recv: nil,
		}
		report(new_bug)
		return
	}

	// strategy 3: number of recvs is 1, and this recv has condition, and send has no condition
	if len(recv_list) == 1 && len(recv_list[0].conds) > 0 && len(send.conds) == 0 {
		new_bug := bug_report{
			ch:   ch,
			send: send,
			recv: nil,
		}
		report(new_bug)
		return
	}

	// get the pre-condition of send
	// get the pre-conditions of recv
	// union all pre-conditions of recv
	// see if pre-condition of send => union of recv

}

// for one recv, see if its precondition can => definitely reach one send or close in another goroutine
func (ch *channel) check_one_recv_liveness(recv *ch_receive, send_list []*ch_send, close_list []*ch_close) {

	// use some strategies to exclude some easy cases that can directly report bug

	// strategy 1: no send and close, and recv is not in select
	if len(send_list)+len(close_list) == 0 && recv.is_in_case == false {
		new_bug := bug_report{
			ch:   ch,
			send: nil,
			recv: recv,
		}
		report(new_bug)
		return
	}

	// strategy 2: no close, all sends are in select, but recv is not
	flag_recv_in_select := recv.is_in_case
	flag_all_send_in_select := true
	for _, send := range send_list {
		if send.is_in_case == false {
			flag_all_send_in_select = false
		}
	}
	if flag_recv_in_select == false && flag_all_send_in_select == true && len(close_list) == 0 {
		new_bug := bug_report{
			ch:   ch,
			send: nil,
			recv: recv,
		}
		report(new_bug)
		return
	}

	// strategy 3: number of send is 1, number of close is 0, and this send has condition, and recv has no condition
	if len(send_list) == 1 && len(close_list) == 0 && len(send_list[0].conds) > 0 && len(recv.conds) == 0 {
		new_bug := bug_report{
			ch:   ch,
			send: nil,
			recv: recv,
		}
		report(new_bug)
		return
	}

	// strategy 4: number of send is 0, number of close is 1, and this close has condition, and recv has no condition and recv is not in select with default
	if len(send_list) == 0 && len(close_list) == 1 && len(close_list[0].conds) > 0 && len(recv.conds) == 0 && recv.is_in_case == false { //(recv.is_in_case == false || (recv.is_in_case && recv.is_case_blocking))

		new_bug := bug_report{
			ch:   ch,
			send: nil,
			recv: recv,
		}
		report(new_bug)
		return
	}

	// get the pre-condition of send
	// get the pre-conditions of recv
	// union all pre-conditions of recv
	// see if pre-condition of send => union of recv

}

func is_ptr_value_in_slice(v *ssa.Value, slice []*ssa.Value) bool {
	for _, v_s := range slice {
		if v_s == v {
			return true
		}
	}
	return false
}

func is_ignore_tuple_in_slice(t ignore_tuple, slice []ignore_tuple) bool {
	for _, t_s := range slice {
		if t_s.fn == t.fn && t_s.line == t.line {
			return true
		}
	}
	return false
}

func is_ch_close_in_slice(c *ch_close, slice []*ch_close) bool {
	for _, c_s := range slice {
		if c_s.inst == c.inst {
			return true
		}
	}
	return false
}

func is_ch_send_in_slice(c *ch_send, slice []*ch_send) bool {
	for _, c_s := range slice {
		if c_s.inst == c.inst {
			return true
		}
	}
	return false
}

func is_ch_receive_in_slice(c *ch_receive, slice []*ch_receive) bool {
	for _, c_s := range slice {
		if c_s.inst == c.inst {
			return true
		}
	}
	return false
}

func is_escape_in_slice(e *ch_escapes, slice []*ch_escapes) bool {
	for _, e_s := range slice {
		if e_s.line == e.line && e_s.whole_line == e.whole_line {
			return true
		}
	}
	return false
}

func is_goroutine_in_slice(g goroutine, slice []goroutine) bool {
	for _, g_s := range slice {
		if g_s.go_inst == g.go_inst && g_s.head_node == g.head_node {
			return true
		}
	}
	return false
}
