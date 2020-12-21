package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa/ssautil"
)

var C7A_reported []bug_report
var Unsorted_chan_values []ssa.Value

func Initialize() {
	C7A_reported = []bug_report{}
	Unsorted_chan_values = []ssa.Value{}
}

func findMakeChanByLineNo(program *ssa.Program, filename string, makelineno int) (*ssa.MakeChan, string) {
	var makeinst *ssa.MakeChan
	makeinst = nil
	var name string
	name = ""
	fset := program.Fset
	for fn := range ssautil.AllFunctions(program) {
		for _, bb := range fn.Blocks {
			for _, ins := range bb.Instrs {
				position := fset.Position(ins.Pos())
				if position.Filename == filename {
					if position.Line == makelineno {
						fmt.Print("[DEBUG] found makechan line no: ", ins)
						fmt.Printf(" %p in function %s %p\n", ins, fn.Name(), fn)
						//Use the first found function is good enough. Example: ethereum 10 and 11
						switch typedins := ins.(type) {
						case *ssa.MakeChan:
							makeinst = typedins
						case *ssa.Alloc:
							name = typedins.Comment
						}
					}
				}
			}
		}
	}
	return makeinst, name
}

func isTheSameReceiver(fn1 *ssa.Function, fn2 *ssa.Function) bool {
	fmt.Printf("[DEBUG] check if %p == %p or %p == %p \n", *fn1, *fn2, fn1.Signature.Recv(), fn2.Signature.Recv())
	return fn1 == fn2 || fn1.Signature.Recv() == fn2.Signature.Recv()
}

func findChanOpsByLineNo(program *ssa.Program, filename string, oplineno int, receiverFunc *ssa.Function) (ssa.Instruction, int) {
	println("[DEBUG] findChanOpsByLineNo")
	var opinst ssa.Instruction
	var gltype int
	_ = gltype
	fset := program.Fset
	for fn := range ssautil.AllFunctions(program) {
		for _, bb := range fn.Blocks {
			for _, ins := range bb.Instrs {
				position := fset.Position(ins.Pos())
				if position.Filename == filename {
					if position.Line == oplineno && receiverFunc != nil {//&& isTheSameReceiver(receiverFunc, ins.Parent()) {
						//println()
						fmt.Print("[DEBUG] found op line no: ", ins)
						fmt.Printf(" %p in function %s %p\n", ins, fn.Name(), fn)
						if sync_check.Is_chan_close(ins) || sync_check.Is_receive_to_channel(ins) {
							gltype = 2
							opinst = ins
						} else if sync_check.Is_send_to_channel(ins) {
							opinst = ins
							gltype = 1
						}
					}
				}
			}
		}
	}
	return opinst, gltype
}

func findSend(ch channel, oplineno int, path string) *ch_send {
	for _, send := range ch.sends {
		ins := send.inst
		fset := ins.Parent().Prog.Fset
		position := fset.Position(ins.Pos())
		if position.Filename == path {
			if position.Line == oplineno {
				//println()
				fmt.Println("[DEBUG] found send line no: ", ins)
				return send
				//fmt.Printf(" %p in function %s %p\n", ins)
			}
		}
	}
	inst, gltype := findChanOpsByLineNo(global.Prog, path, oplineno, ch.make_inst.Parent()) //TODO: sometimes it returns a close or send
	if gltype == 1 {
		ret := ch_send{
			name:             "(found by line number from ssa)",
			inst:             inst,
			is_in_case:       false,
			is_case_blocking: false,
			thread: goroutine{
				go_inst:   nil,
				head_node: nil,
			},
			whole_line:               "",
			path:                     nil,
			bb_path:                  nil,
			conds:                    nil,
			smt:                      nil,
			foundByLineNumberFromSSA: true,
		}
		return &ret
	} else {
		return nil
	}
}

