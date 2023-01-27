package bmoc

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/analysis/pointer"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/instinfo"
	"github.com/system-pclub/GCatch/GCatch/syncgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strings"
)

func Detect() {
	ptrAnalysisResult, syncOps := pointer.AnalyzeAllSyncOp()
	if ptrAnalysisResult == nil || syncOps == nil {
		return
	}

	// When the pointer analysis has any uncertain alias relationship, report Not Sure.
	// This is in GCatch/analysis/pointer/utils.go, func mergeAlias()
	vecChannelOri := pointer.GetChanOps(ptrAnalysisResult, syncOps)

	vecLockerOri := pointer.GetTraditionalOps(ptrAnalysisResult, syncOps) // May delete

	mapDependency := syncgraph.GenDMap(vecChannelOri, vecLockerOri) // May delete

	vecChannel := []*instinfo.Channel{}
	for _, ch := range vecChannelOri {
		if OKToCheck(ch) { // some channels may come from SDK like testing. Ignore them
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

	//p := config.Prog.Fset.Position(ch.MakeInst.Pos())
	//strChHash := ch.MakeInst.Parent().String() + ch.MakeInst.String() + ch.MakeInst.Name() + strconv.Itoa(p.Line)
	//util.Debugfln("strChHash = %s", strChHash)
	//if _, checked := config.MapHashOfCheckedCh[strChHash]; checked {
	//	util.Debugfln("checked. return.")
	//	return
	//}

	boolCheck = true
	//config.MapHashOfCheckedCh[strChHash] = struct{}{}
	//countCh++
	return
}

func Check(prim interface{}, vecChannel []*instinfo.Channel, vecLocker []*instinfo.Locker, mapDependency map[interface{}]*syncgraph.DPrim) {
	defer func() {
		if config.RecoverFromError {
			if r := recover(); r != nil {
				return
			}
		}
	}()

	syncGraph, err := syncgraph.BuildGraph(prim, vecChannel, vecLocker, mapDependency)
	if err != nil { // Met some error
		if config.Print_Debug_Info {
			fmt.Println("-----count_ch:", countCh)
		}
		fmt.Println("Error when building graph: ", err.Error())
		syncgraph.ReportNotSure()
		return
	}

	syncGraph.ComputeFnOnOpPath()
	syncGraph.OptimizeBB_V1()

	syncGraph.SetEnumCfg(2, false, true)

	syncGraph.EnumerateAllPathCombinations()

	for _, ch := range vecChannel {
		// 1. Abort when the channel has dynamic buffer size that we don't know during compile
		if ch.Buffer == instinfo.DynamicSize {
			// If this is a buffered channel with dynamic size and no critical section is found, skip this channel
			fmt.Println("Warning: Contains dynamic sized buffered channel.")
			syncgraph.ReportNotSure()
			return
		}
		//// 2. (Deprecated: this is too conservative and may through away bugs we could detect)
		////Abort if the channel is used in a such a loop that we can't analyze the loop condition
		//for _, recv := range ch.Recvs {
		//	if recvInstUnOp, ok := recv.ChOp.Inst.(*ssa.UnOp); ok {
		//		if recvInstUnOp.Block().Comment == "rangechan.loop" {
		//			fmt.Println("Warning: Contains channel used in such a loop that we can't analyze how many iterations will be executed.")
		//			syncgraph.ReportNotSure()
		//			return
		//		}
		//	}
		//}
		// 3. Abort if any related channel (in select) is from time.Ticker, since developers may not think this is blocking
		// iterate over ops because the ticker.C may not be in vecChannel
		for _, send := range ch.Sends {
			if IsChOpInvolveTime(send) {
				fmt.Println("Warning: Contains channel from time.Ticker, so they may not be considered blocking.")
				syncgraph.ReportNotSure()
				return
			}
		}
		for _, recv := range ch.Recvs {
			if IsChOpInvolveTime(recv) {
				fmt.Println("Warning: Contains channel from time.Ticker, so they may not be considered blocking.")
				syncgraph.ReportNotSure()
				return
			}
		}
	}

	//syncGraph.PrintAllPathCombinations()
	foundBug := syncGraph.CheckWithZ3()
	if foundBug {
		syncgraph.ReportViolation()
	} else {
		syncgraph.ReportNoViolation()
	}
}

func IsChOpInvolveTime(op instinfo.ChanOp) bool {
	if instSelect, ok := op.Instr().(*ssa.Select); ok {
		for _, state := range instSelect.States {
			// Loop for state.Chan to be (*A.C), where A's type is time.Ticker
			if chUnOp, ok := state.Chan.(*ssa.UnOp); ok {
				if field, ok := chUnOp.X.(*ssa.FieldAddr); ok {
					if fieldUnOp, ok := field.X.(*ssa.UnOp); ok {
						if structAlloc, ok := fieldUnOp.X.(*ssa.Alloc); ok {
							strType := structAlloc.Type().String()
							if strings.Contains(strType, "time.") {
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
