package cond

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/prepare"
	"github.com/system-pclub/GCatch/GFix/dispatcher/primitive/locker"
	"github.com/system-pclub/GCatch/GFix/dispatcher/search"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

type Cond struct {
	Name string
	Make_inst ssa.Instruction
	Pkg string

	Signs []*Cond_op
	Broadc []*Cond_op
	Waits []*Cond_op
	All_op []*Cond_op

	Status string
}

const Sign int = 0
const BroadC int = 1
const Wait int = 2

const Edited = "Edited"

type Cond_op struct {
	Name string
	Inst ssa.Instruction
	Type int

	Locks []*locker.Lock_op // This field is specially designed for C6A
	Wrappers []*Wrapper     // This field is specially designed for C6A

	Whole_line string
	Status string
	Parent *Cond
}

var Cond_makes []*Cond_make // A global variable used to record all make lines of a conditional variable

type Cond_make struct {
	Whole_line string
	Locker string
	Inst ssa.Instruction
	Pkg string
}

// Wrapper records a function that "contains" a cond_op. "Contains" means the function directly uses this cond_op, or
// its callee (or callee's callee) uses this cond_op. The maximum layer is C6A_call_chain_layer_for_chan_wrapper
type Wrapper struct {
	Fn *ssa.Function // When compare two Wrapper, can't directly compare fn or inst, because pointer will change during each compilation
	Fn_str string
	Inst ssa.Instruction
	Callee *Wrapper // if callee is nil, then inst is the chan_op itself, else inst is calling to another Wrapper
	Op *Cond_op// the wrapped operation
}

const C_Wait = "C_Wait"
const C_Signal = "C_Signal"
const C_Broadcast = "C_Broadcast"

