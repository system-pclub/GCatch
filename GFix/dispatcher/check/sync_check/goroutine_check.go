package sync_check

import (
	//"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

func Is_make_goroutine(inst ssa.Instruction) bool {
	inst_asGo, ok := inst.(*ssa.Go)
	if ok {
		//fmt.Println("This instruction is a ssa.Go:", inst,"\t")
		//var fn *ssa.Function
		switch val := inst_asGo.Call.Value.(type) {
		case *ssa.Function:
			_ = val
		//	fn = val
		//	fmt.Println("normal Func:",fn)
			return true
		case *ssa.MakeClosure:
			_ = val
		//	fn = val.Fn.(*ssa.Function)
		//	fmt.Println("anonymous Func:",fn)
			return true
		default:
			return false
		}
	} else {
		return false
	}
}