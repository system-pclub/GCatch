package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

func Is_cond_make(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iscond_make := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Cond).NewCond" {
			flag_iscond_make = true
		}
	}

	return flag_iscond_make
}

func Is_cond_broadcast(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iscond_broadcast := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Cond).Broadcast" {
			flag_iscond_broadcast = true
		}
	}

	return flag_iscond_broadcast
}

func Is_cond_signal(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iscond_signal := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Cond).Signal" {
			flag_iscond_signal = true
		}
	}

	return flag_iscond_signal
}

func Is_cond_wait(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_iscond_wait := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.Cond).Wait" {
			flag_iscond_wait = true
		}
	}

	return flag_iscond_wait
}

func Is_cond(inst ssa.Instruction) bool {
	return Is_cond_broadcast(inst) || Is_cond_signal(inst) || Is_cond_wait(inst)
}
