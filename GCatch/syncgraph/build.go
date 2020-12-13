package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/path"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"go/token"
	"go/types"
	"sort"
	"strings"
)

var recursiveCount int

var MapMeetCircularDependPrims map[interface{}]struct{}

var DependMap map[interface{}]*DPrim

func BuildGraph(ch *instinfo.Channel, vecChannel []*instinfo.Channel, vecLocker []*instinfo.Locker, DMap map[interface{}]*DPrim) (*SyncGraph, error) {

	MapMeetCircularDependPrims = make(map[interface{}]struct{})

	DependMap = DMap

	// Before building: add ch into target primitive
	task := newTask()
	task.Step1AddPrim(ch)
	err := task.Step2CompletePrims()
	if err != nil {
		return nil, err
	}

	// create new graph
	newGraph := NewGraph(task)

	boolMetMake := false
	for _, ops := range task.MapLCARoot2Op {
		for _, op := range ops {
			if _, isMake := op.Op.(*instinfo.ChMake); isMake {
				if boolMetMake {
					if config.DISABLE_OPTIMIZATION_CALLEES {
						break
					}
					err := fmt.Errorf("multiple LCA can reach make")
					return nil, err
				}
				boolMetMake = true
				break
			}
		}
	}
	for _, ops := range task.MapLCARoot2Op {
		for _, op := range ops {
			if ch_op, isChanOp := op.Op.(instinfo.ChanOp); isChanOp {
				if strings.Contains(ch_op.Instr().Block().Comment, "rangechan") {
					if config.DISABLE_OPTIMIZATION_CALLEES {
						break
					}
					err := fmt.Errorf("range of channel")
					return nil, err
				}
			}
		}
	}



	for LCA, ops := range task.MapLCARoot2Op {
		// create a new goroutine for head_node
		newGoroutine := newGraph.NewGoroutine(LCA)
		newGraph.HeadGoroutines = append(newGraph.HeadGoroutines, newGoroutine)
		// If this LCA contains the make of ch, this is the MainGoroutine
		boolHasMake := false
		for _, op := range ops {
			if _, is_make := op.Op.(*instinfo.ChMake); is_make {
				boolHasMake = true
				break
			}
		}
		if boolHasMake {
			if newGraph.MainGoroutine != nil {
				//err := fmt.Errorf("Warning: we found more than one MainGoroutine:\n",
				//	newGraph.MainGoroutine.EntryFn,"\n",newGoroutine.EntryFn)
				//return nil, err
			}
			newGraph.MainGoroutine = newGoroutine
		}
		// create a new ctx
		newCtx := newGraph.NewCtx(newGoroutine, LCA)
		// process the head
		recursiveCount = 0
		newGoroutine.HeadNode = ProcessInstGetNode(firstInstOfFn(LCA), newCtx)
	}

	vecHandledUnfinish := []*Unfinish{}

	for len(newGraph.Worklist) > 0 {

		var doThis *Unfinish

		// if another Unfinish.Site is the same as a handled one, do this; this is for a problem in my implementation of defer.
		// 	For example, `defer A(); if then { return 0;} else { return 2;}` If we don't have this step,
		// 	it is possible that `defer A()` will only be analyzed at return 0, and will be skipped at return 2
		for _, unfinish := range newGraph.Worklist {
			for _, handled := range vecHandledUnfinish {
				if handled.Site == unfinish.Site {
					doThis = unfinish
				}
			}
		}

		if doThis == nil {
			// if some unfinish is creating an anonymous function, do this
			doThis = newGraph.findUnfinishGoOfMainGoroutine(newGraph.Worklist)
		}


		newGraph.Task.Update()
		if newGraph.Task.BoolFinished && doThis == nil {
			break
		}

		if doThis == nil {
			wanted_chains := newGraph.Task.WantedList()
			wanted_chains = removeVisitedChains(wanted_chains, newGraph.Visited)
			// if someone is in our wanted list, do this
			for _,unfinished := range newGraph.Worklist {
				has_match := false
				for _,wanted := range wanted_chains {
					if wanted.Contains(unfinished.Ctx.CallChain) {
						has_match = true
						break
					}
				}
				if has_match {
					doThis = unfinished
					break
				}
			}
		}

		/// Code for entering all callees
		if config.DISABLE_OPTIMIZATION_CALLEES {
			if doThis == nil {
				doThis = newGraph.Worklist[0]
			}
			if len(vecHandledUnfinish) == 2000 {
				fmt.Println("More than 2000 handledUnfinish")
			}
		}

		if doThis != nil {
			vecHandledUnfinish = append(vecHandledUnfinish, doThis)
			// we can be sure that nextInst is not nil
			nextInst := doThis.Site.Callee.Func.Blocks[0].Instrs[0]
			switch callerNode := doThis.Unfinished.(type) {
			case *Call:
				callerNode.Calling[doThis.Site] = ProcessInstGetNode(nextInst, doThis.Ctx)
				boolNothingLeft := true
				for _,target := range callerNode.Calling {
					if target == nil {
						boolNothingLeft = false
						break
					}
				}
				if boolNothingLeft {
					newGraph.NodeStatus[callerNode].Str = Done
				}
			case *Go:
				headOfGoroutine := ProcessInstGetNode(nextInst, doThis.Ctx)
				callerNode.MapCreateNodes[doThis.Site] = headOfGoroutine
				workingGoroutine := callerNode.MapCreateGoroutines[doThis.Site]
				workingGoroutine.HeadNode = headOfGoroutine
				newGraph.Goroutines = append(newGraph.Goroutines, workingGoroutine)
				boolNothingLeft := true
				for _,target := range callerNode.MapCreateNodes {
					if target == nil {
						boolNothingLeft = false
						break
					}
				}
				if boolNothingLeft {
					newGraph.NodeStatus[callerNode].Str = Done
				}
			}

			newGraph.Visited = append(newGraph.Visited, doThis.Ctx.CallChain)
			newGraph.Worklist = removeFromWorklist(newGraph.Worklist, doThis)
		} else {
			// no one on wanted list, break
			// TODO: this logic may be incorrect
			fmt.Println("No Unfinished matches wanted list. Exit building")
			break
		}
	}

	// check if task is fulfilled
	newGraph.Task.Update()
	if newGraph.Task.BoolFinished {
		//fmt.Println("A graph is finished")
	} else {
		err := fmt.Errorf("A graph is not finished")
		return nil, err
	}

	// Some remaining fields to fill
	newGraph.fillSyncOp()
	newGraph.BuildNodeInOut()

	recursiveCount = 0
	return newGraph, nil
}

