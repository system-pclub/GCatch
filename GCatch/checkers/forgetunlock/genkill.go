package forgetunlock

import (
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/util"
	"golang.org/x/tools/go/ssa"
)

var mapGen map[ssa.Instruction]string
var mapKill map[ssa.Instruction]string
var mapBefore map[ssa.Instruction]map[string]bool
var mapAfter map[ssa.Instruction]map[string]bool

func CompareTwoMaps(map1 map[string]bool, map2 map[string]bool) bool {
	if len(map1) != len(map2) {
		return false
	}

	for s, _ := range map1 {
		if _, ok := map2[s]; !ok {
			return false
		}
	}

	return true
}

func GetLiveMutex(inputInst ssa.Instruction) map[string]bool {
	return mapBefore[inputInst]
}

func InitGenKillMap(inputFn *ssa.Function) {
	for _, bb := range inputFn.Blocks {
		for _, ii := range bb.Instrs {
			if instinfo.IsDefer(ii) {

			} else if instinfo.IsMutexLock(ii) || instinfo.IsRwmutexLock(ii) || instinfo.IsRwmutexRlock(ii) {
				var strMutexName string
				if instinfo.IsMutexLock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_mutex")
				} else if instinfo.IsRwmutexLock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_rwmutexW")
				} else if instinfo.IsRwmutexRlock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_rwmutexR")
				}

				mapGen[ii] = strMutexName

			} else if instinfo.IsMutexUnlock(ii) || instinfo.IsRwmutexUnlock(ii) || instinfo.IsRwmutexRunlock(ii) {

				var strMutexName string
				if instinfo.IsMutexUnlock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_mutex")
				} else if instinfo.IsRwmutexUnlock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_rwmutexW")
				} else if instinfo.IsRwmutexRunlock(ii) {
					strMutexName = string(instinfo.GetMutexName(ii) + "_rwmutexR")
				}
				mapKill[ii] = strMutexName

			} else {

			}
		}
	}
}

func InitBeforeAfterMap(inputFn *ssa.Function) {
	for _, bb := range inputFn.Blocks {
		for _, ii := range bb.Instrs {
			s1 := make(map[string]bool)
			mapBefore[ii] = s1
			s2 := make(map[string]bool)
			mapAfter[ii] = s2
		}
	}

}

func GenKillAnalysis(inputFn *ssa.Function) {
	mapGen = make(map[ssa.Instruction]string)
	mapKill = make(map[ssa.Instruction]string)
	mapBefore = make(map[ssa.Instruction]map[string]bool)
	mapAfter = make(map[ssa.Instruction]map[string]bool)

	InitGenKillMap(inputFn)

	if len(mapGen) == 0 {
		return
	}

	InitBeforeAfterMap(inputFn)

	vecWorkList := make([]ssa.Instruction, 0)

	for _, bb := range inputFn.Blocks {
		for _, ii := range bb.Instrs {
			vecWorkList = append(vecWorkList, ii)
		}
	}

	for len(vecWorkList) > 0 {
		ii := vecWorkList[len(vecWorkList)-1]
		vecWorkList = vecWorkList[:len(vecWorkList)-1]

		prevIIs := util.GetPrevInsts(ii)

		newBefore := make(map[string]bool)

		for _, prevII := range prevIIs {
			for strMutexName, _ := range mapAfter[prevII] {
				newBefore[strMutexName] = true
			}
		}

		//mapBefore[ii] = newBefore

		mapBefore[ii] = make(map[string]bool)

		for i, _ := range newBefore {
			mapBefore[ii][i] = true
		}

		if strMutexName, ok := mapGen[ii]; ok {
			newBefore[strMutexName] = true
		}

		if strMutexName, ok := mapKill[ii]; ok {
			delete(newBefore, strMutexName)

		}

		if !CompareTwoMaps(newBefore, mapAfter[ii]) {
			mapAfter[ii] = newBefore
			for _, pI := range util.GetSuccInsts(ii) {
				vecWorkList = append(vecWorkList, pI)
			}
		}
	}
}
