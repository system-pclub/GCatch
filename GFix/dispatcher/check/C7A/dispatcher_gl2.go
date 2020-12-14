package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/C7A/gl2"
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

func Map(vs []*ch_send, f func(send *ch_send) ssa.Instruction) []ssa.Instruction {
	vsm := make([]ssa.Instruction, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

/*type Iterable []interface{}
func (this Iterable) Map(fn func(x interface{}) interface{}) Iterable {
	ret := make(Iterable, len(this))
	for i, v := range this {
		ret[i] = fn(v)
	}
	return ret
}*/

func isExitBasicBlock(bb *ssa.BasicBlock) bool {
	//println(bb.Index, ":")
	//printSSAByBB(bb)
	for _, inst := range bb.Instrs {
		switch x := inst.(type) {
		case *ssa.Call:
			name := sync_check.CallName(x.Common())
			if name == "(*testing.common).Fatalf" || name == "(*testing.common).Fatal" || strings.Contains(name, "assert.") {
				println(name)
				println("[DEBUG] Got a Fatalf/Fatal")
				return true
			}
		case *ssa.Return:
			println("[DEBUG] Got a return")
			return true
		case *ssa.Panic:
			println("[DEBUG] Got a panic")
			return true
		}
	}
	return false
}

func getExitBasicBlocks(function *ssa.Function) []*ssa.BasicBlock {
	ret := make([]*ssa.BasicBlock, 0)
	for _, bb := range function.Blocks {
		if isExitBasicBlock(bb) {
			ret = append(ret, bb)
		}
	}
	return ret
}

func MapInstToBasicBlocks(insts []ssa.Instruction) []*ssa.BasicBlock {
	ret := make([]*ssa.BasicBlock, 0)
	for _, inst := range insts {
		ret = append(ret, inst.Block())
	}
	return ret
}

func findGoroutineB(ch *channel) *ssa.Function {
	var ret *ssa.Function = nil
	//we DO NOT assume those are from the same function
	for _, item := range ch.sends {
		if ret == nil {
			ret = item.inst.Parent()
		} else if ret != item.inst.Parent() {
			println("send operations have multiple parent functions.")
			return nil
		}
	}

	for _, item := range ch.closes {
		if ret == nil {
			ret = item.inst.Parent()
		} else if ret != item.inst.Parent() {
			println("close/send operations have multiple parent functions.")
			return nil
		}
	}
	return ret
}

func MapClosesToBasicBlocks(closes []*ch_close) []*ssa.BasicBlock {
	ret := make([]*ssa.BasicBlock, len(closes))
	for i, item := range closes {
		ret[i] = item.inst.Block()
	}
	return ret
}

func findEntry(bbs []*ssa.BasicBlock) *ssa.BasicBlock {
	if len(bbs) > 0 {
		return bbs[0].Parent().Blocks[0]
	} else {
		return nil
	}
}

func findGoInstSitesInFunc(calleeFunc *ssa.Function, callerFunc *ssa.Function) []*ssa.Go {
	//fmt.Println(calleeFunc)
	ret := make([]*ssa.Go, 0)
	//TODO: try to use call graph
	for _, bb := range callerFunc.Blocks {
		for _, ins := range bb.Instrs {
			goinst, ok := ins.(*ssa.Go)
			if ok {
				//fmt.Println(goinst)
				callInstCallee := goinst.Call.StaticCallee()
				fmt.Println("callInstCallee: ", callInstCallee)
				/*
					if callInstCallee != nil {
						fmt.Println(callInstCallee.Name())
					} else {
						fmt.Println(callinst.Call)
					}
					fmt.Println(calleeFunc.Name())*/
				if callInstCallee == calleeFunc { //TODO: use pointer info here
					fmt.Println(callInstCallee)
					ret = append(ret, goinst)
				}
			}
		}
	}
	return ret
}

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

func getGL2PatchLineNoNew(recvobj *ch_receive, ch *channel) (int, []int) {
	// get recv's thread from struct
	// get the thread from struct
	//c. The unblocked operation conducted by B could be skipped, due to panic or return.
	//d. The unblocked operation is only conducted once by B
	//e. There is no data race between instructions after the receiving operation in A and instructions before the unblocking operation in B.
	//a. The leak goroutine(s) are blocked at a receiving operation.
	//b. Suppose the leaked goroutine is A and the other goroutine that can unblock A is B.
	ret2 := make([]int, 0)
	recvInst, ok := recvobj.inst.(*ssa.UnOp)
	_, isSelect := recvobj.inst.(*ssa.Select)
	goroutineB := ch.make_inst.Parent() //TODO: check channel operations are in the same function
	if !ok && !isSelect {
		println("[DEBUG] This should not happen!")
		return -1, ret2
	}
	if recvobj.inst.Parent() == goroutineB {
		println("[DEBUG] We cannot fix this case as recv and make are in the same function.")
		return -1, ret2
	}
	if !isSelect && !isLastReceiveBeforeReturnInSameBB(recvInst) {
		println("[DEBUG] is not the last instruction before return!")
		return -1, ret2
	}
	closeBBs := MapClosesToBasicBlocks(ch.closes)
	sendInsts := Map(ch.sends, func(send *ch_send) ssa.Instruction { return send.inst })
	sendBBs := MapInstToBasicBlocks(sendInsts)
	//put all close and send bbs into one array, and mark them as visited before dfs
	makeGoroutine := ch.make_inst.Parent()

	pathFinder := gl2.NewPathFinder()
	exitBBs := getExitBasicBlocks(goroutineB)
	obstacleBBs := append(closeBBs, sendBBs...)
	println("len of obstacleBBs: ", len(obstacleBBs))
	println("len of exitBBs: ", len(exitBBs))
	entry := findEntry(obstacleBBs)
	fset := makeGoroutine.Prog.Fset

	//handle a special case. Not sure if it is in a more general case.
	if len(goroutineB.Blocks) == 1 {
		println("len(goroutineB.Blocks) == 1!")
		return -1, ret2
	}

	if !pathFinder.IsReachableToEntry(obstacleBBs, exitBBs, entry) {
		println("Not reachable to entry!")
		return -1, ret2
	}
	//strDefer := "defer LOC: "
	for _, x := range ch.closes {
		if x.thread.go_inst != nil && x.thread.go_inst != recvobj.thread.go_inst {
			fmt.Println(x.thread.go_inst)
			fmt.Println(recvobj.thread.go_inst)
			fmt.Println("[DEBUG] not the same go_inst!")
			return -1, ret2
		}
		ret2 = append(ret2, getLineNo(fset, x.inst))
	}
	for _, x := range sendInsts {
		sendInst, ok := x.(*ssa.Send)
		if ok && (isLastSendBeforeReturnInSameBB(sendInst) || isLastSendBeforeReturn(sendInst)) {
			ret2 = append(ret2, getLineNo(fset, x))
		} else {
			println("[DEBUG] is not the last send before return!")
			return -1, ret2
		}
	}
	return getLineNo(fset, ch.make_inst) + 1, ret2
}

//returns -1 if it is not a GL2 bug.
//returns the line number to insert defer operation if it is a GL2 bug.
func getGL2PatchLineNo(bug bug_report) (int, []int) {
	ret2 := make([]int, 0)
	recvInst := bug.recv.inst
	goroutineA := recvInst.Parent()
	goroutineB := findGoroutineB(bug.ch) //send and closes' parent
	if goroutineB == nil {
		println("cannot determine goroutine B")
		return -1, ret2
	}
	goInsts := findGoInstSitesInFunc(goroutineA, goroutineB)
	if len(goInsts) == 0 {
		goInsts = findGoInstSitesInFunc(goroutineB, goroutineA)
	}
	if len(goInsts) == 0 {
		println("Cannot find parent goroutine...")
		return -1, ret2
	} else if len(goInsts) > 1 {
		println("There are multiple go instructions...")
	}

	goInst := goInsts[0]
	if goroutineB == nil {
		fmt.Println("cannot find goroutine B!")
		return -1, ret2
	}

	closeBBs := MapClosesToBasicBlocks(bug.ch.closes)
	sendInsts := Map(bug.ch.sends, func(send *ch_send) ssa.Instruction { return send.inst })
	sendBBs := MapInstToBasicBlocks(sendInsts)
	//put all close and send bbs into one array, and mark them as visited before dfs
	makeGoroutine := bug.ch.make_inst.Parent()

	pathFinder := gl2.NewPathFinder()
	exitBBs := getExitBasicBlocks(goroutineB)
	obstacleBBs := append(closeBBs, sendBBs...)
	println("len of obstacleBBs: ", len(obstacleBBs))
	println("len of exitBBs: ", len(exitBBs))
	entry := findEntry(obstacleBBs)
	fset := makeGoroutine.Prog.Fset

	//handle a special case. Not sure if it is in a more general case.
	if len(goroutineB.Blocks) == 1 {
		return -1, ret2
	}

	if !pathFinder.IsReachableToEntry(obstacleBBs, exitBBs, entry) {
		return -1, ret2
	}
	//strDefer := "defer LOC: "
	for _, x := range bug.ch.closes {
		ret2 = append(ret2, getLineNo(fset, x.inst))
	}
	for _, x := range sendInsts {
		ret2 = append(ret2, getLineNo(fset, x))
	}
	if goroutineB == goInst.Parent() {
		if makeGoroutine == goroutineB {
			//fmt.Println(strDefer, getLineNo(fset, bug.ch.make_inst) + 1)
			//defer after the channel creation
			return getLineNo(fset, bug.ch.make_inst) + 1, ret2
		} else {
			//fmt.Println(strDefer, getLineNo(fset, goInst))
			//defer before the creation of A
			return getLineNo(fset, goInst), ret2
		}
	}
	if goroutineA == goInst.Parent() {
		//fmt.Println(strDefer, getLineNo(fset, goroutineB.Blocks[0].Instrs[0]))
		//defer at the beginning of the B
		return getLineNo(fset, goroutineB.Blocks[0].Instrs[0]), ret2
	}

	return -1, ret2
}
