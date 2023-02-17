package path

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strconv"
	"strings"
)

var LcaErrReachedMax = fmt.Errorf("Reached max layer number")
var LcaErrNilNode = fmt.Errorf("Node in callgraph for target function is nil")
var ErrNotComplete = fmt.Errorf("Incomplete callchains")
var ErrFindNone = fmt.Errorf("Found none")
var ErrNilNode = fmt.Errorf("Nil node in callgraph")
var ErrInaccurateCallgraph = fmt.Errorf("The callgraph is inaccurate: a caller has multiple callees or a callee has multiple callers")

type EdgeChain struct {
	Chain []*callgraph.Edge
	Start *callgraph.Node
}

func boolEdgeChainInSlice(e *EdgeChain, slice []*EdgeChain) bool {
	for _, old := range slice {
		if old.Equal(e) {
			return true
		}
	}
	return false
}

func (e *EdgeChain) Equal(q *EdgeChain) bool {
	if e.Start != q.Start {
		return false
	}
	if len(e.Chain) != len(q.Chain) {
		return false
	}
	for i, oneEdge := range e.Chain {
		if oneEdge != q.Chain[i] {
			return false
		}
	}
	return true
}

// If e.Chain is ABC, q.Chain is A or AB or ABC, return true
func (e *EdgeChain) Contains(q *EdgeChain) bool {
	if e.Start != q.Start {
		return false
	}
	if len(e.Chain) < len(q.Chain) {
		return false
	}
	for i := 0; i < len(q.Chain); i++ {
		if e.Chain[i] != q.Chain[i] {
			return false
		}
	}
	return true
}

// if chain is A-B-C, and lca is B, then return A-B
func cutChain(chain *EdgeChain, lca *ssa.Function) (newChain *EdgeChain) {
	newVecEdge := []*callgraph.Edge{}
	for _, edge := range chain.Chain {
		if edge.Callee.Func == lca {
			break
		}
		newVecEdge = append(newVecEdge, edge)
	}
	newChain = &EdgeChain{
		Chain: newVecEdge,
		Start: chain.Start,
	}
	return
}

// GetCallChains finds call chains from the entry function to each function in the input array.
//func GetCallChains(inputFuncs []*ssa.Function, entry *ssa.Function) (map[*ssa.Function][]*EdgeChain, error) {
//
//}

type LcaFinderWorkspace struct {
	mapChain2Ancestors map[string][]*callgraph.Node // a map from active chains to ancestor nodes of it. Chain ABCD: A,B,C,D
}

type LcaConfig struct {
	GiveUpWhenCallGraphIsInaccurate bool
	GiveUpWhenMaxLayerIsReached     bool
	SkipExternalFuncs               bool
}