func newNormal(inst ssa.Instruction, ctx *CallCtx) (*NormalInst,*Status) {
	normal := &NormalInst{
		Inst: inst,
		Next: nil,
		node: node{
			Instr: inst,
			Ctx:   ctx,
		},
	}

	newStatus := storeGraphInfo(inst,ctx, normal)

	return normal, newStatus
}

func storeGraphInfo(inst ssa.Instruction, ctx *CallCtx, node Node) *Status {
	key := InstCtxKey{
		Inst: inst,
		Ctx:  ctx,
	}
	ctx.Graph.MapInstCtxKey2Node[key] = node
	newStatus := &Status{
		Str:     In_progress,
		Visited: 1,
	}
	ctx.Graph.NodeStatus[node] = newStatus
	return newStatus
}

func storeGraphInfoForDefer(inst ssa.Instruction, ctx *CallCtx, node Node) *Status {
	key := InstCtxKey{
		Inst: inst,
		Ctx:  ctx,
	}
	ctx.Graph.MapInstCtxKey2Defer[key] = append(ctx.Graph.MapInstCtxKey2Defer[key], node)
	newStatus := &Status{
		Str:     In_progress,
		Visited: 1,
	}
	ctx.Graph.NodeStatus[node] = newStatus
	return newStatus
}

func updateTaskOp(op interface{}, prim interface{}, ctx *CallCtx, inst ssa.Instruction) {
	tPrim,ok := ctx.Graph.Task.MapValue2TaskPrimitive[prim]
	if ok {
		chainsToReachOp,ok := tPrim.Ops[op]
		if !ok {
			fmt.Println("An op's primitive is in task, but op is not found in primitive.Ops")
			output.PrintIISrc(inst)
			return
		}
		for i, chain := range chainsToReachOp.Chains {
			if chain.Equal(ctx.CallChain) {
				chainsToReachOp.VecBoolIsChainFinished[i] = true
				break
			}
		}
	}
}

