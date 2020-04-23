package doublelock

import (
	"fmt"
	"github.com/system-pclub/gochecker/output"
	"github.com/system-pclub/gochecker/tools/go/callgraph"
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"github.com/system-pclub/gochecker/tools/go/ssa/ssautil"
)

var AnalyzedFNs map[string] bool
var MapMutex map[string] * StMutex

var vecReportedBugs [] * StDoubleLock
var VecFNsWithLocking   []* ssa.Function

var mapIIStLockingOp map[ssa.Instruction] * StLockingOp
var mapIIStUnlockingOp map[ssa.Instruction] * StUnlockingOp

const MaxCallChainDepth int = 8 // Be careful: Algorithm complexity: e^n, n = C5_call_chain_layer
const MaxInspectedFun int = 100000

var numInspectedFn int



func isReported(bug * StDoubleLock) bool {
	for _, b := range vecReportedBugs {
		if b.PLock1 == bug.PLock1 && b.PLock2 == b.PLock2 {
			return true
		}
	}

	return false
}

func printCallChain(callchain [] * callgraph.Edge) {
	fmt.Println("Call Chain: ")
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
	printCallChain(callchain)
	fmt.Println("\tLocation of the 2 lock operations:")
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

		hasLockingOp := handleFN(fn)

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

func analyzeFN(fn * ssa.Function, callchain []* callgraph.Edge, contextLock map[* StLockingOp] bool, depth int) {
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



	newbugs := GenKillAnalysis(fn, contextLock)

	for _, bug := range newbugs {
		reportDoubleLock(bug, callchain)
	}

	if node, ok := config.CallGraph.Nodes[fn]; ok {

		mapIIContextLock := make(map[* callgraph.Edge] map[* StLockingOp] bool)

		for _, e := range node.Out {
			contextLock := GetLiveMutex(e.Site)
			if len(contextLock) == 0 {
				continue
			}
			if _, ok := e.Site.(*ssa.Go); ok {
				continue
			}
			mapIIContextLock[e] = contextLock
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

			contextLock := GetLiveMutex(e.Site)

			if len(contextLock) == 0 {
				continue
			}

			if _, ok := e.Site.(*ssa.Go); ok {
				continue
			}

			mapIIContextLock[e] = contextLock
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


func Detect() {
	vecReportedBugs = [] * StDoubleLock{}
	MapMutex = make(map[string] * StMutex)
	VecFNsWithLocking =  []* ssa.Function{}
	AnalyzedFNs = make(map[string] bool)

	getFunctionWithLockingOps()


	for _, fn := range VecFNsWithLocking {
		analyzeEntryFN(fn)
	}
}