// Since we are checking from the main function, the returned LCA is always main.
// However, we still need this FindLCA to generate callchains from main to all targetFn, which contains all channel operations
func FindLCA(vecTargetFn []*ssa.Function, lcaConfig LcaConfig, intMaxLayer int) (map[*ssa.Function][]*EdgeChain, error) {

	// A map from each node to the number of chains according to the last map. If only have chains ABC and BCDE, then B:2, E:1
	mapNode2NumChain := make(map[*callgraph.Node]int)

	// A map from chain's hash to chain
	mapEncode2Chain := make(map[string]*EdgeChain)

	// A map from each target to LcaFinderWorkspace. Used to detect circles
	mapTargetFn2workspace := make(map[*ssa.Function]*LcaFinderWorkspace)

	// initialize LcaFinderWorkspace for each target
	for _, fn := range vecTargetFn {
		newWorkspace := &LcaFinderWorkspace{
			mapChain2Ancestors: make(map[string][]*callgraph.Node),
		}
		mapTargetFn2workspace[fn] = newWorkspace

		node := config.CallGraph.Nodes[fn]
		if node == nil || fn == nil {
			return nil, LcaErrNilNode
		}

		initChain := &EdgeChain{
			Chain: nil,
			Start: node,
		}
		encode := encodeChain(initChain, mapEncode2Chain)
		newWorkspace.mapChain2Ancestors[encode] = []*callgraph.Node{node}
		mapNode2NumChain[node] = 1
	}

	countDepth := 0

	// The main loop. A breadth first search
	for {
		TargetToOldChains := make(map[*ssa.Function][]*EdgeChain)
		for target, workspace := range mapTargetFn2workspace {
			for encode, _ := range workspace.mapChain2Ancestors {
				chain := mapEncode2Chain[encode]
				TargetToOldChains[target] = append(TargetToOldChains[target], chain)
			}
		}

		// See if we can find the lowest common ancestor

		m, err, done := TryReportResult(mapTargetFn2workspace, mapNode2NumChain, mapEncode2Chain)
		if done {
			return m, err
		}

		// See if recursive number is reached
		if countDepth >= intMaxLayer && lcaConfig.GiveUpWhenMaxLayerIsReached == true {
			result, m, err, done := ReportExistingCallChains(mapTargetFn2workspace, mapNode2NumChain, mapEncode2Chain, countDepth)
			if done {
				return m, err
			}
			return result, LcaErrReachedMax
		}

		countDepth++

		// For each LcaFinderWorkspace
		for _, workspace := range mapTargetFn2workspace {

			// Update all active chains
			// We can't directly use a "for range LcaFinderWorkspace.mapChain2Ancestors", because we want a breadth-first search, but
			// the new values added to map in one search may be visited before other values in the same depth
			chainEncodes := []string{}
			ancestorGroups := [][]*callgraph.Node{}
			for encode, ancestors := range workspace.mapChain2Ancestors {
				chainEncodes = append(chainEncodes, encode)
				ancestorGroups = append(ancestorGroups, ancestors)
			}
			for i := 0; i < len(chainEncodes); i++ {
				encode := chainEncodes[i]
				ancestors := ancestorGroups[i]

				chain := mapEncode2Chain[encode]
				var lastNode *callgraph.Node // This is actually callee
				if len(chain.Chain) == 0 {
					lastNode = chain.Start
				} else {
					lastNode = chain.Chain[len(chain.Chain)-1].Caller
				}

				if len(lastNode.In) == 0 || !ContainsValidCaller(lastNode) { // skip if the lastNode has no caller
					continue
				}

				if len(lastNode.In) > 1 {
					if lcaConfig.GiveUpWhenCallGraphIsInaccurate {
						return nil, ErrInaccurateCallgraph
					}
					//fmt.Println(ErrInaccurateCallgraph)
				}

				// delete the current chain
				delete(workspace.mapChain2Ancestors, encode)
				for _, ancestor := range ancestors {
					mapNode2NumChain[ancestor] -= 1
				}

				// for each caller, create a new chain for it
				for _, in := range lastNode.In {
					if in.Caller.Func.Pkg == nil {
						//fmt.Printf("func.Pkg == nil, func = %s, func.Signature = %s\n", in.Caller.Func.Name(), in.Caller.Func.Signature)
					} else {
						path := in.Caller.Func.Pkg.Pkg.Path()
						//fmt.Println(path)
						if !config.IsPathIncluded(path) {
							//fmt.Printf("%s.%s: not included\n", path, in.Caller.Func.Name())
							continue
						}
					}
					boolInExist := false
					for _, existingEdge := range chain.Chain {
						if existingEdge == in {
							boolInExist = true
							break
						}
					}
					if boolInExist { // Avoid loop in callgraph
						continue
					}

					if in.Caller.Func.Synthetic != "" && in.Caller.Func.Name() == in.Callee.Func.Name() {
						// if A is interface call, we will have a synthetic A calls concrete A, and we don't want this
						continue
					}

					newChain := &EdgeChain{
						Chain: append(CopyEdgeSlice(chain.Chain), in),
						Start: chain.Start,
					}
					newEncode := encodeChain(newChain, mapEncode2Chain)
					newAncestors := append(CopyNodeSlice(ancestors), in.Caller)
					workspace.mapChain2Ancestors[newEncode] = newAncestors
					for _, ancestor := range newAncestors {
						mapNode2NumChain[ancestor] += 1
					}
				}
			}

			// check if the map from encode to ancestors is correct
			err2 := CheckEncodingToAncestorsCorrectness(workspace, mapEncode2Chain)
			if err2 != nil {
				return nil, err2
			}
		}
	}
}