func findRecv(ch channel, oplineno int, path string) *ch_receive {
	for _, recv := range ch.receives {
		ins := recv.inst
		fset := ins.Parent().Prog.Fset
		position := fset.Position(ins.Pos())
		printSSAByBB(ins.Block())
		fmt.Println("[DEBUG] ", position.Filename, position.Line)
		if position.Filename == path {
			if position.Line == oplineno {
				//println()
				fmt.Println("[DEBUG] found recv line no: ", ins)
				_, ok := ins.(*ssa.Select)
				if ok {
					println("[DEBUG] It is in a select.")
				}
				return recv
				/*				if !ok {
									return recv
								} else {
									println("[DEBUG] It is in a select.")
									inst, gltype := findChanOpsByLineNo(global.Prog, path, oplineno) //TODO: sometimes it returns a close or send
									if gltype == 2 {
										recv.inst = inst
										return recv
									}
								}*/
				//fmt.Printf(" %p in function %s %p\n", ins)
			}
		}
	}
	fmt.Println("[DEBUG] ch.make_inst.Parent() == ", ch.make_inst.Parent())
	inst, gltype := findChanOpsByLineNo(global.Prog, path, oplineno, ch.make_inst.Parent()) //TODO: sometimes it returns a close or send
	if gltype == 2 {
		ret := ch_receive{
			name:             "",
			inst:             inst,
			is_in_case:       false,
			is_case_blocking: false,
			thread:           goroutine{},
			whole_line:       "",
			path:             nil,
			bb_path:          nil,
			conds:            nil,
			smt:              nil,
		}
		return &ret
	} else {
		return nil
	}

}

func Start(path string, makelineno int, oplineno int) {
	//makeinst, name, gltype, opinst := findChanOpsByLineNo(global.Prog, path, makelineno, oplineno)
	makeinst, name := findMakeChanByLineNo(global.Prog, path, makelineno)
	//fmt.Println("[DEBUG] gltype=", gltype)
	//fmt.Println("[DEBUG] opinst=", opinst)
	fmt.Println("[DEBUG] name=", name)
	if makeinst == nil {
		panic("could not found the make instruction.")
	}
	ch := channel{}
	ch.make_inst = makeinst
	ch.name_make = name
	// Find call-graph starting from the function that makes the channel
	ch.prepare_call_graph(global.C7A_max_layer)

	// Store its send, receive, close, but ignore bb_path,conds,smt
	ch.store_SRC(global.C7A_max_call_chain_length)

	// Store all goroutines shown up in send, receive, close
	ch.fill_threads()

	// Store all escape usages of this channel, excluding MakeClosure of ch.threads
	ch.fill_escape(global.C7A_max_call_chain_length)

	// Check for directly use unbuffered channel without create child goroutine
	ch.check_early_usage()

	//See if this channel is only used in this function or goroutines created in this function
	/*if len(ch.escapes) > 0 {
		println("[DEBUG] len(ch.escapes) > 0")
		return
	}*/

	// Calculate bb_path,conds
	ch.fill_SRC_bb_path_cond_whole_line()
	println("[DISPATCH] ", dispatch(path, ch, oplineno))
}

func dispatch(path string, ch channel, oplineno int) int {
	send := findSend(ch, oplineno, path)
	fmt.Println("[DEBUG] send=", send)
	if send != nil {
		if isGL1_new(send, &ch) {
			return 1
		} else {
			println("[DEBUG] GL1 dispatch failed!")
			makeChanLineNo, rewriteLineNos := getGL3PatchLineNo(send, nil, &ch)
			if makeChanLineNo > 0 {
				fmt.Println("[PATCH] ", makeChanLineNo, rewriteLineNos)
				return 3
			}
			return 0
		}
	}
	recv := findRecv(ch, oplineno, path)
	fmt.Println("[DEBUG] recv=", recv)
	if recv != nil {
		insLineNo, delLineNos := getGL2PatchLineNoNew(recv, &ch)
		_ = delLineNos
		if insLineNo != -1 {
			fmt.Println("[PATCH] ", insLineNo, delLineNos)
			return 2
		} else {
			return 0
		}
	}
	return 0
}
