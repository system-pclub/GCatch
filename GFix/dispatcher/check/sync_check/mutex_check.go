package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
	"strings"
)

func case_insensitive_equal(s1, s2 string) bool {
	s1, s2 = strings.ToUpper(s1), strings.ToUpper(s2)
	return s1 == s2
}

func case_insensitive_contains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func Is_mutex_make(inst ssa.Instruction) bool {

	inst_asAlloc,ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if inst_v_type.String() == "sync.Mutex" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func Is_mutex_lock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isMutex_Lock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}


	if call != nil {
		call_full_name := CallName(call)
		if call_full_name == "(*sync.Mutex).Lock" {
			flag_isMutex_Lock = true
		}
		if call.IsInvoke() == true {
			call_part_name := call.Method.Name()
			if case_insensitive_equal(call_part_name,"lock") && case_insensitive_equal(call_part_name,"unlock") == false { //It is impossible to precisely determine if this is a lock or not
				flag_isMutex_Lock = true
			}
		}
	}


	return flag_isMutex_Lock
}

func Is_mutex_unlock(inst ssa.Instruction) bool {
	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isMutex_Unlock := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		call_full_name := CallName(call)
		if call_full_name == "(*sync.Mutex).Unlock" {
			flag_isMutex_Unlock = true
		}
		if call.IsInvoke() == true {
			call_part_name := call.Method.Name()
			if case_insensitive_equal(call_part_name,"unlock") { //It is impossible to precisely determine if this is a lock or not
				flag_isMutex_Unlock = true
			}
		}
	}

	return flag_isMutex_Unlock
}

func Is_mutex(inst ssa.Instruction) bool {

	return Is_mutex_lock(inst) || Is_mutex_unlock(inst)
}

