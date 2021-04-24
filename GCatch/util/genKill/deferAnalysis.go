package genKill

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa/ssautil"
	"github.com/system-pclub/GCatch/GCatch/util"
)

func ComputeDeferMap() (map[ssa.Instruction][]*ssa.Defer, map[*ssa.Defer][]ssa.Instruction) {
	Inst2Defers := make(map[ssa.Instruction][]*ssa.Defer)
	Defer2Insts := make(map[*ssa.Defer][]ssa.Instruction)
	for fn,_ := range ssautil.AllFunctions(config.Prog) {
		computeDeferMapInFunc(fn, Inst2Defers)
	}

	// Delete inst that has no defer (these are added by go's ssa, but can be deleted)
	// Reverse the order of defers, because go will call defers that is executed late
	for inst, defers := range Inst2Defers {
		if len(defers) == 0 {
			delete(Inst2Defers, inst)
			continue
		}
		reverse := []*ssa.Defer{}
		for i := len(defers) - 1; i >= 0; i-- {
			reverse = append(reverse,defers[i])
		}
		Inst2Defers[inst] = reverse
		defers = nil
	}

	if config.MapPrintMod[config.ConstPrintDeferMap] {
		printDeferMap(Inst2Defers)
	}

	for R, Ds := range Inst2Defers {
		for _,D := range Ds {
			Defer2Insts[D] = append(Defer2Insts[D],R)
		}
	}

	return Inst2Defers, Defer2Insts
}

type deferTask struct{
	fnTarget *ssa.Function // the fn we are analyzing
	mapGen map[ssa.Instruction] *ssa.Defer
	mapKill map[ssa.Instruction] *ssa.Defer // always empty
	mapBefore map[ssa.Instruction] []*ssa.Defer
	mapAfter map[ssa.Instruction] []*ssa.Defer
	vecWorkList []ssa.Instruction
}

func (task *deferTask) InitMaps() {
	task.mapGen = make(map[ssa.Instruction] *ssa.Defer)
	task.mapKill = make(map[ssa.Instruction] *ssa.Defer)
	task.mapBefore = make(map[ssa.Instruction] []*ssa.Defer)
	task.mapAfter = make(map[ssa.Instruction] []*ssa.Defer)
}

func (task *deferTask) Clear() {
	task.mapGen = nil
	task.mapKill = nil
	task.mapBefore = nil
	task.mapAfter = nil
	task.vecWorkList = nil
}

func (task *deferTask) InitGenKillMap() {
	for _, bb := range task.fnTarget.Blocks {
		for _, inst := range bb.Instrs {
			inst_as_defer, ok := inst.(*ssa.Defer)
			if ok {
				task.mapGen[inst] = inst_as_defer
			}
		}
	}
}

func (task *deferTask) InitBeforeAfterMap() {
	// no need to init mapBefore and mapAfter
}

func (task *deferTask) Analyze() {
	if task.fnTarget.Name() == "FormatValue" {
		print()
	}
	task.InitMaps()
	task.InitGenKillMap()
	if len(task.mapGen) == 0 {
		return
	}
	task.InitBeforeAfterMap()

	task.vecWorkList = util.GetEntryInsts(task.fnTarget)

	for len(task.vecWorkList) > 0 {
		instTarget := task.vecWorkList[len(task.vecWorkList)-1]
		task.vecWorkList = task.vecWorkList[:len(task.vecWorkList)-1]
		oldAfter, boolExist := task.mapAfter[instTarget]

		newBefore := task.computePrevDefers(instTarget)

		task.mapBefore[instTarget] = newBefore

		newAfter := task.computeAfter(instTarget, newBefore)

		task.mapAfter[instTarget] = newAfter
		if !boolExist || !isVecDefersEqual(newAfter, oldAfter) {
			task.vecWorkList = task.appendNextInstInOrder(instTarget)
		}
	}
}

// if the next inst is the head of a BB, append it to the end of vecWorkList
// if the next inst is not the head of a BB, append it to the beginning of vecWorkList
func (task *deferTask) appendNextInstInOrder(instTarget ssa.Instruction) (result []ssa.Instruction) {

	succInsts := util.GetSuccInsts(instTarget)
	headOfBbInsts := []ssa.Instruction{}
	notHeadOfBbInsts := []ssa.Instruction{}
	for _, succInst := range succInsts {
		if util.IsInstInVec(succInst, task.vecWorkList) {
			continue
		}
		if bb := (succInst).Block(); bb.Instrs[0] == succInst { //succInst is the head of a bb
			headOfBbInsts = append(headOfBbInsts, succInst)
		} else {
			notHeadOfBbInsts = append(notHeadOfBbInsts, succInst)
		}
	}

	for _, notHeadInst := range notHeadOfBbInsts {
		result = append(result, notHeadInst)
	}
	for _,old_inst := range task.vecWorkList {
		result = append(result,old_inst)
	}
	for _,head_inst := range headOfBbInsts {
		result = append(result,head_inst)
	}

	return
}

