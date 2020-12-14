package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"go/token"
	"reflect"
)

type SuccBasicBlockFinder struct {
	entry  *ssa.BasicBlock
	status map[*ssa.BasicBlock]int //0: in progress 1: done
}

func NewSuccBasicBlockFinder() SuccBasicBlockFinder {
	return SuccBasicBlockFinder{
		status: make(map[*ssa.BasicBlock]int),
	}
}

func (this *SuccBasicBlockFinder) dfs(bb *ssa.BasicBlock, prt *ssa.BasicBlock) {
	if prt != nil {
		this.status[bb] = 1 //The entry function may still not marked
	}
	for _, succ := range bb.Succs {
		if !contains(&this.status, succ) {
			this.dfs(succ, bb)
		}
	}
}

func (this *SuccBasicBlockFinder) Analyze(startBB *ssa.BasicBlock) []*ssa.BasicBlock {
	ret := make([]*ssa.BasicBlock, 0)
	this.entry = startBB
	this.dfs(startBB, nil)
	for k, _ := range this.status {
		ret = append(ret, k)
	}
	return ret
}

//end defining SuccBasicBlockFinder

func getLineNo(fset *token.FileSet, inst ssa.Instruction) int {
	pos := inst.Pos()
	return fset.Position(pos).Line
}

func getFileName(fset *token.FileSet, inst ssa.Instruction) string {
	pos := inst.Pos()
	return fset.Position(pos).Filename
}

func printSSAByBB(bb *ssa.BasicBlock) {
	for _, ins := range bb.Instrs {
		fmt.Print("    ")
		printSSAIns(ins)
	}
}

func printSSAIns(ins ssa.Instruction) {
	value, ok := ins.(ssa.Value)
	if ok {
		fmt.Print(value.Name(), "=")
	}
	fmt.Println(ins.String(), reflect.TypeOf(ins))
}
