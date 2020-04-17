package forgetunlock

import (
	"fmt"
	"github.com/system-pclub/gochecker/analysis"
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/instinfo"
	"github.com/system-pclub/gochecker/output"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"github.com/system-pclub/gochecker/tools/go/ssa/ssautil"
	"github.com/system-pclub/gochecker/util"
	"go/token"
	"strings"
)


type stCond struct {
	Cond ssa.Value
	Flag bool
}



var FNOnlyContainUnlock map[string] string
var Bugs [] ssa.Instruction
var AnalyzedFNs map[string] bool

var numFunction int

func strInsensitiveEqual(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return s == substr
}

func strInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func Initialize() {
	FNOnlyContainUnlock = make(map[string] string)
	AnalyzedFNs = make(map[string] bool)
	Bugs = [] ssa.Instruction{}
}



func collectDominator(inputBB * ssa.BasicBlock) map[* ssa.BasicBlock] bool {

	mapDominators := make(map[* ssa.BasicBlock] bool)

	for _, bb := range inputBB.Parent().Blocks {
		if bb.Dominates(inputBB) {
			mapDominators[bb] = true
		}
	}

	return mapDominators
}

func compDominatorIntersection(m1 map[* ssa.BasicBlock] bool, m2 map[* ssa.BasicBlock] bool)  map[ * ssa.BasicBlock] bool {
	mapIntersection := make(map[*ssa.BasicBlock] bool)

	for bb, _ := range m1 {
		if _, ok := m2[bb]; ok {
			mapIntersection[bb] = true
		}
	}

	return mapIntersection
}


func getLowestDominator(m map[* ssa.BasicBlock] bool) * ssa.BasicBlock {

	for bb1, _ := range m {
		flag := true
		for bb2, _ := range m {
			if bb1 == bb2 {
				continue
			}

			if bb1.Dominates(bb2) {
				flag = false
				break
			}
		}

		if flag {
			return bb1
		}
	}

	return nil
}

func getShortestPath(b1 *ssa.BasicBlock, b2 *ssa.BasicBlock) [] * ssa.BasicBlock {

	mapWorkList := make(map[* ssa.BasicBlock] [] * ssa.BasicBlock)
	bbVisited := make(map[* ssa.BasicBlock] bool)

	mapWorkList[b1] = []*ssa.BasicBlock {b1}

	if b1 == b2 {
		return mapWorkList[b1]
	}



	bbVisited[b1] = true

	for len(mapWorkList) > 0 {
		mapWorkList2 := make(map[* ssa.BasicBlock] [] * ssa.BasicBlock )

		for bb, path := range mapWorkList {
			for _, succ := range bb.Succs {
				if succ == b2 {
					path = append(path, b2)
					return path
				}

				if _, ok := bbVisited[succ]; ok {
					continue
				}

				newPath := append(path, succ)
				mapWorkList2[succ] = newPath
			}
		}

		mapWorkList = mapWorkList2
	}

	return [] * ssa.BasicBlock {}

}

func printPath(p []* ssa.BasicBlock) {

	for _, bb := range p[:len(p)-1] {
		fmt.Print(bb.Index, "->")
	}

	fmt.Println(p[len(p)-1].Index)

}


func generatePathConditions(vecPath [] * ssa.BasicBlock, pd * analysis.PostDominator) [] stCond {

	vecKeep := [] bool {}
	index := 0

	for index < len(vecPath) {
		vecKeep = append(vecKeep, true)
		index ++
	}

	ii := len(vecPath) - 1

	for ii >= 0 {
		jj := ii - 1

		for jj >= 0 {
			if pd.Dominate(vecPath[ii], vecPath[jj]) {
				vecKeep[jj] = false
			} else {
				break
			}

			jj --
		}

		ii = jj

	}

	vecPathConds := [] stCond {}

	index = 0

	for index < len(vecPath) - 1 {
		if vecKeep[index] {
			iiLast := vecPath[index].Instrs[len(vecPath[index].Instrs) -1]
			iiIF, ok := iiLast.(*ssa.If)
			if ok {
				var newCond stCond
				newCond.Cond = iiIF.Cond

				if vecPath[index+1] == vecPath[index].Succs[0] {
					newCond.Flag = true
				} else {
					newCond.Flag = false
				}

				vecPathConds = append(vecPathConds, newCond)
			}
		}

		index ++
	}



	return vecPathConds
}

