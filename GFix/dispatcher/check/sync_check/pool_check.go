package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)

func Is_pool_make(inst ssa.Instruction) bool {
	inst_asAlloc,ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.Pool" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Is_pool_get(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_ispool_get := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Pool).Get" {
			flag_ispool_get = true
		}
	}

	return flag_ispool_get
}

func Is_pool_put(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_ispool_put := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Pool).Put" {
			flag_ispool_put = true
		}
	}

	return flag_ispool_put
}

func Is_pool(inst ssa.Instruction) bool {
	return Is_pool_get(inst) || Is_pool_put(inst)
}
