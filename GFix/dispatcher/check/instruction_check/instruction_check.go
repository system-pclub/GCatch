package instruction_check

import (

	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

func Type_inst(inst ssa.Instruction) string {

	if  _,ok := inst.(*ssa.Alloc); ok {
		return "Alloc"
	}

	if  _,ok := inst.(*ssa.BinOp); ok {
		return "BinOp"
	}

	if  _,ok := inst.(*ssa.Call); ok {
		return "Call"
	}

	if  _,ok := inst.(*ssa.Defer); ok {
		return "Defer"
	}

	if  _,ok := inst.(*ssa.Go); ok {
		return "Go"
	}

	if  _,ok := inst.(*ssa.Store); ok {
		return "Store"
	}

	if  _,ok := inst.(*ssa.Send); ok {
		return "Send"
	}

	if  _,ok := inst.(*ssa.Select); ok {
		return "Select"
	}

	if  _,ok := inst.(*ssa.Return); ok {
		return "Return"
	}

	if  _,ok := inst.(*ssa.Phi); ok {
		return "Phi"
	}

	if  _,ok := inst.(*ssa.If); ok {
		return "If"
	}

	if  _,ok := inst.(*ssa.Jump); ok {
		return "Jump"
	}

	if _,ok := inst.(*ssa.FieldAddr); ok {
		return "FieldAddr"
	}

	return "Unknown"
}
