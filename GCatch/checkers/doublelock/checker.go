package doublelock

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa/ssautil"
	"github.com/system-pclub/GCatch/GCatch/util"
)

var AnalyzedFNs map[string] bool
var MapMutex map[string] * StMutex

var vecReportedBugs [] * StDoubleLock
var VecFNsWithLocking   []* ssa.Function

var mapIIStLockingOp map[ssa.Instruction] * StLockingOp
var mapIIStUnlockingOp map[ssa.Instruction] * StUnlockingOp

var mapCallSiteCallee map[ssa.Instruction] map[* ssa.Function] bool



const MaxCallChainDepth int = 6 // Be careful: Algorithm complexity: e^n, n = C5_call_chain_layer
const MaxInspectedFun int = 100000

var numInspectedFn int



func isReported(bug * StDoubleLock) bool {
	for _, b := range vecReportedBugs {
		if b.PLock1 == bug.PLock1 && b.PLock2 == b.PLock2 {
			return true
		}

		/*
		l1 := config.Prog.Fset.Position(bug.PLock1.I.Pos())
		l2 := config.Prog.Fset.Position(b.PLock1.I.Pos())

		if l2.Line > 0 && l1.Filename == l2.Filename && l1.Line == l2.Line {
			l1 := config.Prog.Fset.Position(bug.PLock2.I.Pos())
			l2 := config.Prog.Fset.Position(b.PLock2.I.Pos())

			if l2.Line > 0 && l1.Filename == l2.Filename && l1.Line == l2.Line {
				return true
			}
		}
		 */
		if b.PLock1.NumLine > 0 && bug.PLock1.StrFileName == b.PLock1.StrFileName && bug.PLock1.NumLine == b.PLock1.NumLine {
			if b.PLock2.NumLine > 0 && bug.PLock2.StrFileName == b.PLock2.StrFileName && bug.PLock2.NumLine == b.PLock2.NumLine {
				return true
			}
		}
	}

	return false
}

func printCallChain(callchain [] * callgraph.Edge) {

	mapFun := make(map[string] bool)
	boolRecursive := false
	containFnPointer := false
	for index, e := range callchain {

		if pCall, ok := e.Site.(* ssa.Call); ok {
			if _, ok := pCall.Call.Value.(* ssa.Function); ok {
			} else if _, ok := pCall.Call.Value.(ssa.Instruction); ok {
				containFnPointer = true
			}
		}


		if _, ok := mapFun[e.Caller.Func.String()]; !ok {
			mapFun[e.Caller.Func.String()] = true
		} else {
			boolRecursive = true
		}


		if index == len(callchain) - 1 {
			if _, ok := mapFun[e.Callee.Func.String()]; !ok {
				mapFun[e.Callee.Func.String()] = true
			} else {
				boolRecursive = true
			}
		}
	}

	if boolRecursive {
		fmt.Print("[Recursive] Call Chain")
	} else {
		fmt.Print("Call Chain")
	}

	if containFnPointer {
		fmt.Println(" (with FN Pointer):")
	} else {
		fmt.Println(":")
	}

	for index, e := range callchain {
		fmt.Print(e.Caller.Func.Name())
		loc := (config.Prog.Fset).Position(e.Site.Pos())
		if loc.Line > 0 {
			fmt.Print(" (at ", loc.Filename, ": ", loc.Line, ")")
		}

		//if index < len(callchain) - 1 {
		fmt.Print(" -> ")
		//}

		if index == len(callchain) - 1 {
			fmt.Print(e.Callee.Func.Name())
		}
	}
	fmt.Println()
}

func reportDoubleLock(newbug * StDoubleLock, callchain []* callgraph.Edge) {
	//newbug := StDoubleLock{
	//	PLock1: lock1,
	//	PLock2: lock2,
	//}

	if isReported(newbug) {
		return
	}

	vecReportedBugs = append(vecReportedBugs, newbug)

	config.BugIndexMu.Lock()
	config.BugIndex ++
	fmt.Print("----------Bug[")
	fmt.Print(config.BugIndex)
	fmt.Println("]----------\n\tType: Double Lock \tReason: A Mutex/RWMutex is locked twice. (Note: even double RWMutex.RLock() can produce deadlock bug)\n")
	if len(callchain) == 0 {
		if newbug.PLock1.I == newbug.PLock2.I {
			fmt.Print("Same lock acquired in a loop")
		}
	} else {
		printCallChain(callchain)
	}

	fmt.Println("\tLocation of the 2 lock operations:")

	//newbug.PLock1.I.Parent().WriteTo(os.Stdout)
	//newbug.PLock2.I.Parent().WriteTo(os.Stdout)


	output.PrintIISrc(newbug.PLock1.I)
	output.PrintIISrc(newbug.PLock2.I)
	config.BugIndexMu.Unlock()
}