func Scan_cond_inst_return_value_comment(inst ssa.Instruction) (v ssa.Value, comment string) {
	v = nil
	comment = ""
	if inst.Parent().Pkg == nil {
		return
	}

	inst_call,ok := inst.(ssa.CallInstruction)
	if !ok {
		return
	}

	switch {
	case sync_check.Is_cond_wait(inst):
		args := inst_call.Common().Args
		if len(args) != 1 {
			fmt.Println("Warning: a Cond wait op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.Cond" {
			fmt.Println("Warning: a Cond wait op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],C_Wait
	case sync_check.Is_cond_signal(inst):
		args := inst_call.Common().Args
		if len(args) != 1 {
			fmt.Println("Warning: a Cond signal op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.Cond" {
			fmt.Println("Warning: a Cond signal op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],C_Signal
	case sync_check.Is_cond_broadcast(inst):
		args := inst_call.Common().Args
		if len(args) != 1 {
			fmt.Println("Warning: a Cond broadcast op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.Cond" {
			fmt.Println("Warning: a Cond broadcast op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],C_Broadcast
	default:
		return
	}
}

// Scan_inst_and_record_to_slice take the original slice of *Cond and an inst as inputs, and returns updated slice of *Cond
// The update will happen when inst is Signal/Broadcast/Wait of a sync.Cond
func Scan_inst_and_record_to_slice(all_conds []*Cond, inst ssa.Instruction) (result []*Cond) {
	result = all_conds

	if inst.Parent().Pkg == nil {
		return
	}

	flag_make := sync_check.Is_cond_make(inst)
	if flag_make {
		new_make := &Cond_make{
			Whole_line: search.Read_inst_line(inst),
			Inst:       inst,
			Pkg:		inst.Parent().Pkg.Pkg.Path(),
		}
		locker_str := new_make.Whole_line
		index_NewCond := strings.Index(locker_str,"NewCond(")
		if index_NewCond < 0 {
			fmt.Println("Error in Scan_inst_and_record_to_slice: A conditional variable is made without using NewCond")
			return
		}
		locker_str = locker_str[index_NewCond+8:]
		index_right_p := strings.Index(locker_str,")")
		if index_right_p < 0 {
			return
		}
		locker_str = locker_str[:index_right_p]
		new_make.Locker = locker_str
		for _,cond_make := range Cond_makes {
			if cond_make.Whole_line == new_make.Whole_line {
				return
			}
		}
		Cond_makes = append(Cond_makes,new_make)
	}

	//p := (global.Prog.Fset).Position(inst.Pos())
	flag_sign, flag_broadc, flag_wait := sync_check.Is_cond_signal(inst), sync_check.Is_cond_broadcast(inst), sync_check.Is_cond_wait(inst)

	if flag_sign == false && flag_broadc == false && flag_wait == false {
		return
	}

	cond_name,whole_line := search.Cond_name_and_line(inst)
	if cond_name == "" {
		return
	}

	// find an existing cond in result, or create a new one if not found
	var edit_cond *Cond
	for _,cond := range result {

		if cond.Name == cond_name && cond.Pkg == inst.Parent().Pkg.Pkg.Path() {
			edit_cond = cond
			edit_cond.Status = Edited
			break
		}

	}
	if edit_cond == nil {
		edit_cond = &Cond{
			Name:      cond_name,
			Make_inst: nil,
			Pkg:       inst.Parent().Pkg.Pkg.Path(),
			Signs:     nil,
			Broadc:    nil,
			Waits:     nil,
			Status:    "",
		}
		result = append(result,edit_cond)
	}

	new_op := &Cond_op{
		Name:       cond_name,
		Inst:       inst,
		Type:       -1,
		Locks:      nil,
		Wrappers:   nil,
		Whole_line: whole_line,
		Status:     "",
		Parent:     edit_cond,
	}

	//write to field Wrappers
	parent_wrapper := &Wrapper{
		Fn:     inst.Parent(),
		Fn_str: inst.Parent().String(),
		Inst:   inst,
		Callee: nil,
		Op:     new_op,
	}
	if flag_sign && whole_line == "\tq.cond.Signal()\n" {
		flag_our_sign = 999
	}
	wrappers := recursive_list_wrappers(parent_wrapper,0,[]*Wrapper{parent_wrapper})
	new_op.Wrappers = wrappers

	//write to field Type, and append this new_op to edit_cond
	switch true {
	case flag_sign:
		new_op.Type = Sign
		edit_cond.Signs = append(edit_cond.Signs,new_op)
	case flag_broadc:
		new_op.Type = BroadC
		edit_cond.Broadc = append(edit_cond.Broadc,new_op)
	case flag_wait:
		new_op.Type = Wait
		edit_cond.Waits = append(edit_cond.Waits,new_op)
	}
	edit_cond.All_op = append(edit_cond.All_op,new_op)


	return
}

var flag_our_sign int = -1

func recursive_list_wrappers(this_wrapper *Wrapper, layer int, old_wrappers []*Wrapper) (result []*Wrapper) {
	result = old_wrappers

	//if this_wrapper.Fn.Name() == "worker$1" && strings.Contains(this_wrapper.Fn.String(),"ResourceQuotaController") && flag_our_sign == 999 {
	//	s := this_wrapper.Fn.String()
	//	this_wrapper.Fn.String()
	//	_ = s
	//}

	if layer > global.C6A_call_chain_layer_for_chan_wrapper {
		return
	}

	if prepare.Is_path_include(this_wrapper.Fn.String()) == false {
		return
	}

	this_node,ok := global.Call_graph.Nodes[this_wrapper.Fn]
	if !ok {
		return
	}
	for _,caller_edge := range this_node.In {
		if _,is_go := caller_edge.Site.(*ssa.Go);is_go { // meaning the caller is in another goroutine, can't record this caller
			continue
		}
		new_wrapper := &Wrapper{
			Fn:     caller_edge.Caller.Func,
			Fn_str: caller_edge.Caller.Func.String(),
			Inst:   caller_edge.Site,
			Callee: this_wrapper,
			Op:     this_wrapper.Op,
		}
		result = append(result,new_wrapper)
		result = recursive_list_wrappers(new_wrapper,layer + 1,result)
	}

	return
}

func Append_wrappers_if_not_in(old []*Wrapper, add *Wrapper) (result []*Wrapper) {
	result = old
	if is_wrapper_in_slice(add,result) == false {
		result = append(result,add)
	}
	return
}

func is_wrapper_in_slice(w *Wrapper,s []*Wrapper) bool {
	for _,old := range s {
		if w == old {
			return true
		}
	}
	return false
}