func (task *deferTask) computePrevDefers(instTarget ssa.Instruction) []*ssa.Defer {
	vecResult := []*ssa.Defer{}
	vecPrevInsts := util.GetPrevInsts(instTarget)
	if len(vecPrevInsts) == 0 { // This is the first time this whole worklist loop is invoked
		// Do nothing
	} else if len(vecPrevInsts) == 1 { // Inherit the defers from the previous inst
		prev_defers, ok := task.mapAfter[vecPrevInsts[0]]
		if ok {
			// copy prev_defers into vecResult
			for _, _defer := range prev_defers {
				vecResult = append(vecResult, _defer)
			}
		} // if !ok, then mapAfter is empty, do nothing
	} else {					//Union the mapAfter of all previous inst
		mapPrevDefer := make(map[*ssa.Defer]struct{})
		for _, prevInst := range vecPrevInsts {
			vecDefers, ok := task.mapAfter[prevInst]
			if !ok { // if !ok, then mapAfter is empty, don't store
				continue
			}
			for _, _defer := range vecDefers {
				mapPrevDefer[_defer] = struct{}{}
			}
		}
		for _defer, _ := range mapPrevDefer {
			vecResult = append(vecResult,_defer)
		}
		mapPrevDefer = nil
	}
	return vecResult
}

func (task *deferTask) computeAfter(instTarget ssa.Instruction, newBefore []*ssa.Defer) []*ssa.Defer {
	vecResult := []*ssa.Defer{}
	for _, _defer := range newBefore {
		vecResult = append(vecResult, _defer)
	}

	genDefer := task.mapGen[instTarget]

	boolAlreadyIn := false
	for _, beforeDefer := range newBefore {
		if beforeDefer == genDefer {
			boolAlreadyIn = true
			break
		}
	}
	if boolAlreadyIn || genDefer == nil {
		return vecResult
	} else {
		vecResult = append(vecResult, genDefer)
		return vecResult
	}
}


func computeDeferMapInFunc(fn *ssa.Function, R2D map[ssa.Instruction] []*ssa.Defer) {
	var task deferTask
	task.fnTarget = fn

	task.Analyze()

	mapBefore := task.mapBefore
	for inst, defers := range mapBefore {
		switch inst.(type) {
		case *ssa.RunDefers, *ssa.Panic:
			R2D[inst] = defers
		case *ssa.Call:
			if util.IsInstCallFatal(inst) {
				R2D[inst] = defers
			}
		}
	}

	task.Clear()
}

func printDeferMap(Inst2Defers map[ssa.Instruction][]*ssa.Defer) {
	count := 0
	for rundefer, defers := range Inst2Defers {
		fmt.Println("--------NO.",count)
		fmt.Println("----Location of rundefer")
		output.PrintIISrc(rundefer)
		fmt.Println("----Location of defer")
		for _, _defer := range defers {
			output.PrintIISrc(_defer)
			if _defer.Call.IsInvoke() {
				fmt.Println("\tDefering:", _defer.Call.Method.String()," of interface ", _defer.Call.Value.Type().String())
			} else {
				fmt.Println("\tDefering:", _defer.Call.Value.String())
			}
		}
		count++

		if count%10 == 0 {
			output.WaitForInput()
		}
	}
}

func isVecDefersEqual(vec1, vec2 []*ssa.Defer) bool {
	map1 := make(map[*ssa.Defer]struct{})
	for _, elem := range vec1 {
		map1[elem] = struct{}{}
	}
	map2 := make(map[*ssa.Defer]struct{})
	for _, elem := range vec2 {
		map2[elem] = struct{}{}
	}

	if len(map1) != len(map2) {
		return false
	} else {
		for elem, _ := range map1 {
			_, exist := map2[elem]
			if !exist {
				return false
			}
		}
	}

	return true

	//// reflect.DeepEqual will compare two values. Comparison of slice, maps or pointers will be handled recursively
	//// Note that we can't directly compare slices here, because the order also matters
	//return reflect.DeepEqual(map1, map2)
}
