package syncgraph

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/analysis"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/util"
	"go/constant"
	"go/token"
	"strconv"
	"strings"
	"time"
)

type tupleGoroutinePath struct {
	goroutine *Goroutine
	path      *LocalPath
}

type pathCombination struct { // []'s length is len(g.Goroutines). Every tupleGoroutinePath.Goroutine is unique
	go_paths []*tupleGoroutinePath
}

type EnumeConfigure struct {
	Unfold             int
	IgnoreFn           map[*ssa.Function]struct{} // a map of not interesting
	FlagIgnoreNoSyncFn bool
	FlagIgnoreNormal   bool
}

type LocalPath struct {
	Path                   []Node
	Hash                   string
	mapNodeEdge2IntVisited map[*NodeEdge]int
	mapLoopHead2Visited    map[Node]int
	//finished bool
}

var intEmptyPathId int

func NewEmptyPath() *LocalPath {
	intEmptyPathId++
	return &LocalPath{
		Path:                   []Node{},
		Hash:                   "Empty_path_NO." + strconv.Itoa(intEmptyPathId),
		mapNodeEdge2IntVisited: nil,
		mapLoopHead2Visited:    nil,
	}
}

func (l *LocalPath) IsEmpty() bool {
	return strings.HasPrefix(l.Hash, "Empty_path_NO.")
}

type PNode struct { // A Node will have n PNode if it shows up n times in path
	Path     *LocalPath
	Index    int
	Node     Node
	Blocked  bool
	Executed bool
}

type PPath struct {
	Path      []*PNode
	localPath *LocalPath
}

func (p PPath) IsNodeIn(node Node) bool {
	for _, t_n := range p.Path {
		if t_n.Node == node {
			return true
		}
	}
	return false
}

func (p PPath) SetAllReached() {
	for _, n := range p.Path {
		n.Blocked = false
		n.Executed = true
	}
}

func (p PPath) SetBlockAt(index int) {
	for i, n := range p.Path {
		if i < index {
			n.Executed = true
			n.Blocked = false
		} else if i == index {
			n.Executed = false
			n.Blocked = true
		} else {
			n.Executed = false
			n.Blocked = false
		}
	}
}

var mapHash2Map map[string]*LocalPath

type tupleCallerCallee struct {
	caller Node
	callee Node
}

func deleteNormalFromPath(oldPath *LocalPath) *LocalPath {
	newSlice := []Node{}
	for _, node := range oldPath.Path {
		if _, boolIsNormal := node.(*NormalInst); boolIsNormal {
			continue
		}
		newSlice = append(newSlice, node)
	}

	newLocalPath := &LocalPath{
		Path:                   newSlice,
		Hash:                   hashOfPath(newSlice),
		mapNodeEdge2IntVisited: copyBackedgeMap(oldPath.mapNodeEdge2IntVisited),
		mapLoopHead2Visited:    copyHeaderMap(oldPath.mapLoopHead2Visited),
	}
	oldPath = nil
	return newLocalPath
}

func copyLocalPath(old *LocalPath) *LocalPath {
	newPath := make([]Node, len(old.Path))
	copy(newPath, old.Path)
	newLocalPath := &LocalPath{
		Path:                   newPath,
		Hash:                   old.Hash,
		mapNodeEdge2IntVisited: copyBackedgeMap(old.mapNodeEdge2IntVisited),
		mapLoopHead2Visited:    copyHeaderMap(old.mapLoopHead2Visited),
	}
	return newLocalPath
}

