package sync_check

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/types"
	"strings"
)

func Type_atomic(inst ssa.Instruction) string {
	switch true {

	case Is_atomic_value_load(inst):
		return "atomic_value_load"

	case Is_atomic_value_store(inst):
		return "atomic_value_store"

	case Is_atomic_addint(inst):
		return "atomic_addint"

	case Is_atomic_loadint(inst):
		return "atomic_loadint"

	case Is_atomic_storeint(inst):
		return "atomic_storeint"

	case Is_atomic_swapint(inst):
		return "atomic_swapint"

	case Is_atomic_compareandswapint(inst):
		return "atomic_compareandswapint"

	case Is_atomic_loadpointer(inst):
		return "atomic_loadpointer"

	case Is_atomic_storepointer(inst):
		return "atomic_storepointer"

	case Is_atomic_swappointer(inst):
		return "atomic_swappointer"

	case Is_atomic_compareandswappointer(inst):
		return "atomic_compareandswappointer"

	}

	return "other"


}

func Is_atomic(inst ssa.Instruction) bool { // This function recognize all function calls whose name includes "atomic." as atomic operations
											// So we need further diagnosis by Type_atomic()
	if Is_atomic_value_make(inst) {
		return true
	}

	var call *ssa.CallCommon

	call_, ok := inst.(*ssa.Call)

	flag_isAtomic := false

	if ok {
		call = call_.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		callName := CallName(call)
		if strings.Contains(callName, "atomic.") {
			flag_isAtomic = true
		}
	}

	return flag_isAtomic
}

func Is_atomic_value_make(inst ssa.Instruction) bool {

	local_flag := false

	inst_asAlloc, ok := inst.(*ssa.Alloc)
	if ok {
		inst_v_type := inst_asAlloc.Type().Underlying().(*types.Pointer).Elem()
		if strings.Contains(inst_v_type.String(), "sync/atomic.Value")  {
			local_flag = true
		}
	}

	return local_flag
}

func Is_atomic_value_load(inst ssa.Instruction) bool {
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
		if callName == "(*sync/atomic.Value).Load" {
			local_flag = true
		}
	}

	return local_flag
}

func Is_atomic_value_store(inst ssa.Instruction) bool {
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
		if callName == "(*sync/atomic.Value).Store" {
			local_flag = true
		}
	}

	return local_flag
}

func Is_atomic_addint(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.AddInt32" || callName == "sync/atomic.AddInt64" || callName == "sync/atomic.AddUint32" || callName == "sync/atomic.AddUint64" || callName == "sync/atomic.AddUintptr"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_loadint(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.LoadInt32" || callName == "sync/atomic.LoadInt64" || callName == "sync/atomic.LoadUint32" || callName == "sync/atomic.LoadUint64" || callName == "sync/atomic.LoadUintptr"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_storeint(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.StoreInt32" || callName == "sync/atomic.StoreInt64" || callName == "sync/atomic.StoreUint32" || callName == "sync/atomic.StoreUint64" || callName == "sync/atomic.StoreUintptr"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_swapint(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.SwapInt32" || callName == "sync/atomic.SwapInt64" || callName == "sync/atomic.SwapUint32" || callName == "sync/atomic.SwapUint64" || callName == "sync/atomic.SwapUintptr"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_compareandswapint(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.CompareAndSwapInt32" || callName == "sync/atomic.CompareAndSwapInt64" || callName == "sync/atomic.CompareAndSwapUint32" || callName == "sync/atomic.CompareAndSwapUint64" || callName == "sync/atomic.CompareAndSwapUintptr"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_loadpointer(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.LoadPointer"  {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_storepointer(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.StorePointer" {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_swappointer(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.SwapPointer" {
			local_flag = true
		}
	}


	return local_flag
}

func Is_atomic_compareandswappointer(inst ssa.Instruction) bool {
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
		if callName == "sync/atomic.CompareAndSwapPointer"  {
			local_flag = true
		}
	}

	return local_flag
}