func ContainsValidCaller(lastNode *callgraph.Node) bool {
	ret := false
	for _, in := range lastNode.In {
		ret = IsFunctionIncludedInAnalysis(in.Caller)
	}
	return ret
}

func IsFunctionIncludedInAnalysis(node *callgraph.Node) bool {
	if node.Func.Pkg == nil {
		if strings.HasPrefix(node.Func.Synthetic, "bound method wrapper for") {
			return true
		}
		//fmt.Printf("func.Pkg == nil, func = %s, func.Signature = %s\n", in.Caller.Func.Name(), in.Caller.Func.Signature)
	} else {
		path := node.Func.Pkg.Pkg.Path()
		//fmt.Println(path)
		if !config.IsPathIncluded(path) {
			//fmt.Printf("%s.%s: not included\n", path, in.Caller.Func.Name())
			//continue
		} else {
			return true
		}
	}
	return false
}

func CheckEncodingToAncestorsCorrectness(workspace *LcaFinderWorkspace, mapEncode2Chain map[string]*EdgeChain) error {
	correct := true
	printWarning := func() {
		correct = false
		//fmt.Println("Warning in FindLCA: mapChain2Ancestors is not correct")
	}
	for encode, ancestors := range workspace.mapChain2Ancestors {
		path := mapEncode2Chain[encode]
		if len(path.Chain) == 0 {
			if len(ancestors) != 1 || ancestors[0] != path.Start {
				printWarning()
			}
		} else {
			if len(path.Chain)+1 != len(ancestors) {
				printWarning()
			}
			for i, edge := range path.Chain {
				if edge.Callee != ancestors[i] {
					printWarning()
				}
				if i == len(path.Chain)-1 {
					if edge.Caller != ancestors[i+1] {
						printWarning()
					}
				}
			}
		}
	}
	if correct == false {
		return fmt.Errorf("mapChain2Ancestors is not correct")
	}
	return nil
}

func TryReportResult(mapTargetFn2workspace map[*ssa.Function]*LcaFinderWorkspace, mapNode2NumChain map[*callgraph.Node]int, mapEncode2Chain map[string]*EdgeChain) (map[*ssa.Function][]*EdgeChain, error, bool) {
	intActiveChains := 0
	for _, workspace := range mapTargetFn2workspace {
		intActiveChains += len(workspace.mapChain2Ancestors)
	}

	for node, intNumChains := range mapNode2NumChain {
		// SZ: It seems that when the number of chains matches, that means all existing discovered chains start
		// with the entry function.
		if intNumChains == intActiveChains && (node.Func.Name() == "main" || strings.HasPrefix(node.Func.Name(), "Test")) {
			// Now we find one LCA main that can cover every target
			oneLca := node.Func
			result := make(map[*ssa.Function][]*EdgeChain)
			vecEdgePaths := []*EdgeChain{}

			for _, workspace := range mapTargetFn2workspace {
				for encode, _ := range workspace.mapChain2Ancestors {
					chain := mapEncode2Chain[encode]
					chain = cutChain(chain, oneLca)
					reversedChain := reverseChain(chain)
					if boolEdgeChainInSlice(reversedChain, vecEdgePaths) == false {
						vecEdgePaths = append(vecEdgePaths, reversedChain)
					}
				}
			}

			result[oneLca] = vecEdgePaths
			return result, nil, true
		}
	}
	return nil, nil, false
}