func debugPrintEnumeratedPaths(path_map map[string]*LocalPath) {
	count := 0
	fmt.Println("In total:", len(path_map))
	for _, path := range path_map {
		fmt.Println("-----Path:", count)
		//fmt.Println(path.mapNodeEdge2IntVisited)
		fmt.Println(path.mapLoopHead2Visited)
		count++
		for i, n := range path.Path {
			str := TypeMsgForNode(n)
			if str == "Normal_inst" {
				continue
			}
			fmt.Println(str)
			if i < len(path.Path)-1 {
				var flag_backedge bool
				next := path.Path[i+1]
				for _, out := range n.Out() {
					if out.Succ == next {
						flag_backedge = out.IsBackedge
						break
					}
				}
				if flag_backedge {
					fmt.Println("--Backedge")
				}
			}
			if i == len(path.Path)-1 {
				if TypeMsgForNode(n) != "Return" {
					output.WaitForInput()
				}
			}
		}
		output.WaitForInput()
	}
}

func copyBackedgeMap(old map[*NodeEdge]int) map[*NodeEdge]int {
	n := make(map[*NodeEdge]int)
	for key, value := range old {
		n[key] = value
	}
	return n
}

func copyHeaderMap(old map[Node]int) map[Node]int {
	n := make(map[Node]int)
	for key, value := range old {
		n[key] = value
	}
	return n
}

func copyPathMap(old map[string]*LocalPath) map[string]*LocalPath {
	n := make(map[string]*LocalPath)
	for key, value := range old {
		n[key] = value
	}
	return n
}

func copyPathSlice(old []Node) []Node {
	copy_path := []Node{}
	for _, n := range old {
		copy_path = append(copy_path, n)
	}
	return copy_path
}

func copyIntSlice(old []int) []int {
	n := []int{}
	for _, o := range old {
		n = append(n, o)
	}
	return n
}

func hashOfPath(node []Node) string {
	var buffer bytes.Buffer
	for _, n := range node {
		str := n.GetString() + "_"
		buffer.WriteString(str)
	}
	byte_key := buffer.Bytes()
	hash := sha256.Sum256(byte_key)
	return string(hash[:])
}

func removeFromPathWorklist(old_worklist []*LocalPath, remove *LocalPath) []*LocalPath {
	result := []*LocalPath{}
	for _, o := range old_worklist {
		if o == remove {
			continue
		}
		result = append(result, o)
	}
	return result
}

func (g *SyncGraph) SetEnumCfg(unfold int, flagIgnoreNoSyncFn bool, flagIgnoreNormal bool) {

	ignoreFn := make(map[*ssa.Function]struct{})
	if flagIgnoreNoSyncFn {
		// list some functions containing no sync_op, kill, Go

		vecAllHeads := make(map[Node]struct{})
		for _, goroutine := range g.Goroutines {
			vecAllHeads[goroutine.HeadNode] = struct{}{}
		}

		var flagNotInteresting *bool // need to overwrite at the beginning of entering a function
		entryFn := func(node Node) {
			switch node.(type) {
			case SyncOp, *Kill, *Go:
				*flagNotInteresting = false
			}
		}

		edgeFn := func(edge *NodeEdge) {
			if edge.IsBackedge || edge.IsGo {
				return
			} else if edge.IsCall {
				vecAllHeads[edge.Succ] = struct{}{}
			}
		}

		walkCfg := &WalkConfig{
			NoBackedge: true,
			NoCallee:   true,
			EntryFn:    entryFn,
			EdgeFn:     edgeFn,
			ExitFn:     nil,
		}

		for len(vecAllHeads) > 0 {
			var thisHead Node
			for node, _ := range vecAllHeads {
				thisHead = node
				break
			}
			*flagNotInteresting = true
			Walk(thisHead, walkCfg)
			if *flagNotInteresting == true { // walked all nodes of this function, and no
				if thisHead.Instruction() != nil { // this should always be true
					ignoreFn[thisHead.Instruction().Parent()] = struct{}{}
				}
			}

			delete(vecAllHeads, thisHead)
		}
	}

	cfg := &EnumeConfigure{
		Unfold:             unfold,
		IgnoreFn:           ignoreFn,
		FlagIgnoreNoSyncFn: flagIgnoreNoSyncFn,
		FlagIgnoreNormal:   flagIgnoreNormal,
	}
	g.EnumerateCfg = cfg
}

