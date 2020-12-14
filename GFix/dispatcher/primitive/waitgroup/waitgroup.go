package waitgroup

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

type WG struct {
	Name string
	Make WG_Make
	Adds []WG_Add
	Dones []WG_Done
	Waits []WG_Wait
	Child Goroutine

	Status string
}

type WG_Make struct {
	Name string
	Inst ssa.Instruction
	Value ssa.Value

	Parent *WG
}

type WG_Add struct {
	Name string
	Inst ssa.Instruction
	Is_size_const bool
	Size int

	Parent *WG
}

type WG_Done struct {
	Name string
	Inst ssa.Instruction

	Parent *WG
}

type WG_Wait struct {
	Name string
	Inst ssa.Instruction

	Parent *WG
}

type Goroutine struct {
	Fn_str string
	Fn_may_nil *ssa.Function
	Parent_str string
	Parent_may_nil *ssa.Function

	Creation_site ssa.Instruction

}

const W_Wait = "W_Wait"
const W_Add = "W_Add"
const W_Done = "W_Done"

func Scan_WG_inst_return_value_comment(inst ssa.Instruction) (v ssa.Value, comment string) {
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
	case sync_check.Is_waitgroup_wait(inst):
		args := inst_call.Common().Args
		if len(args) != 1 {
			fmt.Println("Warning: a WaitGroup wait op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.WaitGroup" {
			fmt.Println("Warning: a WaitGroup wait op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],W_Wait
	case sync_check.Is_waitgroup_add(inst):
		args := inst_call.Common().Args
		if len(args) != 2 { //Note: add has 2 arguments
			fmt.Println("Warning: a WaitGroup add op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.WaitGroup" {
			fmt.Println("Warning: a WaitGroup add op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],W_Add
	case sync_check.Is_waitgroup_done(inst):
		args := inst_call.Common().Args
		if len(args) != 1 {
			fmt.Println("Warning: a WaitGroup done op has",len(args),"arguments")
			return
		}
		if args[0].Type().String() != "*sync.WaitGroup" {
			fmt.Println("Warning: a WaitGroup done op's argument has type:",args[0].Type().String())
			return
		}
		return args[0],W_Done
	default:
		return
	}
}
