package util

import "github.com/system-pclub/GCatch/GCatch/tools/go/ssa"

func IsInstCallFatal(inst ssa.Instruction) bool {
	list_fatal := [6]string{"Fatal","Fatalf","FailNow","Skip","Skipf","SkipNow"}

	call, ok := inst.(ssa.CallInstruction)
	if ok && !call.Common().IsInvoke() {
		callee,ok := call.Common().Value.(*ssa.Function)
		if ok {
			if callee.Pkg != nil && callee.Pkg.Pkg != nil {
				if callee.Pkg.Pkg.Name() == "testing" {
					name := callee.Name()
					for _,fatal := range list_fatal {
						if fatal == name {
							return true
						}
					}
				}
			}
		}
	}

	return false
}
