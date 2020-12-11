package instinfo

import (
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"go/types"
	"strings"
)

func IsRwmutexMake(inputInst ssa.Instruction) bool {
	instAlloc,ok := inputInst.(*ssa.Alloc)

	if ok {
		typeInst := instAlloc.Type().Underlying().(*types.Pointer).Elem()
		if typeInst.String() == "sync.RWMutex" {
			return true
		}
	}

	return false
}


func IsRwmutexLock(inputInst ssa.Instruction) bool {
	var fnCall *ssa.CallCommon

	instCall, ok := inputInst.(*ssa.Call)

	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)
	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		callName := GetCallName(fnCall)
		if callName == "(*sync.RWMutex).Lock" {
			return  true
		}
	}

	return false
}


func IsRwmutexUnlock(inputInst ssa.Instruction) bool {
	var fnCall *ssa.CallCommon

	instCall, ok := inputInst.(*ssa.Call)

	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)
	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		callName := GetCallName(fnCall)
		if callName == "(*sync.RWMutex).Unlock" {
			return true
		}
	}

	return false
}


func IsRwmutexRlock(inputInst ssa.Instruction) bool {
	var fnCall *ssa.CallCommon

	instCall, ok := inputInst.(*ssa.Call)

	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)
	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		callName := GetCallName(fnCall)
		if callName == "(*sync.RWMutex).RLock" {
			return true
		}
		if fnCall.IsInvoke() == true {
			strFnName := fnCall.Method.Name()

			if strings.ToLower(strFnName) == "rlock" {
				return true
			}
		}
	}

	return false
}

func IsRwmutexRunlock(inputInst ssa.Instruction) bool {
	var fnCall *ssa.CallCommon

	instCall, ok := inputInst.(*ssa.Call)

	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)
	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		callName := GetCallName(fnCall)
		if callName == "(*sync.RWMutex).RUnlock" {
			return true
		}
		if fnCall.IsInvoke() == true {
			strFnName := fnCall.Method.Name()

			if strings.ToLower(strFnName) == "runlock" {
				return true
			}

		}
	}

	return false
}

func IsRwmutex(inst ssa.Instruction) bool {

	return IsRwmutexLock(inst) || IsRwmutexUnlock(inst) || IsRwmutexRlock(inst) || IsRwmutexRunlock(inst)
}
