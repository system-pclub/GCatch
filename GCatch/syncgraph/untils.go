package syncgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
)

// nextInst mustn't be called upon jump/if/return/panic, which have no or multiple nextInst
func nextInst(inst ssa.Instruction) ssa.Instruction {
	bb := inst.Block()
	insts := bb.Instrs
	for i, other := range insts {
		if other == inst {
			if len(insts) == i+1 {
				return nil
			} else {
				return insts[i+1]
			}
		}
	}
	return nil
}

func checkNodesNumber(newGraph *SyncGraph) {
	// Check if nodes in graph.MapInstCtxKey2Node + graph.Select2Case = graph.NodeStatus
	nodeMap := make(map[Node]bool)
	for _, node := range newGraph.MapInstCtxKey2Node {
		nodeMap[node] = true
	}
	for _, selectCases := range newGraph.Select2Case {
		for _, select_case := range selectCases {
			nodeMap[select_case] = true
		}
	}
	for _, nodes := range newGraph.MapInstCtxKey2Defer {
		for _, node := range nodes {
			nodeMap[node] = true
		}
	}

	for node, _ := range newGraph.NodeStatus {
		_, found := nodeMap[node]
		if !found {
			fmt.Print("Found a node in status but not in 3 maps")
		}
	}
}

const (
	chan_make  = 0
	chan_send  = 1
	chan_recv  = 2
	chan_close = 3
	lock       = 4
	unlock     = 5
	rlock      = 6
	runlock    = 7
)

func opType(op SyncOp) int {
	var intType int
	switch concreteType := op.(type) {
	case *ChanMake:
		intType = chan_make
	case *ChanOp:
		switch concreteType.Op.(type) {
		case *instinfo.ChSend:
			intType = chan_send
		case *instinfo.ChRecv:
			intType = chan_recv
		case *instinfo.ChClose:
			intType = chan_close
		}
	case *SelectCase:
		switch concreteType.Op.(type) {
		case *instinfo.ChSend:
			intType = chan_send
		case *instinfo.ChRecv:
			intType = chan_recv
		case *instinfo.ChClose:
			intType = chan_close
		}
	case *LockerOp:
		switch lock_op := concreteType.Op.(type) {
		case *instinfo.LockOp:
			switch lock_op.IsRLock {
			case true:
				intType = rlock
			case false:
				intType = lock
			}
		case *instinfo.UnlockOp:
			switch lock_op.IsRUnlock {
			case true:
				intType = runlock
			case false:
				intType = unlock
			}
		}
	}
	return intType
}

// See if two operations can synchronize
func twoTypesCanSync(aType, bType int) bool {

	switch aType {
	case chan_make:
		return false
		// do nothing
	case chan_send, chan_close:
		if bType == chan_recv {
			return true
		}
	case chan_recv:
		switch bType {
		case chan_send, chan_close:
			return true
		}
	case lock:
		if bType == unlock {
			return true
		}
	case unlock:
		if bType == lock {
			return true
		}
	case rlock:
		if bType == runlock {
			return true
		}
	case runlock:
		if bType == rlock {
			return true
		}
	}
	return false
}

func firstInstOfFn(fn *ssa.Function) ssa.Instruction {
	if fn == nil {
		return nil
	}
	if len(fn.Blocks) == 0 {
		return nil
	}
	bb := fn.Blocks[0]
	if len(bb.Instrs) == 0 {
		return nil
	}
	return bb.Instrs[0]
}

func isSpecialPrim(prim interface{}) bool {
	if prim == &instinfo.ChanTimer || prim == &instinfo.ChanContext {
		return true
	}
	return false
}

type WalkConfig struct {
	NoBackedge bool
	NoCallee   bool
	EntryFn    func(Node)
	EdgeFn     func(*NodeEdge) // called when meet an edge, before deciding whether to skip edge according to NoBackedge or NoCallee
	ExitFn     func(Node)
}

func Walk(node Node, cfg *WalkConfig) {

	if cfg.EntryFn != nil {
		cfg.EntryFn(node)
	}

	if cfg.ExitFn != nil {
		defer cfg.ExitFn(node)
	}

	for _, out := range node.Out() {
		if cfg.EdgeFn != nil {
			cfg.EdgeFn(out)
		}
		if (out.IsGo || out.IsCall) && cfg.NoCallee {
			continue
		}
		if out.IsBackedge && cfg.NoBackedge {
			continue
		}
		Walk(out.Succ, cfg)
	}
}

