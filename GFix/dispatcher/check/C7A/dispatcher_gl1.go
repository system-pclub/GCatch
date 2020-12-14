package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa/ssautil"
)

var debug bool = true

func findGoInstSites(calleeFunc *ssa.Function, program *ssa.Program) []*ssa.Go {
	//fmt.Println(calleeFunc)
	ret := make([]*ssa.Go, 0)
	for fn := range ssautil.AllFunctions(program) {
		for _, bb := range fn.Blocks {
			for _, ins := range bb.Instrs {
				goinst, ok := ins.(*ssa.Go)
				if ok {
					//fmt.Println(goinst)
					callInstCallee := goinst.Call.StaticCallee()
					/*
						if callInstCallee != nil {
							fmt.Println(callInstCallee.Name())
						} else {
							fmt.Println(callinst.Call)
						}
						fmt.Println(calleeFunc.Name())*/
					if callInstCallee == calleeFunc {
						ret = append(ret, goinst)
					}
				}
			}
		}
	}
	return ret
}

func hasDirectCallSites(calleeFunc *ssa.Function, program *ssa.Program) bool {
	for fn := range ssautil.AllFunctions(program) {
		for _, bb := range fn.Blocks {
			for _, ins := range bb.Instrs {
				callinst, ok := ins.(*ssa.Call)
				if ok {
					callInstCallee := callinst.Call.StaticCallee()
					/*
						if callInstCallee != nil {
							fmt.Println(callInstCallee.Name())
						} else {
							fmt.Println(callinst.Call)
						}
						fmt.Println(calleeFunc.Name())*/
					if callInstCallee == calleeFunc {
						return true
					}
				}
			}
		}
	}
	return false
}

func isSendAndJump(sendInst *ssa.Send) bool {
	bb := sendInst.Block()
	foundSendInst := false
	for _, inst := range bb.Instrs {
		switch x := inst.(type) {
		case *ssa.Send:
			if x == sendInst {
				foundSendInst = true
			}
		case *ssa.Jump:
			if !foundSendInst {
				return false
			}
		default:
			if foundSendInst {
				return false
			}
		}
	}
	return true
}

func isLastSendBeforeReturn(sendInst *ssa.Send) bool {
	bb_start := sendInst.Block()
	printSSAByBB(bb_start)
	if !isSendAndJump(sendInst) {
		return false
	}
	println("[DEBUG] is send and jump.")
	if len(bb_start.Succs) == 1 {
		succ := bb_start.Succs[0]
		printSSAByBB(succ)
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
	} else {
		println("[DEBUG] more than one succ bbs.")
	}
	return false
}

