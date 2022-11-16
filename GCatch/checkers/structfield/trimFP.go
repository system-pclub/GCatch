package structfield

import (
	"strings"

	"github.com/system-pclub/GCatch/GCatch/config"
	"golang.org/x/tools/go/ssa"
)

func FP_loop_pkg(call_insts []ssa.Instruction, target_inst ssa.Instruction) []ssa.Instruction {
	target_pkg := target_inst.Parent().Pkg
	vecAllMethods := List_all_methods()

	for mem_name, _ := range target_pkg.Members { //loop through all members; the member may be a func or a type; if it is type, loop through all its methods

		//check if this member is a type
		mem_as_type := target_pkg.Type(mem_name)
		if mem_as_type != nil {
			//This member is a type

			for _, method := range vecAllMethods {
				if method == nil || method.Pkg != target_pkg {
					continue
				}
				method_prefix_1 := "(*" + target_pkg.Pkg.Path() + "." + mem_name + ")."
				method_prefix_2 := "(" + target_pkg.Pkg.Path() + "." + mem_name + ")."
				if strings.Contains(method.String(), method_prefix_1) || strings.Contains(method.String(), method_prefix_2) {

					if method.Pkg != nil {
						if method.Pkg != target_pkg {
							continue
						}
					}
					//this function is a method of mem_as_type, and it is in pkg
					call_insts = FP_inside_func(method, call_insts, target_inst)
				}
			}
		}

		//check if this member is a function
		mem_as_func := target_pkg.Func(mem_name)
		if mem_as_func != nil {
			//this member is a function
			call_insts = FP_inside_func(mem_as_func, call_insts, target_inst)
		}

	}

	return call_insts
}

func FP_inside_func(fn *ssa.Function, call_insts []ssa.Instruction, target_inst ssa.Instruction) []ssa.Instruction {
	target_position := (config.Prog.Fset).Position(target_inst.Pos())
	_ = target_position
	target_fn_name := target_inst.Parent().Name()

	if fn.Blocks == nil { //meaning this is external function. You will see a lot of them if you use Ssa_build_packages
		return call_insts
	}

	for _, bb := range fn.Blocks { // loop all BBs

		insts := bb.Instrs
		for _, inst := range insts { // loop all instructions

			inst_as_call, ok := inst.(*ssa.Call)
			if !ok {
				continue
			}
			call_common := inst_as_call.Call
			var inst_fn_name string
			if call_common.IsInvoke() == true { //This is a call to a method of an interface
				inst_fn_name = call_common.Method.Name()
			} else { //This is a normal static call
				inst_fn, ok := call_common.Value.(*ssa.Function)
				if !ok {
					continue
				}
				inst_fn_name = inst_fn.Name()
			}
			if target_fn_name == inst_fn_name {
				call_insts = append(call_insts, inst)
			}

		} // end of instruction loop
	} //end of BB loop

	return call_insts
}