// Enumerate all path combinations. A pathCombination: a slice of {one goroutine and one path}. Path can be Not_execute, meaning empty path.
// Note that the number of goroutine_path may be greater than len(g.Goroutines). For example, if a *Go is visited 3 times, we create 3 goroutines
func (g *SyncGraph) EnumerateAllPathCombinations() {

	g.PathCombinations = []*pathCombination{}

	possiblePaths := []map[int]*LocalPath{} // int is path index:0,1,2...   Index 0 is always Not_executed
	for i, goroutine := range g.Goroutines {
		possiblePaths = append(possiblePaths, make(map[int]*LocalPath))

		count := 0
		possiblePaths[i][count] = NewEmptyPath()
		count++

		goroutinePathMap := EnumeratePathWithGoroutineHead(goroutine.HeadNode, g.EnumerateCfg)
		if goroutinePathMap == nil {
			return
		}
		//util.Debugfln("goroutine = %s", goroutine.EntryFn)
		//for _, path := range goroutinePathMap {
		//	for _, node := range path.Path {
		//		util.Debugfln("\t node = %s", node.Instruction())
		//	}
		//}

		for _, path := range goroutinePathMap {
			possiblePaths[i][count] = path
			count++
		}
	}

	n := len(g.Goroutines)

	indices := []int{}
	for i := 0; i < n; i++ {
		indices = append(indices, 0)
	}

	for {
		if len(g.PathCombinations) > config.Max_PATH_ENUMERATE {
			if config.Print_Debug_Info {
				fmt.Println("!!!!")
				fmt.Println("EnumerateAllPathCombinations: reached max enumerate number")
			}
			return
		}

		// store the current combination. During store, do some checks, and may add extra goroutine and path
		goroutines := []*Goroutine{}
		paths := []*LocalPath{}
		for i := 0; i < n; i++ {
			goroutine := g.Goroutines[i]
			goroutines = append(goroutines, goroutine)
			path := possiblePaths[i][indices[i]]
			paths = append(paths, path)
		}
		g.generateNewPathCombinations(goroutines, paths, possiblePaths) // in this function, g.PathCombinations will be added

		// find the rightmost array that has more elements left after the current element in that array
		next := n - 1
		for next >= 0 && (indices[next]+1 >= len(possiblePaths[next])) {
			next -= 1
		}

		// no such array is found so no more combinations left
		if next < 0 {
			break
		}

		// if found move to next element in that array
		indices[next] += 1

		// for all arrays to the right of this array, current index again points to the first element
		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}

// Enumerate all possible paths of the given goroutine. If goroutine starts at function A, A has 3 paths, and in 2 paths
// A calls B, and B has 4 paths. Then we should return (1 + 2 * 4) paths, where B's path is inserted into A's path.
func EnumeratePathWithGoroutineHead(head Node, enumeConfigure *EnumeConfigure) map[string]*LocalPath {

	fakeCaller := Fake_Node()

	// A map from the caller Node to all paths of its callee
	caller2paths := make(map[Node][]*LocalPath)

	// A worklist from head to nodes of callee, all in the same goroutine
	todoFnHeads := make(map[tupleCallerCallee]struct{})
	todoFnHeads[tupleCallerCallee{
		caller: fakeCaller,
		callee: head,
	}] = struct{}{}

	startEnumeAllPaths := time.Now()

	for len(todoFnHeads) > 0 {
		if false || time.Since(startEnumeAllPaths) > config.MAX_PATH_ENUMERATE_SECOND*time.Second {
			if config.Print_Debug_Info {
				fmt.Println("!!!!")
				fmt.Println("Warning in EnumeratePathWithGoroutineHead: timeout")
			}
			return nil
		}
		var thisCallerCallee tupleCallerCallee
		for tupleCallerCallee, _ := range todoFnHeads { // get the first tupleCallerCallee in todo_list. Order doesn't matter
			thisCallerCallee = tupleCallerCallee
			break
		}

		mapHash2Map = make(map[string]*LocalPath)
		enumeratePathBreadthFirst(thisCallerCallee.callee, enumeConfigure.Unfold, todoFnHeads) // updates mapHash2Map
		if enumeConfigure.FlagIgnoreNoSyncFn {                                                 // if callee is a not interesting function (see SetEnumCfg),
			// then reserve only Call Nodes of its paths
			calleeFn := thisCallerCallee.callee.Instruction().Parent()
			if _, ok := enumeConfigure.IgnoreFn[calleeFn]; ok {
				fmt.Println("regen_only_Call_path_map is not complete!")
			}
		}
		for _, path := range mapHash2Map {
			util.Debugfln("path in mapHash2Map: %s", path)
			if enumeConfigure.FlagIgnoreNormal {
				path = deleteNormalFromPath(path)
			}
			if notCorrectUnroll(path) {
				continue
			}
			caller2paths[thisCallerCallee.caller] = append(caller2paths[thisCallerCallee.caller], copyLocalPath(path))
		}

		delete(todoFnHeads, thisCallerCallee)
	}
	for _, path := range caller2paths {
		util.Debugfln("path in caller2paths: %s", path)
	}
	mapHash2Map = nil
	todoFnHeads = nil // Not useful anymore

	// Now we have caller2paths map. Note that we want all paths for a goroutine. If A calls B, and A has 3 paths,
	// B have 4 paths, we should return 12 paths.

	// Time to generate the return value: complete paths that callee's path is inserted into caller's path
	result := make(map[string]*LocalPath)

	// a worklist of incomplete paths
	type unfinishPath struct {
		path                  []Node
		hash                  string
		mapVisitedBackedge    map[*NodeEdge]int
		mapLoopHeadVisited    map[Node]int
		vecUnfinishCallsIndex []int // a list of indexs of *Call Nodes that have callee in caller2paths map,
		// but the callee's path has not been inserted yet. We must use index instead of Node, because Node can duplicate
	}
	worklistPaths := []unfinishPath{}
	// prepare worklistPaths: add paths in entry fn of this goroutine into this list
	for caller, paths := range caller2paths {
		if caller == fakeCaller { // now paths is of the entry function of this goroutine
			for _, path := range paths {
				unfinish_calls_index := []int{}
				for index, node := range path.Path {
					_, ok := caller2paths[node]
					if ok { // this node has callee, and callee has some paths
						unfinish_calls_index = append(unfinish_calls_index, index)
					}
				}

				newPath := make([]Node, len(path.Path))
				copy(newPath, path.Path)
				newUnfinish := unfinishPath{
					path:                  newPath,
					hash:                  path.Hash,
					mapVisitedBackedge:    copyBackedgeMap(path.mapNodeEdge2IntVisited),
					mapLoopHeadVisited:    copyHeaderMap(path.mapLoopHead2Visited),
					vecUnfinishCallsIndex: copyIntSlice(unfinish_calls_index),
				}
				worklistPaths = append(worklistPaths, newUnfinish)
			}
		}
	}

	for len(worklistPaths) > 0 {
		thisUnfinish := worklistPaths[0]
		if len(thisUnfinish.vecUnfinishCallsIndex) == 0 {
			newLocalPath := &LocalPath{
				Path:                   thisUnfinish.path,
				Hash:                   thisUnfinish.hash,
				mapNodeEdge2IntVisited: thisUnfinish.mapVisitedBackedge,
				mapLoopHead2Visited:    thisUnfinish.mapLoopHeadVisited,
			}
			result[thisUnfinish.hash] = newLocalPath
		}
		for _, callIndex := range thisUnfinish.vecUnfinishCallsIndex {
			caller := thisUnfinish.path[callIndex]
			calleePaths := caller2paths[caller]
			for _, calleePath := range calleePaths {
				vecCalleeUnfinishCallIndexs := []int{}
				for index, node := range calleePath.Path {
					_, ok := caller2paths[node]
					if ok { // this node has callee, and callee has some paths
						vecCalleeUnfinishCallIndexs = append(vecCalleeUnfinishCallIndexs, index)
					}
				}

				vecNewUnfinishCallIndexs := []int{}
				for _, callerUnfinishIndex := range thisUnfinish.vecUnfinishCallsIndex {
					if callerUnfinishIndex < callIndex { // reserve this unfinish_index
						vecNewUnfinishCallIndexs = append(vecNewUnfinishCallIndexs, callerUnfinishIndex)
					} else if callerUnfinishIndex == callIndex { // insert vecCalleeUnfinishCallIndexs. Skip callIndex
						for _, callee_unfinish_call_index := range vecCalleeUnfinishCallIndexs {
							vecNewUnfinishCallIndexs = append(vecNewUnfinishCallIndexs, callee_unfinish_call_index+callIndex+1)
						}
					} else { // reserve this unfinish_index, but increase len(calleePath)
						vecNewUnfinishCallIndexs = append(vecNewUnfinishCallIndexs, callerUnfinishIndex+len(calleePath.Path))
					}
				}

				newPath := []Node{}
				for i, caller_node := range thisUnfinish.path {
					newPath = append(newPath, caller_node)
					if i == callIndex {
						for _, callee_node := range calleePath.Path {
							newPath = append(newPath, callee_node)
						}
					}
				}

				newHash := hashOfPath(newPath)

				combinedBackedgeMap := make(map[*NodeEdge]int)
				for key, value := range thisUnfinish.mapVisitedBackedge {
					combinedBackedgeMap[key] = value
				}
				for key, value := range calleePath.mapNodeEdge2IntVisited {
					combinedBackedgeMap[key] = value
				}

				if len(vecNewUnfinishCallIndexs) == 0 { // no unfinished nodes, this can be added to return value

					newLocalPath := &LocalPath{
						Path:                   newPath,
						Hash:                   newHash,
						mapNodeEdge2IntVisited: combinedBackedgeMap,
					}
					result[newHash] = newLocalPath
				} else {
					newUnfinish := unfinishPath{
						path:                  newPath,
						hash:                  newHash,
						mapVisitedBackedge:    combinedBackedgeMap,
						vecUnfinishCallsIndex: vecNewUnfinishCallIndexs,
					}
					worklistPaths = append(worklistPaths, newUnfinish)
				}
			}
		}

		newList := []unfinishPath{}
		for _, oldUnfinish := range worklistPaths {
			if oldUnfinish.hash == thisUnfinish.hash {
				continue
			}
			newList = append(newList, oldUnfinish)
		}
		worklistPaths = newList
	}

	return result
}

func enumeratePathBreadthFirst(head Node, LoopUnfoldBound int, todo_fn_heads map[tupleCallerCallee]struct{}) {
	worklist := []*LocalPath{}

	head_path := []Node{head}
	hash_head_path := hashOfPath(head_path)
	head_local_path := &LocalPath{
		Path:                   head_path,
		Hash:                   hash_head_path,
		mapNodeEdge2IntVisited: make(map[*NodeEdge]int),
		mapLoopHead2Visited:    make(map[Node]int),
	}

	worklist = append(worklist, head_local_path)

	fn := head.Instruction().Parent()
	loopAnalysis := analysis.NewLoopAnalysis(fn)
	mapBackedge2Visited := make(map[*analysis.Edge]int)
	mapLoopHeader2Visited := make(map[*ssa.BasicBlock]int)
	for _, edge := range loopAnalysis.VecBackedge {
		mapBackedge2Visited[edge] = 0
	}
	for headerBB, _ := range loopAnalysis.MapLoopHead2BodyBB {
		mapLoopHeader2Visited[headerBB] = 0
	}

	count := 0
	startPathEnume := time.Now()

	for len(worklist) != 0 {
		count++
		if count > config.Max_PATH_ENUMERATE {
			if config.Print_Debug_Info {
				fmt.Println("!!!!")
				fmt.Println("Warning in enumeratePathBreadthFirst: reached max enumerate number")
			}
			return
		}

		current_local_path := worklist[0]
		worklist = worklist[1:]

		current_path := current_local_path.Path
		last_node := current_path[len(current_path)-1]

		valid_outs := []*NodeEdge{}
		for _, out := range last_node.Out() {
			if out.IsGo {
				continue
			}
			if out.IsCall { // If we encounter call to a Node in another function, add it to todo_list
				calleeFn := out.Succ.Instruction().Parent()
				if _, ok := head.Parent().MapFnOnOpPath[calleeFn]; ok {
					new_caller_callee := tupleCallerCallee{
						caller: last_node,
						callee: out.Succ,
					}
					todo_fn_heads[new_caller_callee] = struct{}{}
					continue
				}
			}
			valid_outs = append(valid_outs, out)
		}
		if len(valid_outs) == 0 {
			util.Debugfln("fn = %s, valid_outs = %s, len = %d", fn.Name(), valid_outs, len(valid_outs))
			util.Debugfln("path = %s", current_local_path)
			//current_local_path.finished = true
			if _, ok := mapHash2Map[current_local_path.Hash]; ok {
				util.Debugfln("update existing path hash: %s", current_local_path.Hash)
			}
			mapHash2Map[current_local_path.Hash] = current_local_path
			continue
		}

	outLoop:
		for _, out := range valid_outs {
			if false && time.Since(startPathEnume) > config.MAX_PATH_ENUMERATE_SECOND*time.Second {
				if config.Print_Debug_Info {
					fmt.Println("Warning in enumeratePathBreadthFirst: timeout")
					for _, prim := range head.Parent().Task.VecTaskPrimitive {
						for op, _ := range prim.Ops {
							if make, ok := op.(*instinfo.ChMake); ok {
								p := config.Prog.Fset.Position(make.Inst.Pos())
								fmt.Println(p.Filename + ":" + strconv.Itoa(p.Line))
								break
							}
						}
					}
				}

				return
			}
			//new_path := copyPathSlice(current_path)
			new_path := make([]Node, len(current_path))
			copy(new_path, current_path)
			new_path = append(new_path, out.Succ)
			hash_new_path := hashOfPath(new_path)
			new_backedge_visited := copyBackedgeMap(current_local_path.mapNodeEdge2IntVisited)
			//newLoopHeaderVisited := copyHeaderMap(current_local_path.mapLoopHead2Visited)

			// update the counter on backedges and loop headers

			// new implementation
			// check if BB has changed
			bbPrev := last_node.Instruction().Block()
			bbSucc := out.Succ.Instruction().Block()
			if bbPrev != bbSucc {
				for backedge, _ := range mapBackedge2Visited {
					if backedge.Pred == bbPrev && backedge.Succ == bbSucc {
						new_backedge_visited[out]++
						if new_backedge_visited[out] > LoopUnfoldBound {
							continue outLoop
						}
					}
				}
			}

			new_local_path := &LocalPath{
				Path:                   new_path,
				Hash:                   hash_new_path,
				mapNodeEdge2IntVisited: new_backedge_visited,
				//finished:         false,
			}

			worklist = append(worklist, new_local_path)
		}
	}

}

// For some path, it contains loop whose iteration number is fixed. Check if such a loop exists and return true if
//the path doesn't unroll it correctly
func notCorrectUnroll(path *LocalPath) bool {
	if path.mapNodeEdge2IntVisited == nil {
		if len(path.Path) != 0 {
			loopAnalysis := analysis.NewLoopAnalysis(path.Path[0].Instruction().Parent())
			for headerBB, _ := range loopAnalysis.MapLoopHead2BodyBB {
				if headerBB_NotCorrectUnroll(headerBB, 0) { // this loop is unrolled 0 times
					return true
				}
			}
		}
		return false
	}
	for backedge, counter := range path.mapNodeEdge2IntVisited {
		// see if the backedge's header block ends with something like "if x < 3"
		//if so, check if the counter == 3
		headerBB := backedge.Succ.Instruction().Block()
		if headerBB_NotCorrectUnroll(headerBB, counter) {
			// detects that a loop's iteration number is different from counter
			return true
		}
	}

	// loop may be unrolled 0 times, then path.mapNodeEdge2IntVisited is nil. Need to find loops first
	if len(path.mapNodeEdge2IntVisited) == 0 && len(path.Path) != 0 {
		loopAnalysis := analysis.NewLoopAnalysis(path.Path[0].Instruction().Parent())
		for headerBB, _ := range loopAnalysis.MapLoopHead2BodyBB {
			if headerBB_NotCorrectUnroll(headerBB, 0) { // this loop is unrolled 0 times
				return true
			}
		}
	}

	return false // we are conservative in this function
}

func headerBB_NotCorrectUnroll(headerBB *ssa.BasicBlock, counter int) bool {
	lastInst := headerBB.Instrs[len(headerBB.Instrs)-1]
	if ssaIf, ok := lastInst.(*ssa.If); ok {
		if ssaBinOp, ok := ssaIf.Cond.(*ssa.BinOp); ok {
			if ssaBinOp.Op == token.LSS {
				if ssaConst, ok := ssaBinOp.Y.(*ssa.Const); ok {
					if ssaConst.Value.Kind() == (constant.Int) {
						// Now we can tell the loop condition is like "if x < 3"

						// check if the counter == loop_condition, meaning we correctly unrolled the loop
						strLoopCondition := ssaConst.Value.ExactString() // this is string "3", we can't directly read int value from package constant
						intLoopCondition, err := strconv.Atoi(strLoopCondition)
						if err != nil {
							return false
						}
						if intLoopCondition != counter {
							// we incorrectly unrolled the loop, return true, which will delete this path
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// With the given paths and goroutines, check some Go rules. If rules passed, return a new pathCombination, else return nil.
// Will create extra goroutines if necessary. The paths of extra goroutines are chosen in random, and we try our best to make them different.
// (e.g., a path visits a *Go node 2 times, then we should have 2 goroutines created by this *Go)
func (g *SyncGraph) generateNewPathCombinations(goroutines []*Goroutine, paths []*LocalPath, possible_paths []map[int]*LocalPath) {
	// See if Go-rule is satisfied. Go-rule: A. If the goroutine is g.MainGoroutine, its path mustn't be nil
	//										 B. If a goroutine has no nil path, its creation must be on another path in pathCombination
	//										 C. If a *Go is in paths, it should create a not nil Goroutine
	//										 D. If there are N same *Go (imagine go in loop), they should create N Goroutines.
	//											Since there is always one goroutine existing, we will add extra (N-1) Goroutines
	//											This is kinda duplicate with C, but it's OK

	for i := 0; i < len(goroutines); i++ {
		goroutine := goroutines[i]
		path := paths[i]

		// check A
		if goroutine == g.MainGoroutine && path.IsEmpty() {
			return
		}

		// check B
		if path.IsEmpty() == false && goroutine.Creator != nil { // No nil path, and not head goroutine.
			// Need to verify this goroutine is created by another path
			flagFoundCreator := false
		loopOtherThread:
			for j, other := range paths {
				if i == j {
					continue
				}
				for _, node := range other.Path {
					if node == goroutine.Creator {
						flagFoundCreator = true
						break loopOtherThread
					}
				}
			}
			if flagFoundCreator == false {
				return
			}
		}

		// check C
		for _, node := range path.Path {
			nodeGo, ok := node.(*Go)
			if !ok {
				continue
			}

			flagFoundCreatedGoroutine := false
			for j, other := range goroutines {
				if other == goroutine {
					continue
				}
				other_path := paths[j]
				if other.Creator == nodeGo { // this goroutine is created by our Go
					flagFoundCreatedGoroutine = true
					if other_path.IsEmpty() {
						return
					}
				}
			}
			if flagFoundCreatedGoroutine == false {
				return
			}
		}
	}

	//// check D. This step may create extra goroutines. Exit the loop when the number of goroutines are stable
	//last_iteration_goroutine_num := -1
	//for last_iteration_goroutine_num != len(goroutines) {
	//	last_iteration_goroutine_num = len(goroutines)
	//
	//	for i := 0; i < len(goroutines); i++ {
	//		goroutine := goroutines[i]
	//		path := paths[i]
	//
	//		go_node_appear_time := make(map[*Go]int)
	//		for _, node := range path.Path {
	//			if go_node, ok := node.(*Go); ok {
	//				go_node_appear_time[go_node] += 1
	//			}
	//		}
	//
	//		// for each *Go shown more than 1 times, check if the number of created goroutines matches
	//		// if not, create extra goroutines, and extra paths. Choose paths randomly but try the best to make them different
	//		for go_node, times := range go_node_appear_time {
	//			if times <= 1 {
	//				continue
	//			}
	//			count_exist_createes := 0
	//			for _, goroutine := range goroutines {
	//				if goroutine.Creator == go_node {
	//					count_exist_createes++
	//				}
	//			}
	//			extra_needed := times - count_exist_createes
	//			// figure out in this pathCombination, which goroutine is created by go_node
	//			var created_thread *Goroutine
	//			for _, createe := range go_node.MapCreateGoroutines {
	//
	//			}
	//			for j := 0; j < extra_needed; j++ {
	//				// create
	//			}
	//		}
	//	}
	//}

	vecNewGoPaths := []*tupleGoroutinePath{}
	for i := 0; i < len(goroutines); i++ {
		newGoPath := &tupleGoroutinePath{
			goroutine: goroutines[i],
			path:      paths[i],
		}
		vecNewGoPaths = append(vecNewGoPaths, newGoPath)

		///DELETE
		GoMap := make(map[*Go]struct{})
		for _, node := range paths[i].Path {
			if nodeGo, ok := node.(*Go); ok {
				_, exist := GoMap[nodeGo]
				if exist {
					//fmt.Println("Found Go node show up twice")
					return
				} else {
					GoMap[nodeGo] = struct{}{}
				}
			}
		}
	}

	newTuple := &pathCombination{go_paths: vecNewGoPaths}
	if g.skipThisPathCombination(newTuple) {
		return
	}
	g.PathCombinations = append(g.PathCombinations, newTuple)
}

type TupleAPIFunc struct {
	PkgName  string
	FuncName string
}

var vecGiveupAPI []TupleAPIFunc = []TupleAPIFunc{
	{"os", "Exit"},
	{"syscall", "Exit"},
	{"reflect", "ValueOf"},
}

// this function uses some rules in "Strategies to reduce FP" document to skip some tuples
func (g *SyncGraph) skipThisPathCombination(t *pathCombination) bool {

	// Rule 2: Handle “for range” operation of channel
	// Implementation: find for range operation, if find, skip this pathCombination
	// See BuildGraph(), line 68. Remember to change

	// Rule 8: Handle special APIs
	for _, go_p := range t.go_paths {
		for _, node := range go_p.path.Path {
			inst := node.Instruction()
			if instCall, ok := inst.(*ssa.Call); ok {
				if callee, ok := instCall.Call.Value.(*ssa.Function); ok {
					if callee.Pkg != nil && callee.Pkg.Pkg != nil {
						pkgName := callee.Pkg.Pkg.Name()
						fnName := callee.Name()
						for _, giveupTuple := range vecGiveupAPI {
							if giveupTuple.PkgName == pkgName && giveupTuple.FuncName == fnName {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}