func isLastSendBeforeReturnInSameBB(sendInst *ssa.Send) bool {
	bb := sendInst.Block()
	state := 0
	//foundDeref := false
	for index, inst := range bb.Instrs {
		switch x := inst.(type) {
		case *ssa.Send:
			if x == sendInst {
				if state != 0 {
					return false
				} else {
					state = 1
				}
			}
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
			if state == 1 {
				_, ok2 := bb.Instrs[index+1].(*ssa.Call)
				if !ok2 {
					return false
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

func findChannelMakeInBB(bb *ssa.BasicBlock) *ssa.MakeChan {
	/*printSSAByBB(bb)
	for _, item := range bb.Instrs {
		makeChanInst, ok := item.(*ssa.MakeChan)
		if ok {
			_ = makeChanInst
		}
	}*/
	return nil
}

func countSend(ch *channel, gor goroutine) int {
	ret := 0
	for _, send := range ch.sends {
		if send.thread == gor {
			ret += 1
		}
	}
	return ret
}

func noRaceCondition() bool {
	return false
}

func isGL1_new(sendobj *ch_send, ch *channel) bool {
	//a. The channel is an unbuffered channel. (Checked by the checker)
	//b. The leaked goroutine (named A) is blocked at a sending operation (Checked by the checker)
	//c. Suppose the leaked goroutine is A and the other goroutine that can unblock A is B.
	//c1. There is no race condition between instructions after A’s sending operation
	//and instructions before B’s receiving operation.
	//d. The sending operation is the last operation of A:
	//d1. the sending operation is not in another called function by A. i.e., the sending operation's function is the
	//same as the function that the goroutine creates.
	goroutineA := sendobj.inst.Parent()     //check if parent function are the same as the goroutine
	goroutineInst := sendobj.thread.go_inst //
	if !CheckCallRelationship(goroutineA, goroutineInst) && sendobj.foundByLineNumberFromSSA == false {
		return false
	}
	//d2. the sending operation is the last operation of the function.
	send := sendobj.inst

	sendInst, ok := send.(*ssa.Send)
	if !ok {
		if debug {
			fmt.Println("Not a send inst!") //TODO: not a correct way to handle error
		}
		return false
	}
	b := checkSafeToUnblock(sendInst)
	if !b {

		return false
	}
	//There should be only one child goroutine that uses the channel.
	for _, itSend := range ch.sends {
		if itSend.thread != ch.main_thread && itSend.thread != sendobj.thread {
			println("the channel was used in more than one child goroutine!")
			return false
		}
	}
	//			fmt.Println("[DEBUG] working on send operation ", it.inst)
	//e. There is one sending operation conducted by A on the channel.
	if countSend(ch, sendobj.thread) != 1 {
		for _, it := range ch.sends {
			if it.thread == sendobj.thread && it.thread != ch.main_thread {
				if it.inst.Parent() == sendInst.Parent() {
					itSendInst, ok := it.inst.(*ssa.Send)
					if !ok {
						fmt.Println("It is not a send instruction!")
						return false
					}
					if !checkSafeToUnblock(itSendInst) {
						fmt.Println("It is not the last instruction before return!")
						return false
					}
				} else {
					fmt.Println("It is not in the same function as the send instruction!", it.inst.Parent().Name(),
						sendInst.Parent().Name())
					return false
				}
			} else {
				fmt.Println("It is not in the same thread!")
				return false
			}
		}
	}
	//f. A is not created inside a loop, unless the sending operation conducted by the leaked goroutine
	//is paired with a receiving operation inside the loop.
	var goinst *ssa.Go
	var parentFunc *ssa.Function
	if sendobj.foundByLineNumberFromSSA {
		parentFunc = ch.make_inst.Parent()
		goinsts := findGoInstSitesInFunc(sendInst.Parent(), ch.make_inst.Parent())
		if len(goinsts) > 1 || len(goinsts) == 0 {
			fmt.Println("[DEBUG] len(goinsts) == ", len(goinsts))
			return false
		} else {
			goinst = goinsts[0]
		}
	} else {
		goinst = sendobj.thread.go_inst
		parentFunc = sendobj.thread.go_inst.Parent()
	}

	loopinfo := NewLoopInfo(parentFunc)
	loopinfo.Analyze()
	_, ok = loopinfo.isLoopBB[goinst.Block()]
	/*for bb, _ := range loopinfo.isLoopBB {
		println("bb ", bb.Index, ":")
		printSSAByBB(bb)
	}*/
	/*if debug {
		println(ok)
	}*/
	if ok {
		println("[DEBUG] is in a loop!")
		makeinst := ch.make_inst
		if makeinst.Block() == goinst.Block() {
			return true
		}
		//findChannelMakeInBB(goinst.Block())
		//findChannelMake(sendInst)
		if debug {
			println("[DEBUG] is in a loop and make and send are not in the same basic block!")
		}
		return false
	}
	return true

}

func CheckCallRelationship(goroutineA *ssa.Function, goroutineInst *ssa.Go) bool {
	if goroutineInst == nil {
		println("[DEBUG] goroutineInst == nil")
		return false
	}

	//println("[DEBUG]", goroutineA, goroutineInst)
	if goroutineA != goroutineInst.Call.StaticCallee() { //sometimes this line crashes because of nil
		prtfunc := goroutineInst.Parent()
		nodes := global.Call_graph.Nodes[prtfunc]
		_ = nodes
		foundCallSite := false
		for _, node := range nodes.Out {
			if node.Site == goroutineInst && node.Callee.Func == goroutineA {
				foundCallSite = true
				break
			}
		}
		if !foundCallSite {
			println("the parent function of sendobj is not the same as goroutine function.")
			//TODO: examine complex call chains and not directly defined goroutine functions. But here there is a bit bug.
			return false
		}
	}
	return true
}

func checkSafeToUnblock(sendInst *ssa.Send) bool {
	println("[DEBUG] checking if safe to unblock")
	if checkDeferSideEffects(sendInst.Parent()) {
		return false
	}
	if !(isLastSendBeforeReturnInSameBB(sendInst) || isLastSendBeforeReturn(sendInst)) {
		if debug {
			printSSAByBB(sendInst.Block())
			fmt.Println("is not the last send before return!")
		}
		lp := NewLoopInfo(sendInst.Parent())
		lp.Analyze()
		_, ok := lp.isLoopBB[sendInst.Block()]
		if ok {
			return false
		}
		if !checkHasSideEffectOnBB(sendInst, true) &&
			!checkHasSideEffect(sendInst.Block(), true) {
			return true
		}
		return false
	}
	return true
}
