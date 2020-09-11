package instinfo

import (
	"github.com/system-pclub/GCatch/tools/go/ssa"
	"go/token"
	"go/types"
	"strconv"
)

func GetCallName(call *ssa.CallCommon) string {

	if call.IsInvoke() {
		return call.String()
	}

	switch v := call.Value.(type) {
	case *ssa.Function:
		fn, ok := v.Object().(*types.Func)
		if !ok {
			return ""
		}
		return fn.FullName()
	case *ssa.Builtin:
		return v.Name()
	}
	return ""
}

type StOpValue struct {
	Inst ssa.Instruction
	Value ssa.Value
	Comment string
}

const Lock, Unlock = "Lock", "Unlock"
const RWMutex, Mutex = "RWMutex", "Mutex"
const Send, Recv, Close, MakeChan = "Send", "Recv", "Close", "MakeChan"

func ScanInstFindLockerValue(inst ssa.Instruction) (v ssa.Value, comment string) {
	v = nil
	comment = ""

	if inst.Parent().Pkg == nil {
		return
	}
	instCall,ok := inst.(ssa.CallInstruction)
	if !ok {
		return
	}

	args := instCall.Common().Args
	switch {
	case IsMutexLock(inst) || IsRwmutexLock(inst):
		if instCall.Common().IsInvoke() == false {
			if len(args) == 0 {
				return
			}
			return args[0], Lock
		} else {
			return instCall.Common().Value, Lock
		}
	case IsMutexUnlock(inst) || IsRwmutexUnlock(inst):
		if instCall.Common().IsInvoke() == false {
			if len(args) == 0 {
				return
			}
			return args[0],Unlock
		} else {
			return instCall.Common().Value,Unlock
		}
	default:
		return
	}
}

// ScanInstFindChanValue return multiple values and strings, only when inst is Select
func ScanInstFindChanValue(inst ssa.Instruction) ([]ssa.Value,[]string) {
	if inst.Parent().Pkg == nil {
		return nil,nil
	}

	boolIsClose := IsChanClose(inst)

	var ssaValueOfPrimitive ssa.Value
	switch concrete := inst.(type) {
	case *ssa.Send:
		return []ssa.Value{concrete.Chan},[]string{Send}
	case *ssa.UnOp:
		if concrete.Op == token.ARROW {
			return []ssa.Value{concrete.X},[]string{Recv}
		}
	case ssa.CallInstruction:
		if boolIsClose {
			return concrete.Common().Args,[]string{Close}
		}
	case *ssa.MakeChan:
		ssaValueOfPrimitive = concrete
		return []ssa.Value{ssaValueOfPrimitive},[]string{MakeChan}
	case *ssa.Select:
		// Return one value and comment for each case in Select
		vecSsaValue := []ssa.Value{}
		vecComment := []string{}
		for i, state := range concrete.States {
			vecSsaValue = append(vecSsaValue, state.Chan)
			if concrete.Blocking {
				if state.Dir == types.SendOnly {
					vecComment = append(vecComment,"Blocking_Select_Send_"+strconv.Itoa(i))
				} else {
					vecComment = append(vecComment,"Blocking_Select_Recv_"+strconv.Itoa(i))
				}
			} else {
				if state.Dir == types.SendOnly {
					vecComment = append(vecComment,"Non_Blocking_Select_Send_"+strconv.Itoa(i))
				} else {
					vecComment = append(vecComment,"Non_Blocking_Select_Recv_"+strconv.Itoa(i))
				}
			}
		}
		return vecSsaValue, vecComment
	}

	return nil,nil
}

func IsChanClose(inst ssa.Instruction) bool {
	var call *ssa.CallCommon
	instCall, ok := inst.(*ssa.Call)



	if ok {
		call = instCall.Common()
	}

	deferIns, ok := inst.(*ssa.Defer)
	if ok {
		call = deferIns.Common()
	}

	if call != nil {
		if !call.IsInvoke() {
			callName := GetCallName(call)
			if callName == "close" {
				return true
			}
		}
	}

	return false
}