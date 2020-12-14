package search

import "github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"

func Is_fn_in_slice(target *ssa.Function, slice []*ssa.Function) bool {
	for _,fn := range slice {
		if fn == target {
			return true
		}
	}
	return false
}

func Is_inst_in_slice(target ssa.Instruction, slice []ssa.Instruction) bool {
	for _,inst := range slice {
		if inst == target {
			return true
		}
	}
	return false
}
