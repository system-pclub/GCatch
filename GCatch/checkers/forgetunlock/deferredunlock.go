package forgetunlock

import (
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"golang.org/x/tools/go/ssa"
)

func SearchDeferredUnlock(inputInst ssa.Instruction) []string {

	mapResults := make(map[string]bool)

	for _, inst := range inputInst.Block().Instrs {
		if inst == inputInst {
			break
		}

		strInst := inst.String()

		if instinfo.IsDefer(inst) && (strInsensitiveContains(strInst, "unlock") || strInsensitiveContains(strInst, "runlock")) {
			strMutexName := instinfo.GetMutexName(inst)
			_, ok := mapResults[strMutexName]

			if !ok {
				mapResults[strMutexName] = true
			}
		} else if instinfo.IsDefer(inst) {
			flag := false
			var closure *ssa.MakeClosure
			var ok1 bool

			for _, op := range inst.Operands(nil) {
				closure, ok1 = (*op).(*ssa.MakeClosure)

				if ok1 {
					flag = true
					break
				}
			}

			if !flag {
				continue
			}

			closureFN, ok := closure.Fn.(*ssa.Function)

			if !ok {
				continue
			}

			for _, bb := range closureFN.Blocks {
				for _, i := range bb.Instrs {
					call, ok := i.(*ssa.Call)
					if !ok {
						continue
					}

					if !call.Call.IsInvoke() {
						if call.Call.Value.Name() != "Unlock" {
							continue
						}
					} else {
						if call.Call.Method.Name() != "Unlock" {
							continue
						}
					}

					strMutexName := instinfo.GetMutexName(i)

					_, ok = mapResults[strMutexName]

					if !ok {
						mapResults[strMutexName] = true
					}
				}
			}

		}

	}

	visitedBB := make(map[*ssa.BasicBlock]bool)
	visitedBB[inputInst.Block()] = true

	for _, prev := range inputInst.Block().Preds {
		_, ok := visitedBB[prev]
		if ok {
			continue
		}

		searchDeferredUnlockBB(mapResults, visitedBB, prev)
	}

	results := make([]string, 0, len(mapResults))

	for k := range mapResults {
		results = append(results, k)
	}

	return results
}

func searchDeferredUnlockBB(mapResults map[string]bool, visitedBB map[*ssa.BasicBlock]bool, BB *ssa.BasicBlock) {
	for _, inst := range BB.Instrs {
		strInst := inst.String()

		if instinfo.IsDefer(inst) && (strInsensitiveContains(strInst, "unlock") || strInsensitiveContains(strInst, "runlock")) {
			strMutexName := instinfo.GetMutexName(inst)
			_, ok := mapResults[strMutexName]

			if !ok {
				mapResults[strMutexName] = true
			}
		} else if instinfo.IsDefer(inst) {
			flag := false
			var closure *ssa.MakeClosure
			var ok1 bool

			for _, op := range inst.Operands(nil) {
				closure, ok1 = (*op).(*ssa.MakeClosure)

				if ok1 {
					flag = true
					break
				}
			}

			if !flag {
				continue
			}

			closureFN, ok := closure.Fn.(*ssa.Function)

			if !ok {
				continue
			}

			for _, bb := range closureFN.Blocks {
				for _, i := range bb.Instrs {
					call, ok := i.(*ssa.Call)
					if !ok {
						continue
					}

					if !call.Call.IsInvoke() {
						if call.Call.Value.Name() != "Unlock" {
							continue
						}
					} else {
						if call.Call.Method.Name() != "Unlock" {
							continue
						}
					}

					strMutexName := instinfo.GetMutexName(i)

					_, ok = mapResults[strMutexName]

					if !ok {
						mapResults[strMutexName] = true
					}
				}
			}
		}

		visitedBB[BB] = true

		for _, prev := range BB.Preds {
			_, ok := visitedBB[prev]
			if ok {
				continue
			}

			searchDeferredUnlockBB(mapResults, visitedBB, prev)
		}
	}
}