func ReportExistingCallChains(mapTargetFn2workspace map[*ssa.Function]*LcaFinderWorkspace, mapNode2NumChain map[*callgraph.Node]int, mapEncode2Chain map[string]*EdgeChain, countDepth int) (map[*ssa.Function][]*EdgeChain, map[*ssa.Function][]*EdgeChain, error, bool) {
	// recursive number is reached. Now we return a map with multiple LCAs, each LCA can cover different vecTargetFn
	// Note: if this is used in our channel checking, you must print "!!!!" with a warning, like GCatch/syncgraph/task.go
	result := make(map[*ssa.Function][]*EdgeChain)

	// A big map of all mapChain2Ancestors
	mapOfMapChain2ancestors := make(map[string][]*callgraph.Node)
	chain2target := make(map[string]*ssa.Function)
	for target, workspace := range mapTargetFn2workspace {
		for encode, chain := range workspace.mapChain2Ancestors {
			mapOfMapChain2ancestors[encode] = chain
			chain2target[encode] = target
		}
	}

	for len(mapOfMapChain2ancestors) > 0 {

		// find a node that can cover the max number of chains
		// first, list all nodes that can cover the max number of chains
		champions := []*callgraph.Node{}
		var intGreatestNumOfChains int
		for node, intNumOfChains := range mapNode2NumChain {
			if intNumOfChains > intGreatestNumOfChains {
				champions = []*callgraph.Node{node}
				intGreatestNumOfChains = intNumOfChains
			} else if intNumOfChains == intGreatestNumOfChains {
				champions = append(champions, node)
				intGreatestNumOfChains = intNumOfChains
			}
		}

		var currentChampion *callgraph.Node
		for _, this := range champions {
			callees := []*callgraph.Node{}
			for _, out := range this.Out {
				callees = append(callees, out.Callee) // callees may have duplicated nodes, and even this itself
			}

			// if this has a callee in champions, remove this
			boolHasCallee := false
		otherLoop:
			for _, other := range champions {
				if other == this {
					continue
				}
				for _, callee := range callees {
					if callee == other {
						boolHasCallee = true
						break otherLoop
					}
				}
			}
			if boolHasCallee == false {
				currentChampion = this
				break
			}
		}
		if currentChampion == nil {
			currentChampion = champions[0] // if nil, choose the first one
		}

		// list all chains that has currentChampion as an ancestor
		vecCoveredEdgeChains := []*EdgeChain{}
		for encode, ancestors := range mapOfMapChain2ancestors {
			boolHasChampion := false
			for _, ancestor := range ancestors {
				if ancestor == currentChampion {
					boolHasChampion = true
					break
				}
			}
			if boolHasChampion {
				// num--
				for _, ancestor := range ancestors {
					if mapNode2NumChain[ancestor] == 0 {
						fmt.Print()
					}
					mapNode2NumChain[ancestor] -= 1
				}
				// add to slice
				chain := mapEncode2Chain[encode]
				reversedChain := reverseChain(chain) // chain is from callee to caller. reverse it

				var coveredEdgeChain *EdgeChain
				// case 1: reversedChain.Chain is nil, and Start is champion
				if len(reversedChain.Chain) == 0 && reversedChain.Start == currentChampion {
					coveredEdgeChain = &EdgeChain{
						Chain: nil,
						Start: currentChampion,
					}
				} else {
					// case 2: reversedChain.Chain has champion as caller or callee
					indexChampionAsCaller := -1
					indexChampionAsCallee := -1
					for i, edge := range reversedChain.Chain {
						if edge.Caller == currentChampion {
							indexChampionAsCaller = i
						}
						if edge.Callee == currentChampion {
							indexChampionAsCallee = i
						}
					}

					if indexChampionAsCaller > -1 { // case 2.1: as caller
						coveredPath := []*callgraph.Edge{}
						for i, edge := range reversedChain.Chain {
							if i < indexChampionAsCaller {
								continue
							} else {
								coveredPath = append(coveredPath, edge)
							}
						}
						coveredEdgeChain = &EdgeChain{
							Chain: coveredPath,
							Start: currentChampion,
						}
					} else if indexChampionAsCallee > -1 { // case 2.2: as callee, meaning champion is the last function
						coveredEdgeChain = &EdgeChain{
							Chain: nil,
							Start: currentChampion,
						}
					} else {
						fmt.Println("Fatal in FindLCA: champion is not in reversedChain")
						panic(1)
					}
				}
				if boolEdgeChainInSlice(coveredEdgeChain, vecCoveredEdgeChains) == false {
					vecCoveredEdgeChains = append(vecCoveredEdgeChains, coveredEdgeChain)
				}

				// delete this chain from big map
				delete(mapOfMapChain2ancestors, encode)
			}
		}

		if mapNode2NumChain[currentChampion] != 0 {
			err := fmt.Errorf("Warning in FindLCA: a node's number of chains is not 0 after we find "+
				"all chains containing it:%d. countDepth is :%d", mapNode2NumChain[currentChampion], countDepth)
			return nil, nil, err, true
		}

		result[currentChampion.Func] = vecCoveredEdgeChains
	}

	fmt.Println("Count Depth:", countDepth)
	return result, nil, nil, false
}