func ProcessInstGetNode(targetInst ssa.Instruction, ctx *CallCtx) Node {

	recursiveCount++
	if recursiveCount > config.MAX_INST_IN_SYNCGRAPH {
		fmt.Println("Warning in ProcessInstGetNode: reached MAX_INST_IN_SYNCGRAPH")
		newEnd := &End{
			Inst:   targetInst,
			Reason: MaxRecursive,
			node:      node{
				Instr: targetInst,
				Ctx:   ctx,
			},
		}
		return newEnd
	}

	if targetInst == nil {
		fmt.Println("Fatal in ProcessInstGetNode: nil inst")
	}

	key := InstCtxKey{
		Inst: targetInst,
		Ctx:  ctx,
	}
	if existNode,ok := ctx.Graph.MapInstCtxKey2Node[key]; ok {
		return existNode
	}

	// Check for dependency map
	metPrims := []*DPrim{}
	vecChOp, ok := instinfo.MapInst2ChanOp[targetInst]
	if ok {
		for _, chOp := range vecChOp {
			isBlocking := false
			switch chOp.(type) {
			case *instinfo.ChSend, *instinfo.ChRecv:
				isBlocking = true
			}
			if isBlocking == false { // Not interested in non blocking operations
				continue
			}

			metPrimNode,ok := DependMap[chOp.Prim()]
			if ok {
				metPrims = append(metPrims, metPrimNode)
			}
		}
	} else {
		muOp, ok := instinfo.MapInst2LockerOp[targetInst]
		lockOp, isBlocking := muOp.(*instinfo.LockOp)
		if ok && isBlocking {
			metPrimNode,ok := DependMap[lockOp.Prim()]
			if ok {
				metPrims = append(metPrims, metPrimNode)
			}
		}
	}
	for _, metPrimNode := range metPrims {
		for _, targetPrim := range ctx.Graph.Task.VecTaskPrimitive {
			targetPrimNode,ok := DependMap[targetPrim.Primitive]
			if ok {
				isCircularDepend := false
				for _, depend := range targetPrimNode.Circular_depend {
					if depend.Callee == metPrimNode {
						isCircularDepend = true
						break
					}
				}
				if isCircularDepend {
					MapMeetCircularDependPrims[metPrimNode] = struct{}{}
				}
			}
		}
	}

	switch inst := targetInst.(type) {

	case *ssa.MakeChan:
		var op instinfo.ChanOp

		ops,ok := instinfo.MapInst2ChanOp[inst] // The slice op can at most have 1 element
		if !ok {
			fmt.Println("Warning in ProcessInstGetNode: can't find op for a channel make")
			output.PrintIISrc(inst)
			//op = instinfo.Anytime_make(inst)
			normal, newStatus := newNormal(inst,ctx)

			normal.Next = ProcessInstGetNode(nextInst(inst), ctx)
			newStatus.Str = Done

			return normal
		} else {
			if len(ops) > 1 { // when channel can be overwritten, vecChOp can have multiple elements
				//debugPrintMultiChAlias(inst,ops,"chan make")
			}
			op = ops[0]
		}
		ch := op.(*instinfo.ChMake).Parent

		updateTaskOp(op,ch,ctx,inst)

		newMakeChan := &ChanMake{
			Inst:    inst,
			Channel: ch,
			MakeOp:  instinfo.MapInst2ChanOp[inst][0],
			Next:    nil,
			syncNode: syncNode{
				Prim:                  ch,
				BoolIsAllAliasInGraph: ctx.Graph.Task.IsPrimATarget(ch),
				AliasOp:               make(map[SyncOp]bool),
				SyncOp:                make(map[SyncOp]bool),
				node:                  node{
					Instr: inst,
					Ctx:   ctx,
				},
			},
		}
		ctx.Graph.MapPrim2VecSyncOp[ch] = append(ctx.Graph.MapPrim2VecSyncOp[ch], newMakeChan)

		newStatus := storeGraphInfo(inst,ctx, newMakeChan)

		newMakeChan.Next = ProcessInstGetNode(nextInst(inst),ctx)
		newStatus.Str = Done

		return newMakeChan


	case *ssa.Jump:
		newJump := &Jump{
			Inst:            inst,
			Next:            nil,
			BoolIsBackedge:  false,
			BoolIsNextexist: false,
			node:        node{
				Instr: inst,
				Ctx:   ctx,
			},
		}

		nextBB := inst.Block().Succs[0]
		if len(nextBB.Instrs) == 0 {
			fmt.Println("Warning in ProcessInstGetNode: a jump's target is an empty bb")
			output.PrintIISrc(inst)

			newStatus := storeGraphInfo(inst,ctx, newJump)
			newStatus.Str = Done
			return newJump
		}

		nextInst := nextBB.Instrs[0]
		key := InstCtxKey{
			Inst: nextInst,
			Ctx:  ctx,
		}
		nextNode,ok := ctx.Graph.MapInstCtxKey2Node[key]
		if ok{
			newJump.BoolIsNextexist = true
			if ctx.Graph.NodeStatus[nextNode].Str == In_progress {
				newJump.BoolIsBackedge = true
			}
		}

		newStatus := storeGraphInfo(inst,ctx, newJump)

		if ok {
			newJump.Next = nextNode
		} else {
			// inst not met before
			newJump.Next = ProcessInstGetNode(nextInst,ctx)
		}
		newStatus.Str = Done

		return newJump

	case *ssa.If:
		newIf := &If{
			Inst:               inst,
			Cond:               inst.Cond,
			Then:               nil,
			Else:               nil,
			BoolIsThenBackedge: false,
			BoolIsElseBackedge: false,
			node:             node{
				Instr: inst,
				Ctx:   ctx,
			},
		}

		thenBB := inst.Block().Succs[0]
		elseBB := inst.Block().Succs[1]
		if len(thenBB.Instrs) == 0 || len(elseBB.Instrs) == 0 {
			fmt.Println("Warning in ProcessInstGetNode: a If's target is an empty bb")
			output.PrintIISrc(inst)
			return nil
		}
		thenInst := thenBB.Instrs[0]
		elseInst := elseBB.Instrs[0]
		thenKey := InstCtxKey{
			Inst: thenInst,
			Ctx:  ctx,
		}
		elseKey := InstCtxKey{
			Inst: elseInst,
			Ctx:  ctx,
		}
		thenNode, thenOk := ctx.Graph.MapInstCtxKey2Node[thenKey]
		if thenOk {
			if ctx.Graph.NodeStatus[thenNode].Str == In_progress {
				newIf.BoolIsThenBackedge = true
			}
		}
		elseNode, elseOk := ctx.Graph.MapInstCtxKey2Node[elseKey]
		if elseOk {
			if ctx.Graph.NodeStatus[elseNode].Str == In_progress {
				newIf.BoolIsElseBackedge = true
			}
		}

		newStatus := storeGraphInfo(inst,ctx, newIf)

		if thenOk {
			newIf.Then = thenNode
		} else {
			newIf.Then = ProcessInstGetNode(thenInst,ctx)
		}
		if elseOk {
			newIf.Else = elseNode
		} else {
			newIf.Else = ProcessInstGetNode(elseInst,ctx)
		}
		newStatus.Str = Done

		return newIf

	case *ssa.Call, *ssa.RunDefers:

		todoInsts := []ssa.CallInstruction{}
		var flagRundefer bool
		if rundefer,ok := inst.(*ssa.RunDefers); ok {
			flagRundefer = true
			vecAllDefers,ok := config.Inst2Defers[rundefer]
			if !ok {
				newNormal, newStatus := newNormal(inst,ctx)

				newNormal.Next = ProcessInstGetNode(nextInst(inst), ctx)
				newStatus.Str = Done

				return newNormal
			} else {
				for _, aDefer := range vecAllDefers {
					todoInsts = append(todoInsts, aDefer)
				}
			}
		} else if call,ok := inst.(*ssa.Call); ok {
			flagRundefer = false
			todoInsts = []ssa.CallInstruction{call}
		}

		firstNode := handleTodoInsts(inst, ctx, flagRundefer, todoInsts)

		return firstNode

	case *ssa.Return:
		newReturn := &Return{
			Inst:                 inst,
			BoolIsEndOfGoroutine: inst.Parent() == ctx.Goroutine.EntryFn,
			Caller:               ctx.CallSite,
			node:                node{
				Instr: inst,
				Ctx:   ctx,
			},
		}

		newStatus := storeGraphInfo(inst,ctx, newReturn)
		newStatus.Str = Done

		return newReturn

	case *ssa.Go:
		newGo := &Go{
			Inst:                inst,
			MapCreateGoroutines: make(map[*callgraph.Edge]*Goroutine),
			MapCreateNodes:      make(map[*callgraph.Edge]Node),
			NextLocal:           nil,
			node:       node{
				Instr: inst,
				Ctx:   ctx,
			},
		}

		newStatus := storeGraphInfo(inst,ctx, newGo)

		// An Go inst may have multiple call edges if calling interface method
		mapAllEdges,ok := config.Inst2CallSite[inst]
		if ok {
			for edge, _ := range mapAllEdges {
				callee := edge.Callee.Func
				if callee == nil || len(callee.Blocks) == 0 || len(callee.Blocks[0].Instrs) == 0 {
					continue
				}

				newEdgePath := &path.EdgeChain{
					Chain:  append(ctx.CallChain.Chain, edge),
					Start: ctx.CallChain.Start,
				}

				newGoroutine := &Goroutine{
					Creator: newGo,
					EntryFn: callee,
					IsMain:  false,
					Graph:   ctx.Graph,
				}

				newGo.MapCreateNodes[edge] = nil
				newGo.MapCreateGoroutines[edge] = newGoroutine

				newCtx := &CallCtx{
					CallChain: newEdgePath,
					Goroutine: newGoroutine,
					CallSite:  newGo,
					Graph:     ctx.Graph,
				}

				newUnfinish := &Unfinish{
					UnfinishedFn: inst.Parent(),
					Unfinished:    newGo,
					IsGo:          true,
					Site:          edge,
					Dir:           true,
					Ctx:           newCtx,
				}

				ctx.Graph.Worklist = append(ctx.Graph.Worklist, newUnfinish)
			}
		} else {
			//fmt.Println("Warning in ProcessInstGetNode: Go: can't find Sites of inst by config.Inst2CallSite")
			//output.PrintIISrc(inst)
		}

		newGo.NextLocal = ProcessInstGetNode(nextInst(inst), ctx)
		newStatus.Str = Done

		return newGo

	case *ssa.Select:
		newSelect := &Select{
			Inst:           inst,
			Cases:          make(map[int]*SelectCase),
			BoolHasDefault: !inst.Blocking,
			DefaultCase:    nil,
			node:         node{
				Instr: inst,
				Ctx:   ctx,
			},
		}

		// Prepare the ops for each case. Note that some cases may have no op, if pointer analysis failed to find
		ops, ok := instinfo.MapInst2ChanOp[inst]
		if !ok {
			fmt.Println("Warning in ProcessInstGetNode: a select has no corresponding OP")
			output.PrintIISrc(inst)
		}

		// Update task
		for _, op := range ops {
			switch concrete := op.(type) {
			case *instinfo.ChSend:
				ch := concrete.Parent
				updateTaskOp(op,ch,ctx,inst)
			case *instinfo.ChRecv:
				ch := concrete.Parent
				updateTaskOp(op,ch,ctx,inst)
			}
		}

		index2op := make(map[int]instinfo.ChanOp)
		for _,op := range ops {
			var index int
			switch concrete := op.(type) {
			case *instinfo.ChSend:
				if concrete.CaseIndex == -1 {
					fmt.Println("Warning in ProcessInstGetNode: a send is not in select, but we are dealing with select inst:")
					output.PrintIISrc(inst)
					continue
				}
				index = concrete.CaseIndex
			case *instinfo.ChRecv:
				if concrete.CaseIndex == -1 {
					fmt.Println("Warning in ProcessInstGetNode: a recv is not in select, but we are dealing with select inst:")
					output.PrintIISrc(inst)
					continue
				}
				index = concrete.CaseIndex
			}
			_,ok := index2op[index]
			if ok { // This index is used by another op, which doesn't make sense
				//fmt.Println("Warning in ProcessInstGetNode: one case index show up twice in a select:")
				//output.PrintIISrc(inst)
			}
			index2op[index] = op
		}
		if inst.Blocking == false { // Has Default
			index2op[-1] = nil
		}
		// for those cases without op, we create an unsure op
		for i := 0; i < len(inst.States); i++ {
			if _,ok := index2op[i]; ok {
				continue
			}
			switch inst.States[i].Dir {
			case types.SendOnly:
				op := instinfo.AddNotDependSend(inst)
				index2op[i] = op
				op.IsCaseBlocking = inst.Blocking
				op.CaseIndex = i
			case types.RecvOnly:
				op := instinfo.AddNotDependRecv(inst)
				index2op[i] = op
				op.IsCaseBlocking = inst.Blocking
				op.CaseIndex = i
			}
		}

		index2next := make(map[int]ssa.Instruction)
		index2next,err := path.FindSelectNexts(inst)
		if err != nil {
			fmt.Println("Fatal error: can't deal with a select, because:")
			fmt.Println(err.Error())
			return newSelect
		}

		// Check the 2 maps are consistent
		if len(index2op) != len(index2next) {
			fmt.Println("Fatal error: can't deal with a select, because index2op is not the same length as index2next. Select:")
			output.PrintIISrc(inst)
			return newSelect
		}
		for i,_ := range index2op {
			flagFound := false
			for j,_ := range index2next {
				if i == j {
					flagFound = true
					break
				}
			}
			if flagFound == false {
				fmt.Println("Fatal error: can't deal with a select, because index2op doesn't contains the same indexes as index2next. Select:")
				output.PrintIISrc(inst)
				return newSelect
			}
		}

		type opNext struct {
			op instinfo.ChanOp
			next ssa.Instruction
		}

		mapIndex2opNext := make(map[int]opNext)
		for index,op := range index2op {
			next := index2next[index]
			mapIndex2opNext[index] = opNext{
				op:   op,
				next: next,
			}
		}

		newSelectStatus := storeGraphInfo(inst,ctx, newSelect)

		indexs := []int{}
		for i,_ := range mapIndex2opNext {
			indexs = append(indexs, i)
		}
		sort.Ints(indexs)

		for _,index := range indexs {
			opNext := mapIndex2opNext[index]
			var ch *instinfo.Channel
			switch concrete := opNext.op.(type) {
			case *instinfo.ChSend:
				ch = concrete.Parent
			case *instinfo.ChRecv:
				ch = concrete.Parent
			}
			newSelectcase := &SelectCase{
				Channel:        ch,
				Op:             opNext.op,
				BoolIsDefault:  index == -1,
				Index:          index,
				Next:           nil,
				BoolIsBackedge: false,
				Select:         newSelect,
				syncNode:  syncNode{
					Prim:                  ch,
					BoolIsAllAliasInGraph: ctx.Graph.Task.IsPrimATarget(ch),
					AliasOp:               make(map[SyncOp]bool),
					SyncOp:                make(map[SyncOp]bool),
					node:                  node{
						Instr: inst,
						Ctx:   ctx,
					},
				},
			}
			key := InstCtxKey{
				Inst: opNext.next,
				Ctx:  ctx,
			}
			nextNode,ok := ctx.Graph.MapInstCtxKey2Node[key]
			if ok{
				if ctx.Graph.NodeStatus[nextNode].Str == In_progress {
					newSelectcase.BoolIsBackedge = true
				}
			}
			ctx.Graph.MapPrim2VecSyncOp[ch] = append(ctx.Graph.MapPrim2VecSyncOp[ch], newSelectcase)

			ctx.Graph.Select2Case[newSelect] = append(ctx.Graph.Select2Case[newSelect], newSelectcase)
			newCaseStatus := &Status{
				Str:     In_progress,
				Visited: 1,
			}
			ctx.Graph.NodeStatus[newSelectcase] = newCaseStatus

			newSelectcase.Next = ProcessInstGetNode(opNext.next, ctx)
			newCaseStatus.Str = Done

			newSelect.Cases[index] = newSelectcase
			if index == -1 {
				newSelect.DefaultCase = newSelectcase
			}
		}

		newSelectStatus.Str = Done
		return newSelect

	case *ssa.Send:
		var chOp instinfo.ChanOp
		chOps,ok := instinfo.MapInst2ChanOp[inst]
		if !ok {
			fmt.Println("Warning in ProcessInstGetNode: can't find op for a send")
			output.PrintIISrc(inst)
			chOp = instinfo.AddNotDependSend(inst)
		}

		if len(chOps) == 0 { // when channel can be overwritten, vecChOp can have multiple elements
			//debugPrintMultiChAlias(inst,vecChOp,"send")
			new_normal, new_status := newNormal(inst,ctx)

			new_normal.Next = ProcessInstGetNode(nextInst(inst), ctx)
			new_status.Str = Done

			return new_normal
		}

		chOp = chOps[0]
		sendOp := chOp.(*instinfo.ChSend)

		ch := sendOp.Parent
		updateTaskOp(sendOp,ch,ctx,inst)

		newSend := &ChanOp{
			Channel: sendOp.Parent,
			Op:      sendOp,
			Next:    nil,
			syncNode:      syncNode{
				Prim:                  sendOp.Parent,
				BoolIsAllAliasInGraph: ctx.Graph.Task.IsPrimATarget(sendOp.Parent),
				AliasOp:               make(map[SyncOp]bool),
				SyncOp:                make(map[SyncOp]bool),
				node:                  node{
					Instr: inst,
					Ctx:   ctx,
				},
			},
		}
		ctx.Graph.MapPrim2VecSyncOp[sendOp.Parent] = append(ctx.Graph.MapPrim2VecSyncOp[sendOp.Parent], newSend)

		newStatus := storeGraphInfo(inst,ctx, newSend)

		newSend.Next = ProcessInstGetNode(nextInst(inst), ctx)
		newStatus.Str = Done

		return newSend

	case *ssa.UnOp:
		if inst.Op == token.ARROW {
			var chOp instinfo.ChanOp
			vecChOp, ok := instinfo.MapInst2ChanOp[inst] //vecChOp can have at most 1 element

			if !ok {
				fmt.Println("Warning in ProcessInstGetNode: can't find op for a receive")
				output.PrintIISrc(inst)
				//chOp = instinfo.Anytime_recv(inst)
				newNormal, newStatus := newNormal(inst,ctx)

				newNormal.Next = ProcessInstGetNode(nextInst(inst), ctx)
				newStatus.Str = Done

				return newNormal
			} else {
				if len(vecChOp) > 1 { // when channel can be overwritten, vecChOp can have multiple elements
					//debugPrintMultiChAlias(inst,vecChOp,"receive")
				}
				chOp = vecChOp[0]
			}

			recvOp := chOp.(*instinfo.ChRecv)

			ch := recvOp.Parent
			updateTaskOp(recvOp,ch,ctx,inst)

			newRecv := &ChanOp{
				Channel: ch,
				Op:      chOp,
				Next:    nil,
				syncNode:      syncNode{
					Prim:                  ch,
					BoolIsAllAliasInGraph: ctx.Graph.Task.IsPrimATarget(ch),
					AliasOp:               make(map[SyncOp]bool),
					SyncOp:                make(map[SyncOp]bool),
					node:                  node{
						Instr: inst,
						Ctx:   ctx,
					},
				},
			}
			ctx.Graph.MapPrim2VecSyncOp[ch] = append(ctx.Graph.MapPrim2VecSyncOp[ch], newRecv)

			newStatus := storeGraphInfo(inst,ctx, newRecv)

			newRecv.Next = ProcessInstGetNode(nextInst(inst), ctx)
			newStatus.Str = Done

			return newRecv

		}

		newNormal, newStatus := newNormal(inst,ctx)

		newNormal.Next = ProcessInstGetNode(nextInst(inst), ctx)
		newStatus.Str = Done

		return newNormal

	case *ssa.Defer:
		return ProcessInstGetNode(nextInst(inst),ctx)

	case *ssa.Panic:
		newKill := &Kill{
			Inst:        inst,
			BoolIsPanic: true,
			BoolIsFatal: false,
			node:     node{
				Instr: inst,
				Ctx:   ctx,
			},
		}
		newStatus := storeGraphInfo(inst,ctx, newKill)

		vecAllDefers,ok := config.Inst2Defers[inst]
		if !ok {
			newKill.Next = nil
		} else {
			todoInsts := []ssa.CallInstruction{}
			for _,a_defer := range vecAllDefers {
				todoInsts = append(todoInsts,a_defer)
			}
			firstNode := handleTodoInsts(inst, ctx, true, todoInsts)

			newKill.Next = firstNode
		}

		newStatus.Str = Done

		return newKill

	default:
		newNormal, new_status := newNormal(inst,ctx)

		newNormal.Next = ProcessInstGetNode(nextInst(inst), ctx)
		new_status.Str = Done

		return newNormal
	}
}