func collectUnlockedMutex(fn * ssa.Function, pd * analysis.PostDominator) map[string] bool {

	//fmt.Println("call function collectUnlockedMutex")

	mapResult := make(map[string] bool)

	mapLockingOperation := make(map[string] map[ssa.Instruction] bool)
	mapUnlockingOperation := make(map[string] map[ssa.Instruction] bool)

	for _, bb := range fn.Blocks {
		for _, ii := range bb.Instrs {
			if instinfo.IsMutexLock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_mutex"
				if _, ok := mapLockingOperation[strMutexName]; !ok {
					mapLockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapLockingOperation[strMutexName][ii] = true
			} else if instinfo.IsRwmutexLock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_rwmutexW"
				if _, ok := mapLockingOperation[strMutexName]; !ok {
					mapLockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapLockingOperation[strMutexName][ii] = true
			} else if instinfo.IsRwmutexRlock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_rwmutexR"
				if _, ok := mapLockingOperation[strMutexName]; !ok {
					mapLockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapLockingOperation[strMutexName][ii] = true
			} else if instinfo.IsMutexUnlock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_mutex"
				if _, ok := mapUnlockingOperation[strMutexName]; !ok {
					mapUnlockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapUnlockingOperation[strMutexName][ii] = true
			} else if instinfo.IsRwmutexUnlock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_rwmutexW"
				if _, ok := mapUnlockingOperation[strMutexName]; !ok {
					mapUnlockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapUnlockingOperation[strMutexName][ii] = true
			} else if instinfo.IsRwmutexRunlock(ii) {
				strMutexName := instinfo.GetMutexName(ii) + "_rwmutexR"
				if _, ok := mapUnlockingOperation[strMutexName]; !ok {
					mapUnlockingOperation[strMutexName] = make(map[ssa.Instruction] bool)
				}
				mapUnlockingOperation[strMutexName][ii] = true
			}
		}
	}

	for strMutexName, mapLockingOp := range mapLockingOperation {
		if len(mapLockingOp) != 1 {
			continue
		}

		//var mapUnlockingOp map[ssa.Instruction] bool

		if _, ok := mapUnlockingOperation[strMutexName]; !ok {
			continue
		}

		mapUnlockingOp, _ := mapUnlockingOperation[strMutexName]

		if len(mapUnlockingOp) != 1 {
			continue
		}

		var bbLocking *ssa.BasicBlock
		for ii, _ := range mapLockingOp {
			bbLocking = ii.Block()
		}

		var bbUnlocking * ssa.BasicBlock
		for ii, _ := range mapUnlockingOp {
			bbUnlocking = ii.Block()
		}

		//fmt.Println(strMutexName)
		//fmt.Println(strMutexName, bbLocking.Index, bbUnlocking.Index)

		//if strMutexName != "s.locks[j]_rwmutexW" {
		//	continue
		//}

		dominatorLocking := collectDominator(bbLocking)
		dominatorUnlocking := collectDominator(bbUnlocking)

		dominatorIntersection := compDominatorIntersection(dominatorLocking, dominatorUnlocking)

		if len(dominatorIntersection) == 0 {
			continue
		}

		bbLowest := getLowestDominator(dominatorIntersection)

		if bbLowest == nil {
			continue
		}

		//fmt.Println("Lowest common:", bbLowest.Index)

		path1 := getShortestPath(bbLowest, bbLocking)
		path2 := getShortestPath(bbLowest, bbUnlocking)


		vecCond1 := generatePathConditions(path1, pd)
		vecCond2 := generatePathConditions(path2, pd)

		if len(vecCond1) != len(vecCond2) {
			continue
		}

		strConstraints1 := ConvertCondsToContraints(vecCond1)
		strConstraints2 := ConvertCondsToContraints(vecCond2)

		if strConstraints1 == strConstraints2 {
			mapResult[strMutexName] = true
		}
	}

	return mapResult
}


func IsDominatedbyShutdown(inputInst ssa.Instruction) bool {
	for _, bb := range inputInst.Parent().Blocks {
		for _, ii := range bb.Instrs {
			call, ok := ii.(*ssa.Call)
			if !ok {
				continue
			}

			if call.Call.IsInvoke() {
				continue
			}

			fnCall, ok := call.Call.Value.(*ssa.Function)

			if !ok {
				continue
			}

			if (fnCall.Pkg.Pkg.Path() == "os" && fnCall.Name() == "Exit") || (fnCall.Pkg.Pkg.Path() == "syscall" && fnCall.Name() == "Kill") {
				if ii.Block().Dominates(inputInst.Block()) {
					return true
				}
			}
		}
	}

	return false
}

