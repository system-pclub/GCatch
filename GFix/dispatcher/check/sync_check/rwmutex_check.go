package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
)

func Is_rwmutex_make(inst ssa.Instruction) bool {
	inst_asAlloc,ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.RWMutex" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}


func Is_rwmutex_lock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isrwmutex_lock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.RWMutex).Lock" {
			flag_isrwmutex_lock = true
		}
	}

	return flag_isrwmutex_lock
}

func Is_rwmutex_unlock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isrwmutex_unlock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.RWMutex).Unlock" {
			flag_isrwmutex_unlock = true
		}
	}

	return flag_isrwmutex_unlock
}

func Is_rwmutex_rlock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isrwmutex_rlock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.RWMutex).RLock" {
			flag_isrwmutex_rlock = true
		}
		if call.IsInvoke() == true {
			str_call_name := call.Method.Name()
			if case_insensitive_equal(str_call_name,"rlock") { //It is impossible to precisely determine if this is a lock or not
				flag_isrwmutex_rlock = true
			}
		}
	}

	return flag_isrwmutex_rlock
}

func Is_rwmutex_runlock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isrwmutex_runlock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if callName == "(*sync.RWMutex).RUnlock" {
			flag_isrwmutex_runlock = true
		}
		if call.IsInvoke() == true {
			str_call_name := call.Method.Name()
			if case_insensitive_equal(str_call_name,"runlock") { //It is impossible to precisely determine if this is a lock or not
				flag_isrwmutex_runlock = true
			}

		}
	}

	return flag_isrwmutex_runlock
}

func Is_rwmutex(inst ssa.Instruction) bool {

	return Is_rwmutex_lock(inst) || Is_rwmutex_unlock(inst) || Is_rwmutex_rlock(inst) || Is_rwmutex_runlock(inst)
}