var intNextId int

// Walk through node and its succs, to update node.In, node.Out and node.Num_Paths. Don't walk through backedge.
// Don't walk into callee, just append them to head_of_fn
func WalkGenEdge(node Node, mapHeadOfFn map[Node]bool) {

	if len(node.Out()) > 0 { // This node is walked before. Note that this can still happen even if we don't walk backedge
		return
	}

	node.SetId(intNextId)
	intNextId++

	str := node.Instruction().String()
	if value, ok := node.Instruction().(ssa.Value); ok {
		str = value.Name() + str
	}
	node.SetString(str)

	switch concrete := node.(type) {
	case *Jump:
		edge := &NodeEdge{
			Prev:       node,
			Succ:       concrete.Next,
			IsBackedge: concrete.BoolIsBackedge,
			IsCall:     false,
			IsGo:       false,
			AddValue:   0,
		}
		node.OutAdd(edge)
		concrete.Next.InAdd(edge)
		if concrete.BoolIsBackedge == false {
			WalkGenEdge(concrete.Next, mapHeadOfFn)
		}
		return

	case *If:
		edge1 := &NodeEdge{
			Prev:       node,
			Succ:       concrete.Then,
			IsBackedge: concrete.BoolIsThenBackedge,
			IsCall:     false,
			IsGo:       false,
			AddValue:   0,
		}
		edge2 := &NodeEdge{
			Prev:       node,
			Succ:       concrete.Else,
			IsBackedge: concrete.BoolIsElseBackedge,
			IsCall:     false,
			IsGo:       false,
			AddValue:   0,
		}
		node.OutAdd(edge1)
		node.OutAdd(edge2)
		concrete.Then.InAdd(edge1)
		concrete.Else.InAdd(edge2)
		if concrete.BoolIsThenBackedge == false {
			WalkGenEdge(concrete.Then, mapHeadOfFn)
		}
		if concrete.BoolIsElseBackedge == false {
			WalkGenEdge(concrete.Else, mapHeadOfFn)
		}
		return

	case *Call:
		for _, callee := range concrete.Calling {
			if callee == nil {
				continue
			}
			edge := &NodeEdge{
				Prev:       node,
				Succ:       callee,
				IsBackedge: false,
				IsCall:     true,
				IsGo:       false,
				AddValue:   0,
			}
			node.OutAdd(edge)
			callee.InAdd(edge)

			mapHeadOfFn[callee] = true
		}

	case *End, *Overwrite: // In current implementation, these are not in the graph
		return
	case *Return:
		return
	case *Kill:
		if concrete.Next == nil { // No defer executed
			return
		}
		edge := &NodeEdge{
			Prev:       node,
			Succ:       concrete.Next,
			IsBackedge: false,
			IsCall:     false,
			IsGo:       false,
			AddValue:   0,
		}
		node.OutAdd(edge)
		concrete.Next.InAdd(edge)
		WalkGenEdge(concrete.Next, mapHeadOfFn)
		return
	case *Go:
		for _, callee := range concrete.MapCreateNodes {
			if callee == nil {
				continue
			}
			edge := &NodeEdge{
				Prev:       node,
				Succ:       callee,
				IsBackedge: false,
				IsCall:     false,
				IsGo:       true,
				AddValue:   0,
			}
			node.OutAdd(edge)
			callee.InAdd(edge)

			mapHeadOfFn[callee] = true
		}

	case *Select:
		for _, select_case := range concrete.Cases {
			edge := &NodeEdge{
				Prev:       node,
				Succ:       select_case,
				IsBackedge: false,
				IsCall:     false,
				IsGo:       false,
				AddValue:   0,
			}
			node.OutAdd(edge)
			select_case.InAdd(edge)
			WalkGenEdge(select_case, mapHeadOfFn)
		}

		return
	case *SelectCase:
		edge := &NodeEdge{
			Prev:       node,
			Succ:       concrete.Next,
			IsBackedge: concrete.BoolIsBackedge,
			IsCall:     false,
			IsGo:       false,
			AddValue:   0,
		}
		node.OutAdd(edge)
		concrete.Next.InAdd(edge)
		if concrete.BoolIsBackedge == false {
			WalkGenEdge(concrete.Next, mapHeadOfFn)
		}
		return
	}

	var next Node
	switch concrete := node.(type) {
	case *Call:
		next = concrete.NextLocal
	case *Go:
		next = concrete.NextLocal
	case *ChanMake:
		next = concrete.Next
	case *ChanOp:
		next = concrete.Next
	case *LockerOp:
		next = concrete.Next
	case *NormalInst:
		next = concrete.Next
	default:
		return
	}

	edge := &NodeEdge{
		Prev:       node,
		Succ:       next,
		IsBackedge: false,
		IsCall:     false,
		IsGo:       false,
		AddValue:   0,
	}

	node.OutAdd(edge)
	node.InAdd(edge)

	WalkGenEdge(next, mapHeadOfFn)
}

