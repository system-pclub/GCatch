package instruction_check

import "github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"

func Is_defer(inst ssa.Instruction) bool {
	_,ok := inst.(*ssa.Defer)
	return ok
}