func getFunctionWithLockingOps() {

	for fn, _ := range ssautil.AllFunctions(config.Prog) {
		if fn == nil {
			continue
		}

		if config.IsPathIncluded(fn.String()) == false {
			continue
		}

		if _, ok := AnalyzedFNs[fn.String()]; ok {
			continue
		}

		AnalyzedFNs[fn.String()] = true

		//fmt.Println(fn.String())

		hasLockingOp := handleFN(fn)

		//if fn.String() == "(*google.golang.org/grpc/xds/internal/balancer/balancergroup.BalancerGroup).newSubConn" {
		//	fmt.Println("found", hasLockingOp)
		//}

		if hasLockingOp {
			VecFNsWithLocking = append(VecFNsWithLocking, fn)
		}
	}

	mapIIStLockingOp = make(map[ssa.Instruction] * StLockingOp)
	mapIIStUnlockingOp = make(map[ssa.Instruction] * StUnlockingOp)


	for _, stMutex := range MapMutex {
		for ii, l := range stMutex.MapLockingOps {
			mapIIStLockingOp[ii] = l
		}

		for ii, ul := range stMutex.MapUnlockingOps {
			mapIIStUnlockingOp[ii] = ul
		}
	}
}

func getCallSiteCalleeMapping() {
	mapCallSiteCallee = make(map[ssa.Instruction] map[* ssa.Function] bool)

	for _, node := range config.CallGraph.Nodes {
		for _, e := range node.Out {
			if ii, ok := e.Site.(ssa.Instruction); ok {
				if _, ok := mapCallSiteCallee[ii]; !ok {
					mapCallSiteCallee[ii] = make(map[* ssa.Function] bool)
				}
				mapCallSiteCallee[ii][e.Callee.Func] = true
			}
		}
	}

	/*
	numSites := 0
	numCount := 0
	maxSite := 0


	for _, m := range mapCallSiteCallee {
		if len(m) > 1 {
			numCount ++
		}

		if len(m) > maxSite {
			maxSite = len(m)
		}

		numSites += len(m)
	}

	fmt.Println("# call sites: ", numSites, "# of ambigous sites: ", numCount, "max # of callees ", maxSite)
	 */
}

func analyzeFN(fn * ssa.Function, callchain []* callgraph.Edge, context map[* StLockingOp] bool, depth int) {


	numInspectedFn ++

	if numInspectedFn > MaxInspectedFun {
		return
	}

	if depth > MaxCallChainDepth {
		return
	}

	vecNameChain := []  string {}

	for index, e := range callchain {
		vecNameChain = append(vecNameChain, e.Caller.Func.String())

		if index == len(callchain) -1 {
			vecNameChain = append(vecNameChain, e.Callee.Func.String())
		}
	}

	for _, fn1 := range vecNameChain {
		vecIndex := [] int{}
		for i, fn2 := range vecNameChain {
			if fn2 == fn1 {
				vecIndex = append(vecIndex, i)
			}
		}

		if len(vecIndex) >= 3 {
			return
		}

		if len(vecIndex) == 2 {
			if vecIndex[1] + 1 < len(vecNameChain) {
				s1 := vecNameChain[vecIndex[0] + 1]
				s2 := vecNameChain[vecIndex[1] + 1]
				if s1 == s2 {
					return
				}
			}
		}
	}

	newbugs := GenKillAnalysis(fn, context)

	for _, bug := range newbugs {
		reportDoubleLock(bug, callchain)
	}


	if node, ok := config.CallGraph.Nodes[fn]; ok {
		mapIIContextLock := make(map[* callgraph.Edge] map[* StLockingOp] bool)

		for _, e := range node.Out {

			if config.BoolDisableFnPointer {
				//if mapCallSiteCallee[]
				if ii, ok := e.Site.(ssa.Instruction); ok {
					if  m, ok := mapCallSiteCallee[ii]; ok {
						if len(m) > 1 {
							continue
						}
					}
				}
			}

			if _, ok := e.Site.(* ssa.Defer); ok {
				IIs := util.GetExitInsts(fn)
				for _, ii := range IIs {
					contextLock := GetLiveMutex(ii)
					if len(contextLock) == 0 {
						continue
					}
					mapIIContextLock[e] = contextLock
				}
			} else {

				contextLock := GetLiveMutex(e.Site)

				if len(contextLock) == 0 {
					continue
				}

				if _, ok := e.Site.(*ssa.Go); ok {
					continue
				}

				if _, ok = mapIIStUnlockingOp[e.Site]; ok {
					continue
				}

				mapIIContextLock[e] = contextLock
			}



			/*


			contextLock := GetLiveMutex(e.Site)
			if len(contextLock) == 0 {
				continue
			}

			//fmt.Println("contextLock: ", len(contextLock))
			if _, ok := e.Site.(*ssa.Go); ok {
				continue
			}
			mapIIContextLock[e] = contextLock

			 */
			//callchain = append(callchain, e)
			//analyzeFN(e.Callee.Func, callchain, contextLock, depth)
			//callchain = callchain[:len(callchain)-1]
		}

		for e, contextLock := range mapIIContextLock {
			callchain = append(callchain, e)
			analyzeFN(e.Callee.Func, callchain, contextLock, depth + 1)
			callchain = callchain[: len(callchain) - 1]
		}
	}
}

