package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)

func Is_syncmap_make(inst ssa.Instruction) bool {
	inst_asAlloc,ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.Map" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Is_syncmap_delete(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_issyncmap_delete := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Map).Delete" {
			flag_issyncmap_delete = true
		}
	}

	return flag_issyncmap_delete
}

func Is_syncmap_load(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	local_flag := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Map).Load" {
			local_flag = true
		}
	}

	return local_flag
}

func Is_syncmap_loadorstore(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	local_flag := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Map).LoadOrStore" {
			local_flag = true
		}
	}

	return local_flag
}

func Is_syncmap_range(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_issyncmap_store := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Map).Range" {
			flag_issyncmap_store = true
		}
	}

	return flag_issyncmap_store
}

func Is_syncmap_store(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	local_flag := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Map).Store" {
			local_flag = true
		}
	}

	return local_flag
}
