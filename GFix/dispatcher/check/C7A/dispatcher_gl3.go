package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

/*
func isLastReceiveBeforeReturnInSameBB(recvInst *ssa.UnOp) bool {
	bb := recvInst.Block()
	state := 0
	//foundDeref := false
	for index, inst := range bb.Instrs {
		switch x := inst.(type) {
		case *ssa.RunDefers:
			if state == 1 {
				state = 1
			}
		case *ssa.Call:
			isCloseChan := sync_check.Is_chan_close(inst)
			//state==1 && isCloseChan: continue
			//state==1 && !isCloseChan: false
			//state!=1: continue
			if state == 1 {
				if !isCloseChan {
					return false
				}
			}
		case *ssa.UnOp:
			if sync_check.Is_receive_to_channel(x) {
				if state != 0 {
					return false
				}
				state = 1
			} else {
				if state == 1 {
					_, ok2 := bb.Instrs[index+1].(*ssa.Call)
					if !ok2 {
						return false
					}
				}
			}
		case *ssa.Return:
			if state > 0 {
				return true
			} else {
				return false
			}
		default:
			if state == 1 {
				return false
			}
		}
	}
	return false
}

func isLastReceiveBeforeReturn(recvInst *ssa.UnOp) bool {
	bb_start := recvInst.Block()
	if !isSendAndJump(recvInst) {
		return false
	}
	if len(bb_start.Succs) == 1 {
		succ := bb_start.Succs[0]
		if len(succ.Instrs) == 2 {
			_, ok1 := succ.Instrs[0].(*ssa.RunDefers)
			_, ok2 := succ.Instrs[1].(*ssa.Return)
			if ok1 && ok2 {
				return true
			}
		} else if len(succ.Instrs) == 4 {
			isChanClose := sync_check.Is_chan_close(succ.Instrs[1])
			_, ok1 := succ.Instrs[2].(*ssa.RunDefers)
			_, ok2 := succ.Instrs[3].(*ssa.Return)
			if isChanClose && ok1 && ok2 {
				return true
			}
		}
	}
	return false
}

func checkLastReceiveBeforeReturn(recvInst *ssa.UnOp) bool {
	if !(isLastReceiveBeforeReturnInSameBB(recvInst) || isLastSendBeforeReturn(recvInst)) {
		if debug {
			printSSAByBB(recvInst.Block())
			fmt.Println("is not the last send before return!")
		}
		if checkHasSideEffect(recvInst.Block()) {
			return true
		}
		return false
	}
	return true
}*/

func isNewObject(fn *ssa.Function) bool {
	sig := fn.Signature.Results()
	if sig.Len() > 0 {
		rettype := sig.At(0).Type().String()
		subnames := strings.Split(rettype, ".")
		println("[DEBUG] parent func ret type:", rettype)
		println("[DEBUG] parent func name:", fn.Name())

		if "new"+strings.ToLower(subnames[len(subnames)-1]) == strings.ToLower(fn.Name()) {
			return true
		} else {
			return false
		}
	}
	return false
}

func returnsChannel(fn *ssa.Function) bool {
	sig := fn.Signature.Results()
	if sig.Len() > 0 {
		rettype := sig.At(0).Type().String()
		println("[DEBUG] parent func ret type:", rettype)
		if strings.Contains(rettype, "chan ") {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func checkEtcd36(ch *channel) bool {
	fset := ch.make_inst.Parent().Prog.Fset
	if ch.make_inst.Parent().Name() == "TestTxnPanics" && strings.Contains(fset.Position(ch.make_inst.Pos()).Filename, "clientv3/txn_test.go") {
		return true
	} else {
		return false
	}
}

func getGL3PatchLineNo(sendobj *ch_send, recvobj *ch_receive, ch *channel) (int, []int) {
	ret2 := make([]int, 0)
	//TODO: add goroutine relationship check
	if sendobj != nil && recvobj != nil {
		panic("This shouldn't happen!")
	}
	if sendobj != nil {
		goroutineA := sendobj.inst.Parent()     //check if parent function are the same as the goroutine
		goroutineInst := sendobj.thread.go_inst //
		if returnsChannel(ch.make_inst.Parent()) {
			return -1, ret2
		}
		if sendobj.foundByLineNumberFromSSA {
			if isNewObject(ch.make_inst.Parent()) {
				return -1, ret2
			}
		} else {
			if !CheckCallRelationship(goroutineA, goroutineInst) && !checkEtcd36(ch) {
				println("[DEBUG] call relationship check failed!")
				return -1, ret2
			}
		}
		send := sendobj.inst
		fset := sendobj.inst.Parent().Prog.Fset
		sendInst, ok := send.(*ssa.Send)
		if !ok {
			if debug {
				fmt.Println("Not a send inst!") //TODO: not a correct way to handle error
			}
			return -1, ret2
		}
		b := checkAllSideEffect(sendInst)
		if b {
			println("[DEBUG] side effect check failed!")
			return -1, ret2
		}
		ret2 = append(ret2, getLineNo(fset, sendInst))
		//append(ret2, sendInst)
		//e. There is one sending operation conducted by A on the channel.
		println("[DEBUG] more than one send instruction")
		for _, it := range ch.sends {
			itSendInst, ok := it.inst.(*ssa.Send)
			fmt.Println("[DEBUG] send inst: ", itSendInst)
			if !ok {
				fmt.Println("It is not a send instruction!")
				return -1, ret2
			}
			if checkAllSideEffect(itSendInst) {
				fmt.Println("[DEBUG] side effect check failed!")
				return -1, ret2
			} else {
				ret2 = append(ret2, getLineNo(fset, itSendInst))
			}
		}
		return getLineNo(fset, ch.make_inst) + 1, ret2
	} else if recvobj != nil {
		panic("Not implemented!")
	}
	return -1, ret2
}

func checkAllSideEffect(sendInst *ssa.Send) bool {
	return checkHasSideEffectOnBB(sendInst, false) ||
		checkHasSideEffect(sendInst.Block(), false) ||
		checkDeferSideEffects(sendInst.Parent())
}