func analyzeEntryFN(fn * ssa.Function) {

	numInspectedFn = 0
	depth := 0
	callchain := make([] * callgraph.Edge, 0)
	contextLock := make(map[* StLockingOp] bool)

	newbugs := GenKillAnalysis(fn, contextLock)

	for _, bug := range newbugs {
		reportDoubleLock(bug, callchain)
	}

	if node, ok := config.CallGraph.Nodes[fn]; ok {

		mapIIContextLock := make(map[* callgraph.Edge] map[* StLockingOp] bool)

		for _, e := range node.Out {

			//fmt.Println(e.Site, e.Callee.Func.Name())

			if config.BoolDisableFnPointer {
				//if mapCallSiteCallee[]
				if ii, ok := e.Site.(ssa.Instruction); ok {
					if  m, ok := mapCallSiteCallee[ii]; ok {
						if len(m) > 1 {
							continue
						}
					}
				}
			}

			if _, ok := e.Site.(* ssa.Defer); ok {
				if ii, ok := e.Site.(ssa.Instruction); ok {
					if util.IsFirstDefer(ii) {
						continue
					}
				}

				IIs := util.GetExitInsts(fn)
				for _, ii := range IIs {
					contextLock := GetLiveMutex(ii)
					if len(contextLock) == 0 {
						continue
					}
					mapIIContextLock[e] = contextLock
				}
			} else {
				contextLock := GetLiveMutex(e.Site)

				if len(contextLock) == 0 {
					continue
				}

				if _, ok := e.Site.(*ssa.Go); ok {
					continue
				}

				if _, ok = mapIIStUnlockingOp[e.Site]; ok {
					continue
				}

				mapIIContextLock[e] = contextLock
			}

			//callchain = append(callchain, e)
			//analyzeFN(e.Callee.Func, callchain, contextLock, depth)
			//callchain = callchain[:len(callchain)-1]
		}

		for e, contextLock := range mapIIContextLock {
			callchain = append(callchain, e)
			analyzeFN(e.Callee.Func, callchain, contextLock, depth)
			callchain = callchain[: len(callchain) -1]
		}
	}
}

func Initialize() {
	vecReportedBugs = [] * StDoubleLock{}
}


func Detect() {

	MapMutex = make(map[string] * StMutex)
	VecFNsWithLocking =  []* ssa.Function{}
	AnalyzedFNs = make(map[string] bool)

	getFunctionWithLockingOps()


	if config.BoolDisableFnPointer {
		getCallSiteCalleeMapping()

	}

	//for strKey, _ := range MapMutex {
	//	fmt.Println(strKey)
	//}

	/*
	vecFun := make([] string, 0)

	for _, fn := range VecFNsWithLocking {
		//analyzeEntryFN(fn)
		vecFun = append(vecFun, fn.String())
	}

	fmt.Println(len(vecFun))

	sort.Strings(vecFun)

	for _, s := range vecFun {
		fmt.Println(s)
	}
	*/


	for _, fn := range VecFNsWithLocking {
		analyzeEntryFN(fn)
	}

}
