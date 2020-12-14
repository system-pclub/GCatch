package C7A

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

type BackEdge struct {
	src *ssa.BasicBlock
	dst *ssa.BasicBlock
}

type LoopInfo struct {
	status    map[*ssa.BasicBlock]int //0: in progress 1: done
	isLoopBB  map[*ssa.BasicBlock]struct{}
	backEdges []BackEdge
	fn        *ssa.Function
}

func NewLoopInfo(function *ssa.Function) *LoopInfo {
	ret := LoopInfo{
		status:    make(map[*ssa.BasicBlock]int),
		isLoopBB:  make(map[*ssa.BasicBlock]struct{}),
		backEdges: make([]BackEdge, 0),
		fn:        function,
	}
	return &ret
}

func contains(this *map[*ssa.BasicBlock]int, bb *ssa.BasicBlock) bool {
	_, ok := (*this)[bb]
	if ok {
		return true
	} else {
		return false
	}
}

func (this *LoopInfo) dfs(bb *ssa.BasicBlock) {
	this.status[bb] = 0
	for _, succ := range bb.Succs {
		if !contains(&this.status, succ) {
			this.dfs(succ)
		} else if this.status[succ] == 0 {
			this.backEdges = append(this.backEdges, BackEdge{
				src: bb,
				dst: succ,
			})
		}
	}
	this.status[bb] = 1
}

func (this *LoopInfo) revDfs(bb *ssa.BasicBlock) {
	this.status[bb] = 1
	for _, pred := range bb.Preds {
		if !contains(&this.status, pred) {
			this.revDfs(pred)
		}
	}
	this.status[bb] = 1
}

func (this *LoopInfo) Analyze() {
	//for each back edges:
	//clear status map
	//reverse dfs to find loop body bbs.
	this.dfs(this.fn.Blocks[0]) //dfs to find back edges.
	for _, backedge := range this.backEdges {
		this.status = make(map[*ssa.BasicBlock]int)
		this.status[backedge.dst] = 1 //the loop header node
		this.revDfs(backedge.src)
		for k, v := range this.status {
			if k != backedge.dst && v == 1 {
				this.isLoopBB[k] = struct{}{}
			}
		}
		if backedge.src == backedge.dst {
			this.isLoopBB[backedge.src] = struct{}{}
		}
	}
}
