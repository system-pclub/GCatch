package bmoc

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/analysis/pointer"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/syncgraph"
	"strconv"
)

func Detect() {
	stPtrResult, vecStOpValue := pointer.AnalyzeAllSyncOp()
	if stPtrResult == nil || vecStOpValue == nil {
		return
	}

	// When the pointer analysis has any uncertain alias relationship, report Not Sure.
	// This is in GCatch/analysis/pointer/utils.go, func mergeAlias()
	vecChannelOri := pointer.WithdrawAllChan(stPtrResult, vecStOpValue)

	vecLockerOri := pointer.WithdrawAllTraditionals(stPtrResult, vecStOpValue) // May delete

	mapDependency := syncgraph.GenDMap(vecChannelOri, vecLockerOri) // May delete

	vecChannel := []*instinfo.Channel{}
	for _, ch := range vecChannelOri {
		if OKToCheck(ch) == true { // some channels may come from SDK like testing. Ignore them
			vecChannel = append(vecChannel, ch)
		}
	}

	vecLocker := []*instinfo.Locker{}
	for _, l := range vecLockerOri {
		if OKToCheckLocker(l) {
			vecLocker = append(vecLocker, l)
		}
	}

	if len(vecChannel) == 0 && len(vecLocker) == 0 { // Definitely no channel/mutex safety of liveness violations
		syncgraph.ReportNoViolation()
		return
	}

	if len(vecChannel) > 0 {
		// Check all channels together. We can just check the first channel from main(), and let all other channels be checked together with it
		Check(vecChannel[0], vecChannel, vecLocker, mapDependency)
	} else {
		Check(vecLocker[0], vecChannel, vecLocker, mapDependency)
	}

}

var countCh int
var countUnbufferBug int
var countBufferBug int

func OKToCheckLocker(locker *instinfo.Locker) (boolCheck bool) {
	boolCheck = false

	if locker.Value == nil {
		return
	}
	if locker.Value.Parent() == nil {
		return
	}

	pkg := locker.Value.Parent().Pkg
	if pkg == nil {
		return
	}
	pkgOfPkg := pkg.Pkg
	if pkgOfPkg == nil {
		return
	}
	if config.IsPathIncluded(pkgOfPkg.Path()) == false {
		return
	}
	return true
}

func OKToCheck(ch *instinfo.Channel) (boolCheck bool) {
	boolCheck = false

	if ch.MakeInst == nil {
		return
	}
	pkg := ch.MakeInst.Parent().Pkg
	if pkg == nil {
		return
	}
	pkgOfPkg := pkg.Pkg
	if pkgOfPkg == nil {
		return
	}
	if config.IsPathIncluded(pkgOfPkg.Path()) == false {
		return
	}

	p := config.Prog.Fset.Position(ch.MakeInst.Pos())
	strChHash := ch.MakeInst.Parent().String() + ch.MakeInst.String() + ch.MakeInst.Name() + strconv.Itoa(p.Line)
	if _, checked := config.MapHashOfCheckedCh[strChHash]; checked {
		return
	}

	boolCheck = true
	config.MapHashOfCheckedCh[strChHash] = struct{}{}
	countCh++
	return
}

func Check(prim interface{}, vecChannel []*instinfo.Channel, vecLocker []*instinfo.Locker, mapDependency map[interface{}]*syncgraph.DPrim) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	syncGraph, err := syncgraph.BuildGraph(prim, vecChannel, vecLocker, mapDependency)
	if err != nil { // Met some error
		if config.Print_Debug_Info {
			fmt.Println("-----count_ch:", countCh)
		}
		fmt.Println("Error when building graph", err.Error())
		syncgraph.ReportNotSure()
		return
	}

	syncGraph.ComputeFnOnOpPath()
	syncGraph.OptimizeBB_V1()

	syncGraph.SetEnumCfg(1, false, true)

	syncGraph.EnumerateAllPathCombinations()

	boolSkip := false
	if primCh, ok := prim.(*instinfo.Channel); ok {
		if primCh.Buffer == instinfo.DynamicSize {
			// If this is a buffered channel with dynamic size and no critical section is found, skip this channel
			fmt.Println("Contains dynamic sized buffered channel.")
			syncgraph.ReportNotSure()
			boolSkip = true
		}
	}

	if boolSkip == false {
		foundBug := syncGraph.CheckWithZ3()
		if foundBug {
			syncgraph.ReportViolation()
		} else {
			syncgraph.ReportNoViolation()
		}
	}
	return
}
