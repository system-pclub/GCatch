package analysis

import "github.com/system-pclub/GCatch/GCatch/tools/go/ssa"

type LoopAnalysis struct {
	FN *ssa.Function
	VecBackedge []*Edge
	MapLoopHead2BodyBB map[*ssa.BasicBlock][]*ssa.BasicBlock
	MapBodyBb2LoopHead map[*ssa.BasicBlock][]*ssa.BasicBlock
}

type Edge struct {
	Pred *ssa.BasicBlock
	Succ *ssa.BasicBlock
}

const Unmarked, InProgress, Done, Visited = 0, 1, 2, 3

func NewLoopAnalysis(fn *ssa.Function) *LoopAnalysis {
	if fn == nil || len(fn.Blocks) == 0 {
		return nil
	}
	loopAnalysis := &LoopAnalysis{
		FN:                 fn,
		VecBackedge:        []*Edge{},
		MapLoopHead2BodyBB: make(map[*ssa.BasicBlock][]*ssa.BasicBlock),
		MapBodyBb2LoopHead: make(map[*ssa.BasicBlock][]*ssa.BasicBlock),
	}

	// Depth first search to find all backedges
	firstBB := fn.Blocks[0]
	mapBB2Status := make(map[*ssa.BasicBlock]int)
	for _, bb := range fn.Blocks {
		mapBB2Status[bb] = Unmarked
	}
	dfs(firstBB, mapBB2Status, loopAnalysis)

	mapLoopHead2BodyBB := make(map[*ssa.BasicBlock][]*ssa.BasicBlock)
	mapBodyBb2LoopHead := make(map[*ssa.BasicBlock][]*ssa.BasicBlock)

	// For each backedge, find all loop headers and corresponding loop bodies
	for _, backedge := range loopAnalysis.VecBackedge {
		// check if there is a loop.
		// A loop (called a natural loop) is identified for every backedge from a node Y to a node X such that X dominates Y
		var header, aNode *ssa.BasicBlock
		header = backedge.Succ // X
		aNode = backedge.Pred // Y
		if header.Dominates(aNode) == false {
			// Not a loop
			continue
		}
		for _, bb := range fn.Blocks {
			mapBB2Status[bb] = Unmarked
		}
		mapBB2Status[header] = Visited
		mapVisitedNode := make(map[*ssa.BasicBlock]struct{})
		backwardsDfs(aNode, mapBB2Status, mapVisitedNode)
		for visited, _ := range mapVisitedNode {
			mapLoopHead2BodyBB[header] = append(mapLoopHead2BodyBB[header], visited)
			mapBodyBb2LoopHead[visited] = append(mapBodyBb2LoopHead[visited], header)
		}
	}

	// delete redundant bbs
	for header, vecBodyBB := range mapLoopHead2BodyBB {
		mapBodyBB := make(map[*ssa.BasicBlock]struct{})
		for _, bb := range vecBodyBB {
			mapBodyBB[bb] = struct{}{}
		}
		for bb, _ := range mapBodyBB {
			loopAnalysis.MapLoopHead2BodyBB[header] = append(loopAnalysis.MapLoopHead2BodyBB[header], bb)
		}
		mapBodyBB = nil
	}
	for body, vecHeader := range mapBodyBb2LoopHead {
		mapHeader := make(map[*ssa.BasicBlock]struct{})
		for _, bb := range vecHeader {
			mapHeader[bb] = struct{}{}
		}
		for bb, _ := range mapHeader {
			loopAnalysis.MapBodyBb2LoopHead[body] = append(loopAnalysis.MapBodyBb2LoopHead[body], bb)
		}
		mapHeader = nil
	}
	mapLoopHead2BodyBB = nil
	mapBodyBb2LoopHead = nil

	return loopAnalysis
}

func backwardsDfs(bb *ssa.BasicBlock, mapBB2Status map[*ssa.BasicBlock]int, mapVisitedNode map[*ssa.BasicBlock]struct{}) {
	mapBB2Status[bb] = Visited
	mapVisitedNode[bb] = struct{}{}
	for _, pred := range bb.Preds {
		mark, _ := mapBB2Status[pred]
		if mark != Visited {
			backwardsDfs(pred, mapBB2Status, mapVisitedNode)
		}
	}
}

func dfs(bb *ssa.BasicBlock, mapBB2Status map[*ssa.BasicBlock]int, loopAnalysis *LoopAnalysis) {
	mapBB2Status[bb] = InProgress
	for _, suc := range bb.Succs {
		mark, _ := mapBB2Status[suc]
		if mark == Unmarked {
			dfs(suc, mapBB2Status, loopAnalysis)
		} else if mark == InProgress {
			newBackEdge := &Edge{
				Pred: bb,
				Succ: suc,
			}
			loopAnalysis.VecBackedge = append(loopAnalysis.VecBackedge, newBackEdge)
		}
	}
	mapBB2Status[bb] = Done
}
