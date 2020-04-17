package analysis

import (
	"fmt"
	"github.com/system-pclub/gochecker/tools/go/ssa"
)

type PostDominator struct {
	FN * ssa.Function
	mapBefore map[*ssa.BasicBlock] map[*ssa.BasicBlock] bool
	mapAfter map[*ssa.BasicBlock] map[*ssa.BasicBlock] bool
}



func NewPostDominator(fn * ssa.Function) * PostDominator {
	pd := new(PostDominator)
	pd.FN = fn
	pd.mapBefore = make(map[* ssa.BasicBlock] map[*ssa.BasicBlock] bool)
	pd.mapAfter = make(map[* ssa.BasicBlock] map[*ssa.BasicBlock] bool)

	pd.initBeforeAfterMap()
	pd.conductAnalysis()

	return pd
}

func (pd * PostDominator) initBeforeAfterMap() {
	for _, bb := range pd.FN.Blocks {
		pd.mapBefore[bb] = make(map[* ssa.BasicBlock] bool)
		pd.mapAfter[bb] = make(map[* ssa.BasicBlock] bool)
	}
}

func intersectMap(m1 map[* ssa.BasicBlock] bool, m2 map[* ssa.BasicBlock] bool) map[* ssa.BasicBlock] bool {
	mapResult := map[* ssa.BasicBlock] bool {}

	for bb, _ := range m1 {
		if _, ok := m2[bb]; ok {
			mapResult[bb] = true
		}
	}

	return mapResult
}

func compareMap(m1 map[* ssa.BasicBlock] bool, m2 map[* ssa.BasicBlock] bool ) bool {
	if len(m1) != len(m2) {
		return false
	}

	for bb, _ := range m1 {
		if _, ok := m2[bb]; !ok {
			return false
		}
	}

	return true
}

func (pd * PostDominator) conductAnalysis() {
	vecWorkList := []* ssa.BasicBlock {}
	for _, bb := range pd.FN.Blocks {
		vecWorkList = append(vecWorkList, bb)
	}

	for len(vecWorkList) > 0 {
		bb := vecWorkList[len(vecWorkList) - 1]
		vecWorkList = vecWorkList[:len(vecWorkList)-1]

		newAfter := make(map[* ssa.BasicBlock] bool)

		if len(bb.Succs) > 0 {
			for b, _ := range pd.mapBefore[bb.Succs[0]] {
				newAfter[b] = true
			}
			for _, succ := range bb.Succs[1:] {
				newAfter = intersectMap(newAfter, pd.mapBefore[succ])
			}
		}

		pd.mapAfter[bb] = make(map[* ssa.BasicBlock] bool)
		for b, _ := range newAfter {
			pd.mapAfter[bb][b] = true
		}

		newAfter[bb] = true

		if !compareMap(pd.mapBefore[bb], newAfter) {
			pd.mapBefore[bb] = newAfter

			for _, pred := range bb.Preds {
				vecWorkList = append(vecWorkList, pred)
			}
		}
	}
}


func (pd * PostDominator) Print() {
	for _, bb := range pd.FN.Blocks {
		fmt.Printf("%d: ", bb.Index)

		for b, _ := range pd.mapAfter[bb] {
			fmt.Printf("%d ", b.Index)
		}
		fmt.Println()
	}
}



func (pd * PostDominator) Dominate(b1 * ssa.BasicBlock, b2 * ssa.BasicBlock) bool {
	if _, ok := pd.mapAfter[b2][b1]; ok {
		return true
	}

	return false
}


