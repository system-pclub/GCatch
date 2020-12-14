package search

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/check/sync_check"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

func Find_prim(inst ssa.Instruction) *ssa.Value {
	inst_sync_type := sync_check.Type_sop(inst)
	prim_type := sync_check.Op_to_prim(inst_sync_type)

	switch prim_type {

	case "chan":
		return nil

	case "atomic":
		return nil

	case "atomic_value":
		return nil

	case "goroutine":
		return nil

	case "syncmap":
		return nil

	default:	//including mutex rwmutex waitgroup once cond pool
				//we can these normal primitives
		return find_normal_prim(inst)
	}
}



func case_insensitive_contains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func case_insensitive_equal(s1, s2 string) bool {
	s1, s2 = strings.ToUpper(s1), strings.ToUpper(s2)
	return s1 == s2
}

