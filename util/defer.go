package util

import "github.com/system-pclub/GCatch/tools/go/ssa"

func IsFirstDefer(ii ssa.Instruction) bool {

	if len(ii.Parent().Blocks) == 0 {
		return false
	}

	bb := ii.Parent().Blocks[0]

	if bb != ii.Block() {
		return false
	}

	for _, i := range bb.Instrs {
		if _, ok := i.(* ssa.Defer); ok {
			if i == ii {
				return true
			}
		}
	}

	return false
}