// This function updates Node.Num_paths and NodeEdge.AddValue
func (g *SyncGraph) BuildNodeInOut() {

	headOfFn := make(map[Node]bool)
	for _, headGoroutine := range g.HeadGoroutines {
		headOfFn[headGoroutine.HeadNode] = true
	}
	intNextId = 0
	for len(headOfFn) > 0 { // for the head of every function
		var next_head Node
		for head, ok := range headOfFn {
			if ok {
				next_head = head
				break
			}
		}
		WalkGenEdge(next_head, headOfFn)

		//Build_DAG(next_head)

		delete(headOfFn, next_head)
	}
}

func TypeMsgForNode(node Node) string {
	strNodeType := ""
	switch node.(type) {
	case *Jump:
		strNodeType = "Jump"
	case *If:
		strNodeType = "If"
	case *Call:
		strNodeType = "Call"
	case *End: // In current implementation, End is not in the graph
		strNodeType = "End"
	case *Overwrite: // In current implementation, Overwrite is not in the graph
		strNodeType = "Overwrite"
	case *Return:
		strNodeType = "Return"
	case *Go:
		strNodeType = "Go"
	case *ChanMake:
		strNodeType = "ChanMake"
	case *Select:
		strNodeType = "Select"
	case *SelectCase:
		strNodeType = "Select_case"
	case *ChanOp:
		strNodeType = "Chan_op"
	case *LockerOp:
		strNodeType = "Locker_op"
	case *NormalInst:
		strNodeType = "Normal_inst"
	case *Kill:
		strNodeType = "Kill"
	case nil:
		strNodeType = "Not Built"
	}
	return strNodeType
}

func (p *PPath) PrintPPath() {
	const tick = '\u2713'
	const cross = '\u2717'
	for _, pn := range p.Path {
		strTypeMsg := TypeMsgForNode(pn.Node)
		if strTypeMsg == "Normal_inst" {
			continue
		}
		fmt.Print(strTypeMsg)
		fmt.Print(" :", output.GetLoc(pn.Node.Instruction()))
		if pn.Blocked {
			fmt.Print("\t Blocking")
		} else if pn.Executed {
			fmt.Printf("\t %q", tick)
		} else {
			fmt.Printf("\t %q", cross)
		}
		fmt.Println()
	}
}

func (g *SyncGraph) PrintAllPathCombinations() {
	count := 0
	for _, tuple := range g.PathCombinations {
		fmt.Println("=======combination NO.", count)
		count++
		for _, goroutinePath := range tuple.go_paths {
			if goroutinePath.goroutine.EntryFn != nil {
				fmt.Println("-----Goroutine:", goroutinePath.goroutine.EntryFn.String())
			} else {
				fmt.Println("-----Goroutine: Entry")
			}
			intLastLine := -1
			for _, node := range goroutinePath.path.Path {
				strNodeType := TypeMsgForNode(node)
				if strNodeType == "Normal_inst" {
					return
				}
				p := config.Prog.Fset.Position(node.Instruction().Pos())
				if p.Line == intLastLine {

				} else {
					fmt.Print(p.Line)
					fmt.Print("\t", strNodeType)

					fmt.Print(" :", output.GetLoc(node.Instruction()))
					fmt.Println()
				}

			}
		}
	}
	output.WaitForInput()
}

