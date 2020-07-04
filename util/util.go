package util

import "github.com/system-pclub/GCatch/tools/go/ssa"

func IsInstInVec(inst ssa.Instruction, vec []ssa.Instruction) bool {
	for _, elem := range vec {
		if elem == inst {
			return true
		}
	}
	return false
}
