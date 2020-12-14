package prepare

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa/ssautil"
	"strings"
)

func List_all_methods() []*ssa.Function {
	methodset := *new([]*ssa.Function)

	fns_in_prog := ssautil.AllFunctions(global.Prog)
	for fn_in_prog, _ := range fns_in_prog { // a cumbersome loop, looping through all functions in the program
		method_prefix := ")."
		var str string
		if fn_in_prog.Pkg == nil {
			str = fn_in_prog.String()
		} else {
			if Is_path_include(fn_in_prog.Pkg.Pkg.Path()) == false {
				continue
			}
			str = fn_in_prog.RelString(fn_in_prog.Pkg.Pkg)
		}
		if strings.Contains(str, method_prefix) {
			//this function is a method of mem_as_type, and it is in pkg
			methodset = append(methodset, fn_in_prog)
		}
	}

	var result []*ssa.Function = *new([]*ssa.Function)
	for _,method := range methodset {
		if method.Pkg != nil && method.Synthetic == "" {
			result = append(result,method)
		}
	}

	return result
}