// Walk the whole graph, print type of node on path
func (g *SyncGraph) PrintGraphAllNodesType() {
	count := 0
	vecAllHeads := []Node{}
	for _, headGoroutine := range g.HeadGoroutines {
		vecAllHeads = append(vecAllHeads, headGoroutine.HeadNode)
	}

	intLastLine := -1
	entryFn := func(node Node) {
		strNodeType := TypeMsgForNode(node)
		if strNodeType == "Normal_inst" {
			return
		}

		p := config.Prog.Fset.Position(node.Instruction().Pos())
		if intLastLine == p.Line && intLastLine != 0 {

		} else {
			intLastLine = p.Line
			fmt.Print(p.Line)
			fmt.Print("\t", strNodeType)
			fmt.Print(" :", output.GetLoc(node.Instruction()))
			fmt.Println()
		}
		count++
		for _, out := range node.Out() {
			if out.IsCall || out.IsGo {
				if out.Succ == nil {
					fmt.Println("Warning in PrintGraphAllNodesType: an out edge that is Call or Go has nil succ")
					output.WaitForInput()
				}
				vecAllHeads = append(vecAllHeads, out.Succ)
			}
		}
	}

	exit_fn := func(node Node) {
		if TypeMsgForNode(node) == "Normal_inst" {
			return
		}
		fmt.Println("\t\tMove Back")
	}

	cfg := &WalkConfig{
		NoBackedge: true,
		NoCallee:   true,
		EntryFn:    entryFn,
		EdgeFn:     nil,
		ExitFn:     exit_fn,
	}

	for len(vecAllHeads) > 0 {
		this_head := vecAllHeads[0]
		fmt.Print("--------Printing the graph of: ")
		callchain := this_head.CallCtx().CallChain
		fmt.Print(callchain.Start.Func.Name())
		for _, edge := range callchain.Chain {
			fmt.Print(" ---> ", edge.Callee.Func.Name())
		}
		fmt.Print("\n\n")
		count = 0
		Walk(this_head, cfg)
		vecAllHeads = removeFromHeadList(vecAllHeads, this_head)
	}

	output.WaitForInput()
}

func removeFromHeadList(old []Node, delete Node) []Node {
	result := []Node{}
	for _, o := range old {
		if o == delete {
			continue
		}
		result = append(result, o)
	}
	return result
}

func isKillThread(inst ssa.Instruction) bool {
	call, ok := inst.(ssa.CallInstruction)
	if !ok {
		return false
	}
	return isFatal(call)
}

func isFatal(call ssa.CallInstruction) bool {

	if call.Common().IsInvoke() {
		return false
	}
	fn, ok := call.Common().Value.(*ssa.Function)
	if !ok {
		return false
	}
	if fn.Pkg == nil || fn.Pkg.Pkg == nil {
		return false
	}

	fnName := fn.Name()
	pkgName := fn.Pkg.Pkg.Name()
	//pkg_path := fn.Pkg.Pkg.Path()

	var listFatal []string

	if pkgName == "testing" {
		listFatal = []string{"Fatal", "Fatalf", "FailNow"} // "Skip","Skipf","SkipNow" are not considered here, because we want to check test functions that start with a Skip

	} else if pkgName == "assert" { // github.com/stretchr/testify/assert
		listFatal = []string{"Fail"}

		//} else if pkg_path == "github.com/cockroachdb/cockroach/pkg/testutils"{
		//	listFatal = []string{"SucceedsSoon"}
	} else {
		return false
	}
	for _, fatal := range listFatal {
		if fatal == fnName {
			return true
		}
	}
	return false
}