func inspectInst(inputInst ssa.Instruction, isBrutal bool, usingSolver bool, pd * analysis.PostDominator) bool {
	if util.IsFnEnd(inputInst) {
		instPanic, ok := inputInst.(*ssa.Panic)
		if ok {
			panicvalue := instPanic.X
			ii, ok := panicvalue.(ssa.Instruction)
			if ok {
				if strings.Contains(ii.String(), "select") {
					return false
				}
			}
		}

		mapLiveMutex := GetLiveMutex(inputInst)


		if len(mapLiveMutex) == 0 {
			return false
		}


		if usingSolver {
			mapUnlockedMutex := collectUnlockedMutex(inputInst.Parent(), pd)
			newMapLiveMutex := map[string] bool {}

			for strMutexName, _ := range mapLiveMutex {
				if _, ok := mapUnlockedMutex[strMutexName]; ok {
					continue
				}

				newMapLiveMutex[strMutexName] = true
			}

			mapLiveMutex = newMapLiveMutex
		}

		vecDeferredMutex := SearchDeferredUnlock(inputInst)

		if len(mapLiveMutex) > len(vecDeferredMutex) {
			return true
		}

	}
	return false
}


func inspectFunc1(fn * ssa.Function) {

	numFunction ++
	GenKillAnalysis(fn)
	pd := analysis.NewPostDominator(fn)

	for _, bb := range fn.Blocks {
		for _, ii := range bb.Instrs {
			if inspectInst(ii, false, false , pd ) {

				if IsDominatedbyShutdown(ii) {
					continue
				}


				if inspectInst(ii, true, true, pd) {
					flag := false
					for _, bug := range Bugs {

						if bug.Pos() == ii.Pos() && ii.Pos() != token.NoPos {
							flag = true
							break
						}
					}

					//fmt.Println(fn.Name())

					if ! flag {
						Bugs = append(Bugs, ii)
					}

				}
			}
		}
	}
}


func inspectFunc(fn * ssa.Function, isMethod bool) {
	if fn.Blocks == nil {
		//meaning this is external function. You will see a lot of them if you use Ssa_build_packages
		return
	}

	//If this function's name contains lock or unlock, skip this function
	if strInsensitiveContains(fn.Name(), "lock") || strInsensitiveContains(fn.Name(), "unlock") {
		return
	}


	inspectFunc1(fn)

	for _, anonyFN := range fn.AnonFuncs { // loop through all anonymous functions in fn
		inspectFunc1(anonyFN)
	}
}


func Detect() {

	Bugs = [] ssa.Instruction{}

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


		/*
		fn.WriteTo(os.Stdout)

		for _, bb := range fn.Blocks {
			for _, ii := range bb.Instrs {
				loc := (config.Prog.Fset).Position(ii.Pos())

				fmt.Println(ii.String())
				fmt.Println(loc.Filename, loc.Line)

			}
		}
		 */


		if fn.Signature.Recv() == nil {
			inspectFunc(fn, false)
		} else {
			inspectFunc(fn, true)
		}
	}


	mapCombinedIndex := make(map[int] bool)
	mapIndexGroup := make(map[int] map[int] bool)


	/*
	for i, _ := range Bugs {
		Bugs[i].Parent().WriteTo(os.Stdout)
	}



	for i, _ :=  range Bugs {
		for j:= i + 1; j < len(Bugs); j ++ {
			fmt.Printf("%p %p %d\n",Bugs[i].Parent(), Bugs[j].Parent(), Bugs[i].Parent() == Bugs[j].Parent())
		}
	}

	*/

	for i, _ := range Bugs {
		if _, ok := mapCombinedIndex[i]; ok {
			continue
		}
		m1 := make(map[int] bool)
		m1[i] = true
		mapIndexGroup[i] = m1
		mapCombinedIndex[i] = true

		for j := i +1; j < len(Bugs); j ++ {
			if _, ok := mapCombinedIndex[j]; ok {
				continue
			}

			if Bugs[i].Parent() == Bugs[j].Parent() {
				mapIndexGroup[i][j] = true
				mapCombinedIndex[j] = true
			}

		}
	}


	for _, mapIndex := range mapIndexGroup {
		vecToPrint := [] ssa.Instruction {}

		for i, _ := range mapIndex {
			vecToPrint = append(vecToPrint, Bugs[i])
		}

		config.BugIndexMu.Lock()
		config.BugIndex ++
		fmt.Print("----------Bug[")
		fmt.Print(config.BugIndex)
		config.BugIndexMu.Unlock()
		fmt.Print("]----------\n\tType: Missing Unlock \tReason: Unlock operation of a Mutex/RWMutex is missing.\n")
		fmt.Print("\tLocation of multiple buggy instructions:\n")
		output.PrintInsts(vecToPrint)
	}


}