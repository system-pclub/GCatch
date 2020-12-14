package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)

func Is_waitgroup_make(inst ssa.Instruction) bool {
	inst_asAlloc, ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.WaitGroup" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Is_waitgroup_add(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iswaitgroup_add := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.WaitGroup).Add" {
			flag_iswaitgroup_add = true
		}
	}

	return flag_iswaitgroup_add
}

func Is_waitgroup_done(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iswaitgroup_done := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.WaitGroup).Done" {
			flag_iswaitgroup_done = true
		}
	}

	return flag_iswaitgroup_done
}

func Is_waitgroup_wait(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iswaitgroup_wait := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.WaitGroup).Wait" {
			flag_iswaitgroup_wait = true
		}
	}

	return flag_iswaitgroup_wait
}

func Is_waitgroup(inst ssa.Instruction) bool {
	return Is_waitgroup_add(inst) || Is_waitgroup_done(inst) || Is_waitgroup_wait(inst)
}
