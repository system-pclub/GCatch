package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)

func Is_once_make(inst ssa.Instruction) bool {
	inst_asAlloc,ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.Once" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Is_once_do(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isonce_do := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Once).Do" {
			flag_isonce_do = true
		}
	}

	return flag_isonce_do
}

func Is_once(inst ssa.Instruction) bool {
	return Is_once_do(inst)
}