func encodeChain(chain *EdgeChain, encode2chain map[string]*EdgeChain) string {
	str := ""
	str += chain.Start.Func.String() + " "
	for _, edge := range chain.Chain {
		boolNilSite := edge.Site == nil
		if boolNilSite {
			str += "NilSite" + fmt.Sprintf("%p", edge) + edge.Caller.String() + " "
		} else {
			str += strconv.Itoa(int(edge.Site.Pos())) + edge.Caller.String() + " "
		}
	}
	encode2chain[str] = chain
	return str
}

func reverseChain(old *EdgeChain) *EdgeChain {
	reversedEdgeSlice := []*callgraph.Edge{}
	for i := len(old.Chain) - 1; i >= 0; i-- {
		reversedEdgeSlice = append(reversedEdgeSlice, old.Chain[i])
	}
	var start *callgraph.Node
	if len(old.Chain) == 0 {
		start = old.Start
	} else {
		start = reversedEdgeSlice[0].Caller
	}
	return &EdgeChain{
		Chain: reversedEdgeSlice,
		Start: start,
	}
}

func CopyEdgeSlice(old []*callgraph.Edge) []*callgraph.Edge {
	result := []*callgraph.Edge{}
	for _, o := range old {
		result = append(result, o)
	}
	return result
}

func CopyNodeSlice(old []*callgraph.Node) []*callgraph.Node {
	result := []*callgraph.Node{}
	for _, o := range old {
		result = append(result, o)
	}
	return result
}

// Every Select is only used in Extract. Among all Extract, only one Extract's Index is 0 (0 means obtaining the case index)
// This Extract will be used by N BinOp(where OP is ==), N is the number of cases (default is not counted)
// If the case is not empty, the next inst will be If; if the case is empty, the next inst can be anything
// If Select has case 0,1,2,(default), then BinOp.Y will be 0,1,2. Find the BinOp.Y == 2. If the next inst is If,
// default is If.Else. If the next is not If, meaning default and the last case are going to the same place
func FindSelectNexts(s *ssa.Select) (map[int]ssa.Instruction, error) {
	var e *ssa.Extract
	for _, r := range *(s.Referrers()) {
		if rAsExtract, ok := r.(*ssa.Extract); ok {
			if rAsExtract.Index == 0 {
				e = rAsExtract
				break
			}
		}
	}

	if e == nil { // This should never happen
		output.PrintIISrc(s)
		return nil, fmt.Errorf("Warning in find_select_nexts: no Extract found")
	}

	mapIndex2binop := make(map[int]*ssa.BinOp)
	for _, r := range *(e.Referrers()) {
		binop, ok := r.(*ssa.BinOp)
		if !ok { // This should never happen
			output.PrintIISrc(s)
			output.PrintIISrc(binop)
			return nil, fmt.Errorf("Warning in find_select_nexts: one of Extract.Referrer is not BinOp")
		}
		defer func() {
			if config.RecoverFromError {
				if r := recover(); r != nil {
					fmt.Println("Warning in find_select_nexts: BinOp.Y can't be converted into int")
					output.PrintIISrc(s)
					output.PrintIISrc(binop)
				}
			}
		}()
		intCaseIndex := int(binop.Y.(*ssa.Const).Int64()) // This should never panic
		mapIndex2binop[intCaseIndex] = binop
	}

	result := make(map[int]ssa.Instruction)
	for index, binop := range mapIndex2binop {
		referrers := *binop.Referrers()
		if len(referrers) == 1 { // 1 referrer. This must be If. If.Then is the case. When index is the
			// last index of all cases, and default exists, then If.Else is the default
			inst_if, ok := referrers[0].(*ssa.If)
			if !ok { // This should never happen
				output.PrintIISrc(s)
				output.PrintIISrc(referrers[0])
				return nil, fmt.Errorf("Warning in find_select_nexts: binop's one and only referrer is not If")
			}
			thenBB, elseBB := inst_if.Block().Succs[0], inst_if.Block().Succs[1]
			thenInst, elseInst := thenBB.Instrs[0], elseBB.Instrs[0]
			result[index] = thenInst

			if s.Blocking == false { // Has default
				if index == len(mapIndex2binop)-1 { // This is the last case
					result[-1] = elseInst
				}
			}

		} else if len(referrers) == 0 { //No referrer. Meaning there is no If. This is an empty case
			result[index] = nextInst(binop)
			if s.Blocking == false { // Has default
				if index == len(mapIndex2binop)-1 { // This is the last case
					result[-1] = nextInst(binop)
				}
			}
		} else { // This should never happen
			output.PrintIISrc(s)
			for _, r := range referrers {
				output.PrintIISrc(r)
			}
			return nil, fmt.Errorf("Warning in find_select_nexts: binop has multiple referrers")
		}
	}

	return result, nil
}

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