func canSync(op1, op2 SyncOp) bool {
	switch op1Concrete := op1.(type) {
	case *ChanOp:
		op2Concrete, isOp2Chan := op2.(*ChanOp)
		if isOp2Chan == false {
			return false
		}
		switch op1Concrete.Op.(type) {
		case *instinfo.ChSend:
			_, isOp2Recv := op2Concrete.Op.(*instinfo.ChRecv)
			if isOp2Recv {
				return true
			} else {
				return false
			}
		case *instinfo.ChRecv:
			_, isOp2Send := op2Concrete.Op.(*instinfo.ChSend)
			if isOp2Send {
				return true
			} else {
				return false
			}
		}
	case *LockerOp:
		op2Concrete, isOp2Locker := op2.(*LockerOp)
		if isOp2Locker == false {
			return false
		}
		switch op1Concrete.Op.(type) {
		case *instinfo.LockOp:
			_, isOp2UnLock := op2Concrete.Op.(*instinfo.UnlockOp)
			if isOp2UnLock {
				return true
			} else {
				return false
			}
		case *instinfo.UnlockOp:
			_, isOp2Lock := op2Concrete.Op.(*instinfo.LockOp)
			if isOp2Lock {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func canSyncOpTriggerGl(op SyncOp) bool {
	switch concrete := op.(type) {
	case *SelectCase:
		return false // Select can't be the blocking operation of GL, because select involves other primitives,
		// and if they are blocking, the bug is circular wait bug
	case *ChanMake:
		return false
	case *ChanOp: // close, recv, send. Not in select
		switch concrete.Op.(type) {
		case *instinfo.ChClose:
			return false
		default:
			return true
		}
	case *LockerOp: // lock, unlock
		switch concrete.Op.(type) {
		case *instinfo.LockOp:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

// send should be equal or less than recv+buffer. If no close, recv should be equal to or less than send
func checkChOpsLegal(ch *instinfo.Channel, ops []Node) bool {
	var intNumSend, intNumRecv, intNumClose int
	for _, op := range ops {
		switch concrete_node := op.(type) {
		case *ChanMake:
		case *ChanOp:
			switch concrete_node.Op.(type) {
			case *instinfo.ChSend:
				intNumSend++
			case *instinfo.ChClose:
				intNumClose++
			case *instinfo.ChRecv:
				intNumRecv++
			}
		case *SelectCase:
			switch concrete_node.Op.(type) {
			case *instinfo.ChSend:
				intNumSend++
			case *instinfo.ChRecv:
				intNumRecv++
			}
		}
	}

	if intNumSend > 0 {
		if ch.Buffer == instinfo.DynamicSize { // Todo: we can try to analyze the real value of buffer

		} else {
			if intNumSend > intNumRecv+ch.Buffer {
				return false
			}
		}
	}

	if intNumRecv > 0 {
		if intNumClose > 0 { // then intNumRecv can be any value

		} else {
			if intNumRecv > intNumSend {
				return false
			}
		}
	}

	return true
}

// lock should be same as unlock, or just 1 above unlock
func checkLockerOpsLegal(l *instinfo.Locker, ops []Node) bool {
	var intNumLock, intNumUnlock int
	for _, op := range ops {
		switch concrete_node := op.(type) {
		case *LockerOp:
			switch concrete_node.Op.(type) {
			case *instinfo.LockOp:
				intNumLock++
			case *instinfo.UnlockOp:
				intNumUnlock++
			}
		}
	}

	if intNumUnlock == intNumLock {
		return true
	} else if intNumUnlock > intNumLock {
		return false
	} else { // intNumUnlock < intNumLock
		if intNumUnlock == intNumLock - 1 {
			return true
		} else {
			return false
		}
	}
}

func fnsForInstsNoDupli(insts []ssa.Instruction) []*ssa.Function {
	result := []*ssa.Function{}

	mapFns := make(map[*ssa.Function]struct{})
	for _, inst := range insts {
		mapFns[inst.Parent()] = struct{}{}
	}

	for fn, _ := range mapFns {
		result = append(result, fn)
	}

	return result
}

func (g *SyncGraph) syncOpsOfTargetChans() []SyncOp {
	vecOpsOfTarget := []SyncOp{}
	for _, tPrim := range g.Task.VecTaskPrimitive {
		if _, ok := tPrim.Primitive.(*instinfo.Channel); !ok {
			continue
		}
		ops := g.MapPrim2VecSyncOp[tPrim.Primitive]
		for _, op := range ops {
			vecOpsOfTarget = append(vecOpsOfTarget, op)
		}
	}
	return vecOpsOfTarget
}

func isSyncOpInSlice(target SyncOp, slice []SyncOp) bool {
	for _, op := range slice {
		if op == target {
			return true
		}
	}
	return false
}

func (g *SyncGraph) findUnfinishGoOfMainGoroutine(unfinishes []*Unfinish) *Unfinish {
	for _, unfinished := range unfinishes {
		caller := unfinished.Site.Caller.Func
		//callee := unfinished.Site.Callee.Func
		if unfinished.IsGo && caller == g.MainGoroutine.EntryFn {
			//if is_b_anonymous_in_a(caller, callee) {
			//
			//}
			return unfinished
		}
	}
	return nil
}

func isBBSliceEqual(s1, s2 []*ssa.BasicBlock) bool {
	map1 := make(map[*ssa.BasicBlock]struct{})
	map2 := make(map[*ssa.BasicBlock]struct{})
	for _, bb := range s1 {
		map1[bb] = struct{}{}
	}
	for _, bb := range s2 {
		map2[bb] = struct{}{}
	}
	if len(map1) != len(map2) {
		return false
	}
	for bb, _ := range map1 {
		_, existInMap2 := map2[bb]
		if existInMap2 == false {
			return false
		}
	}
	return true
}
