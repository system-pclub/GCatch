package conflictinglock

import (
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/instinfo"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"github.com/system-pclub/gochecker/util"
	"strings"
)

const Unknown = "Unknown"
const Edited = "Edited"


type StLockingOp struct {
	StrName string
	I ssa.Instruction
	IsRWLock bool
	IsDefer bool

	Parent * StMutex

	StrFileName string
	NumLine int
}

type StUnlockingOp struct {
	StrName string
	I ssa.Instruction
	IsRWLock bool
	IsDefer bool

	Parent * StMutex
}


type StMutex struct{
	StrName string
	StrType string
	StrBastStructType string
	MapLockingOps map[ssa.Instruction] *StLockingOp
	MapUnlockingOps map[ssa.Instruction] *StUnlockingOp
	Pkg string // Don't use *ssa.Package here! It's not reliable



	StrStatus string
}

type StLockPair struct {
	PLock1 * StLockingOp
	PLock2 * StLockingOp
	CallChainID int
}

type stMutexPair struct {
	PMutex1 * StMutex
	PMutex2 * StMutex

	VecLockingPair [] * StLockPair
}

func getLockingOpInfo(inputInst ssa.Instruction) (strName string, strMutexType string, strOpType string, isDefer bool, isLockingOp bool) {

	strName = Unknown
	strMutexType = Unknown
	strOpType = Unknown
	isDefer = false
	isLockingOp = true

	var pCall * ssa.CallCommon

	switch pType := inputInst.(type) {
	case *ssa.Call:
		pCall = pType.Common()
	case *ssa.Defer:
		pCall = pType.Common()
		isDefer = true
	default :
		isLockingOp = false
		return
	}

	switch instinfo.GetCallName(pCall) {
	case "(*sync.Mutex).Lock":
		strMutexType = "Mutex"
		strOpType = "Lock"

	case "(*sync.Mutex).Unlock":
		strMutexType = "Mutex"
		strOpType = "Unlock"

	case "(*sync.RWMutex).Lock":
		strMutexType = "RWMutex"
		strOpType = "Lock"

	case "(*sync.RWMutex).Unlock":
		strMutexType = "RWMutex"
		strOpType = "Unlock"

	case "(*sync.RWMutex).RLock":
		strMutexType = "RWMutex"
		strOpType = "RLock"

	case "(*sync.RWMutex).RUnlock":
		strMutexType = "RWMutex"
		strOpType = "RUnlock"

	default:
		var strCallName string = ""
		if pCall.IsInvoke() {
			strCallName = pCall.Method.Name()
		} else {
			if callee, ok := pCall.Value.(*ssa.Function); ok {
				strCallName = callee.Name()
			}
		}
		if strCallName != "" {
			strCallName = strings.ToLower(strCallName)
			switch {
			case strCallName == "lock":
				strOpType = "Lock"
			case strCallName == "unlock":
				strOpType = "Unlock"
			case strCallName == "rLock":
				strOpType = "RLock"
			case strCallName == "runlock":
				strOpType = "RUnlock"
			default:
				isLockingOp = false
				return
			}
		} else {
			isLockingOp = false
			return
		}
	}

	strName = instinfo.GetMutexName(inputInst)
	if strName == "" {
		strName = Unknown
	}

	return
}


func handleInst(inputInst ssa.Instruction) (isLockingOp bool) {
	isLockingOp = false

	if inputInst.Parent().Pkg == nil {
		return
	}

	strName, strMutexType, strOpType, isDefer, isLockingOp := getLockingOpInfo(inputInst)

	if !isLockingOp {
		return
	}

	if strName == Unknown || strOpType == Unknown {
		return
	}

	var pMutex * StMutex
	strBaseType := ""

	if pCall, ok := inputInst.(*ssa.Call); ok {
		if len(pCall.Common().Args) == 0 {
			strBaseType = pCall.Common().Signature().Recv().Type().String() //util.GetBaseType(pcall.Common().Signature().Recv())
		} else {
			strBaseType = util.GetBaseType(pCall.Common().Args[0]).String()
		}
	}

	if pm, ok := MapMutex[inputInst.Parent().Pkg.Pkg.Path() + ": " + strName + " (" + strBaseType + ")"]; ok {
		pMutex = pm
	} else {
		pMutex = &StMutex{
			StrName:    		strName,
			StrType:    		strMutexType,
			StrBastStructType:	strBaseType,
			MapLockingOps:		map[ssa.Instruction] *StLockingOp {},
			MapUnlockingOps:	map[ssa.Instruction] *StUnlockingOp {},
			Pkg:				inputInst.Parent().Pkg.Pkg.Path(),
			StrStatus:  		Edited,
		}

		MapMutex[inputInst.Parent().Pkg.Pkg.Path() + ": " + strName + " (" + strBaseType + ")"] = pMutex
	}

	isRWLocking := false

	if strOpType == "RLock" || strOpType == "RUnlock" {
		isRWLocking = true
	}

	if strOpType == "Lock" || strOpType == "RLock" {

		loc := config.Prog.Fset.Position(inputInst.Pos())

		newLocking := &StLockingOp{
			StrName:		strName,
			I:				inputInst,
			IsRWLock:		isRWLocking,
			IsDefer:		isDefer,
			Parent:			pMutex,
			StrFileName: 	loc.Filename,
			NumLine: 		loc.Line,
		}
		pMutex.MapLockingOps[inputInst] = newLocking
	} else if strOpType == "Unlock" || strOpType == "RUnlock" {
		newUnlocking := & StUnlockingOp {
			StrName:	strName,
			I:			inputInst,
			IsRWLock:	isRWLocking,
			IsDefer: 	isDefer,
			Parent:		pMutex,
		}
		pMutex.MapUnlockingOps[inputInst] = newUnlocking
	}

	return
}

func handleFN(fn *ssa.Function) bool {
	hasLockingOp := false
	for _, bb := range fn.Blocks {
		for _, ii := range bb.Instrs {
			isLockingOp := handleInst(ii)
			if isLockingOp {
				hasLockingOp = true
			}
		}
	}
	return hasLockingOp
}