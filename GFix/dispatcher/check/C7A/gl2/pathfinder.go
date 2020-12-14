package gl2

import "github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"

type PathFinder struct {
	status map[*ssa.BasicBlock]int
}

func NewPathFinder() *PathFinder {
	ret := PathFinder{status: map[*ssa.BasicBlock]int{}}
	return &ret
}

func contains(this map[*ssa.BasicBlock]int, bb *ssa.BasicBlock) bool {
	_, ok := this[bb]
	if ok {
		return true
	} else {
		return false
	}
}

func (this *PathFinder) revDfs(bb *ssa.BasicBlock) {
	this.status[bb] = 1
	for _, pred := range bb.Preds {
		if !contains(this.status, pred) {
			this.revDfs(pred)
		}
	}
	this.status[bb] = 1
}

func (this *PathFinder) IsReachableToEntry(obstacles []*ssa.BasicBlock, startingPoints []*ssa.BasicBlock, entry *ssa.BasicBlock) bool {
	for _, startingPoint := range startingPoints {

		this.status = make(map[*ssa.BasicBlock]int)
		this.revDfs(startingPoint)
		isReachableAtTheBeginning := contains(this.status, entry)
		for _, obstacle := range obstacles {
			this.status[obstacle] = 1
		}
		this.revDfs(startingPoint)
		if !contains(this.status, entry) && isReachableAtTheBeginning {
			return false
		}
	}
	return true
}
