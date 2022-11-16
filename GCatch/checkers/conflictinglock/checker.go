package conflictinglock

import (
	"fmt"
	"strings"

	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/util"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var AnalyzedFNs map[string]bool
var MapMutex map[string]*StMutex

var VecFNsWithLocking []*ssa.Function

var mapIIStLockingOp map[ssa.Instruction]*StLockingOp
var mapIIStUnlockingOp map[ssa.Instruction]*StUnlockingOp

var mapCallSiteCallee map[ssa.Instruction]map[*ssa.Function]bool

var mapIDCallChain map[int][]*callgraph.Edge

var numCurrentCallChainID int

const MaxCallChainDepth int = 6 // Be careful: Algorithm complexity: e^n, n = C5_call_chain_layer
const MaxInspectedFun int = 100000

var numInspectedFn int

var mapPairIDMutexPair map[string]*stMutexPair

var reportedPairs map[string]bool

func printCallChain(callchain []*callgraph.Edge) {

	mapFun := make(map[string]bool)
	boolRecursive := false
	containFnPointer := false
	for index, e := range callchain {

		if pCall, ok := e.Site.(*ssa.Call); ok {
			if _, ok := pCall.Call.Value.(*ssa.Function); ok {
			} else if _, ok := pCall.Call.Value.(ssa.Instruction); ok {
				containFnPointer = true
			}
		}

		if _, ok := mapFun[e.Caller.Func.String()]; !ok {
			mapFun[e.Caller.Func.String()] = true
		} else {
			boolRecursive = true
		}

		if index == len(callchain)-1 {
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

		fmt.Print(" -> ")

		if index == len(callchain)-1 {
			fmt.Print(e.Callee.Func.Name())
		}
	}

	fmt.Println()
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

	//mapIIStLockingOp = make(map[ssa.Instruction] * StLockingOp)
	//mapIIStUnlockingOp = make(map[ssa.Instruction] * StUnlockingOp)

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
	//mapCallSiteCallee = make(map[ssa.Instruction] map[* ssa.Function] bool)

	for _, node := range config.CallGraph.Nodes {
		for _, e := range node.Out {
			if ii, ok := e.Site.(ssa.Instruction); ok {
				if _, ok := mapCallSiteCallee[ii]; !ok {
					mapCallSiteCallee[ii] = make(map[*ssa.Function]bool)
				}
				mapCallSiteCallee[ii][e.Callee.Func] = true
			}
		}
	}
}

func collectLockingPair(vecLockingPair []*StLockPair, callchain []*callgraph.Edge) {

	if len(vecLockingPair) == 0 {
		return
	}

	tmp := make([]*callgraph.Edge, 0)

	for _, c := range callchain {
		tmp = append(tmp, c)
	}

	mapIDCallChain[numCurrentCallChainID] = tmp

	for _, p := range vecLockingPair {

		p.CallChainID = numCurrentCallChainID

		strMutexID1 := p.PLock1.I.Parent().Pkg.Pkg.Path() + ": " + p.PLock1.StrName + " (" + p.PLock1.Parent.StrBastStructType + ")"
		strMutexID2 := p.PLock2.I.Parent().Pkg.Pkg.Path() + ": " + p.PLock2.StrName + " (" + p.PLock2.Parent.StrBastStructType + ")"

		pairID := strMutexID1 + "\t\t" + strMutexID2

		if _, ok := mapPairIDMutexPair[pairID]; !ok {
			mapPairIDMutexPair[pairID] = &stMutexPair{
				PMutex1:        p.PLock1.Parent,
				PMutex2:        p.PLock2.Parent,
				VecLockingPair: make([]*StLockPair, 0),
			}
		}

		mapPairIDMutexPair[pairID].VecLockingPair = append(mapPairIDMutexPair[pairID].VecLockingPair, p)
	}

	numCurrentCallChainID++
}

func analyzeFN(fn *ssa.Function, callchain []*callgraph.Edge, context map[*StLockingOp]bool, depth int) {

	numInspectedFn++

	if numInspectedFn > MaxInspectedFun {
		return
	}

	if depth > MaxCallChainDepth {
		return
	}

	vecNameChain := []string{}

	for index, e := range callchain {
		vecNameChain = append(vecNameChain, e.Caller.Func.String())

		if index == len(callchain)-1 {
			vecNameChain = append(vecNameChain, e.Callee.Func.String())
		}
	}

	for _, fn1 := range vecNameChain {
		vecIndex := []int{}
		for i, fn2 := range vecNameChain {
			if fn2 == fn1 {
				vecIndex = append(vecIndex, i)
			}
		}

		if len(vecIndex) >= 3 {
			return
		}

		if len(vecIndex) == 2 {
			if vecIndex[1]+1 < len(vecNameChain) {
				s1 := vecNameChain[vecIndex[0]+1]
				s2 := vecNameChain[vecIndex[1]+1]
				if s1 == s2 {
					return
				}
			}
		}
	}

	newPair := GenKillAnalysis(fn, context)

	collectLockingPair(newPair, callchain)

	if node, ok := config.CallGraph.Nodes[fn]; ok {
		mapIIContextLock := make(map[*callgraph.Edge]map[*StLockingOp]bool)

		for _, e := range node.Out {
			if config.BoolDisableFnPointer {
				//if mapCallSiteCallee[]
				if ii, ok := e.Site.(ssa.Instruction); ok {
					if m, ok := mapCallSiteCallee[ii]; ok {
						if len(m) > 1 {
							continue
						}
					}
				}
			}

			if _, ok := e.Site.(*ssa.Defer); ok {
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
		}

		for e, contextLock := range mapIIContextLock {
			callchain = append(callchain, e)
			analyzeFN(e.Callee.Func, callchain, contextLock, depth+1)
			callchain = callchain[:len(callchain)-1]
		}
	}
}

func analyzeEntryFN(fn *ssa.Function) {

	//if fn.Name() != "Value" {
	//	return
	//}

	//fn.WriteTo(os.Stdout)

	numInspectedFn = 0
	depth := 0
	callchain := make([]*callgraph.Edge, 0)
	contextLock := make(map[*StLockingOp]bool)

	newPair := GenKillAnalysis(fn, contextLock)

	//fmt.Println(len(newPair))
	collectLockingPair(newPair, callchain)

	if node, ok := config.CallGraph.Nodes[fn]; ok {

		mapIIContextLock := make(map[*callgraph.Edge]map[*StLockingOp]bool)
		for _, e := range node.Out {
			if config.BoolDisableFnPointer {
				if ii, ok := e.Site.(ssa.Instruction); ok {
					if m, ok := mapCallSiteCallee[ii]; ok {
						if len(m) > 1 {
							continue
						}
					}
				}
			}

			if _, ok := e.Site.(*ssa.Defer); ok {
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
		}

		for e, contextLock := range mapIIContextLock {
			callchain = append(callchain, e)
			analyzeFN(e.Callee.Func, callchain, contextLock, depth)
			callchain = callchain[:len(callchain)-1]
		}
	}
}

func Initialize() {
	reportedPairs = make(map[string]bool)
}

func Detect() {

	MapMutex = make(map[string]*StMutex)
	VecFNsWithLocking = []*ssa.Function{}
	AnalyzedFNs = make(map[string]bool)

	mapIIStLockingOp = make(map[ssa.Instruction]*StLockingOp)
	mapIIStUnlockingOp = make(map[ssa.Instruction]*StUnlockingOp)

	mapCallSiteCallee = make(map[ssa.Instruction]map[*ssa.Function]bool)

	mapIDCallChain = make(map[int][]*callgraph.Edge)

	numCurrentCallChainID = 0

	mapPairIDMutexPair = make(map[string]*stMutexPair)
	mapReportedPair = make(map[string]struct{})

	getFunctionWithLockingOps()

	if config.BoolDisableFnPointer {
		getCallSiteCalleeMapping()

	}

	for _, fn := range VecFNsWithLocking {
		analyzeEntryFN(fn)
	}

	//fmt.Println("Whole map:")
	//fmt.Println(mapPairIDMutexPair)

	for k, _ := range mapPairIDMutexPair {
		mutexIDs := strings.Split(k, "\t\t")
		k2 := mutexIDs[1] + "\t\t" + mutexIDs[0]

		if _, ok := mapPairIDMutexPair[k2]; ok {
			_, boolReported1 := mapReportedPair[k]
			_, boolReported2 := mapReportedPair[k2]
			if boolReported1 == false && boolReported2 == false {
				fmt.Println("found Conflict Lock")
				fmt.Println(k)
				mapReportedPair[k] = struct{}{}
				mapReportedPair[k2] = struct{}{}
			}

		}
	}

}

var mapReportedPair map[string]struct{}