// Enumerate all paths between bb1 and bb2. Requires that bb2 post-dominates bb1
func EnumeratePathForPostDomBBs(bb1, bb2 *ssa.BasicBlock) [][]*ssa.BasicBlock {
	result := [][]*ssa.BasicBlock{}
	pathWorklist := [][]*ssa.BasicBlock{[]*ssa.BasicBlock{bb1}}

	for len(pathWorklist) > 0 {
		path := pathWorklist[len(pathWorklist)-1]
		pathWorklist = pathWorklist[:len(pathWorklist)-1]

		lastBB := path[len(path)-1]
		if lastBB == bb2 {
			result = append(result, path)
			continue
		}

		for _, suc := range lastBB.Succs {
			// if suc has appeared in path, skip it
			boolSkip := false
			for _, existBB := range path {
				if suc == existBB {
					boolSkip = true
					break
				}
			}
			if boolSkip {
				continue
			}
			newPath := append(path, suc)
			pathWorklist = append(pathWorklist, newPath)
		}
	}

	return result
}

func PathBetweenInst(inst1, inst2 ssa.Instruction) (result []*ssa.BasicBlock) {
	result = []*ssa.BasicBlock{}
	if inst1.Parent() == inst2.Parent() {
		result = PathBetweenLocalInst(inst1, inst2)
	}
	return
}

func PathBetweenLocalInst(inst1, inst2 ssa.Instruction) (result []*ssa.BasicBlock) {
	result = []*ssa.BasicBlock{}
	bb1 := inst1.Block()
	bb2 := inst2.Block()
	if bb1 == bb2 {
		bb := bb1
		var index1, index2 int
		for i, inst := range bb.Instrs {
			if inst == inst1 {
				index1 = i
			} else if inst == inst2 {
				index2 = i
			}
		}
		if index1 <= index2 {
			result = append(result, bb)
		}
		return
	}
	result = DFSPathLocalBB(bb1, bb2, result)

	return
}

func DFSPathLocalBB(bbCur, bbTarget *ssa.BasicBlock, path []*ssa.BasicBlock) (result []*ssa.BasicBlock) {
	if bbCur == bbTarget {
		result = path
		return
	}
	for _, bbSuc := range bbCur.Succs {
		isSucInPath := false
		for _, bb := range path {
			if bb == bbSuc {
				isSucInPath = true
				break
			}
		}
		if isSucInPath == false {
			newPath := append(path, bbSuc)
			result = DFSPathLocalBB(bbSuc, bbTarget, newPath)
		}
	}
	return
}
