package fatal

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/config"
	"github.com/system-pclub/GCatch/GCatch/output"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa/ssautil"
	"strings"
)

var C8_done_fn []string

func init() {
	C8_done_fn = []string{}
}

func cleanup() {

}

func Detect() {

	cleanup()
	loop_fns()
}

func report(inst ssa.Instruction, parent *ssa.Function) {
	config.BugIndexMu.Lock()
	config.BugIndex++
	fmt.Print("----------Bug[")
	fmt.Print(config.BugIndex)
	config.BugIndexMu.Unlock()
	fmt.Print("]----------\n\tType: API-Fatal \tReason: testing.Fatal()/FailNow()/SkipNow()/... can only be used in test goroutine.\n")
	fmt.Print("\tLocation of call:\n")
	output.PrintIISrc(inst)
}

func loop_fns() {

fn_loop:
	for fn, _ := range ssautil.AllFunctions(config.Prog) {

		if fn == nil {
			continue
		}
		if config.IsPathIncluded(fn.String()) == false {
			continue
		}
		//Actually we don't need to measure the function name, since testing.Fatal() won't be used in normal functions
		//if strings.Contains(fn.Name(),"test") == false && strings.Contains(fn.Name(),"Test") == false {
		//	continue
		//}
		fn_str := fn.String()
		for _, done_fn := range C8_done_fn {
			if done_fn == fn_str {
				continue fn_loop
			}
		}
		C8_done_fn = append(C8_done_fn, fn.String())

		inside_func(fn)
	}
}

func inside_func(fn *ssa.Function) {

	for _, bb := range fn.Blocks {
		for _, inst := range bb.Instrs {
			//p := (config.Prog.Fset).Position(inst.Pos())

			inst_go, ok := inst.(*ssa.Go)
			if !ok {
				continue
			}

			if inst_go.Call.IsInvoke() == true {
				continue
			}

			callee := inst_go.Call.Value
			var interesting_fn *ssa.Function
			switch concrete := callee.(type) {
			case *ssa.Function:
				interesting_fn = concrete
			case *ssa.Builtin:
			case *ssa.MakeClosure:
				var ok bool
				interesting_fn, ok = concrete.Fn.(*ssa.Function)
				if !ok {
					fmt.Println("Warning in C8: Unknown MakeClosure callee in:", fn.String(), "\tinst:", inst)
				}
			default:
				node, ok := config.CallGraph.Nodes[fn]
				if !ok {
					continue
				}
				for _, out := range node.Out {
					if out.Site == inst {
						if out.Callee.Func != nil {
							if strings.Contains(out.Callee.Func.String(), fn.Name()) { // make sure the callee is created in this function, or there will be a lot of FPs
								find_fatal_in_fn(out.Callee.Func, fn)
							}
						}
					}
				}

			}
			if interesting_fn == nil {
				continue
			} else {
				find_fatal_in_fn(interesting_fn, fn)
			}
		}
	}
}

func find_fatal_in_fn(target, parent *ssa.Function) {
	for _, bb := range target.Blocks {
		for _, inst := range bb.Instrs {
			inst_call, ok := inst.(*ssa.Call)
			if !ok {
				continue
			}

			if inst_call.Call.IsInvoke() {
				continue
			}

			callee_fn, ok := inst_call.Call.Value.(*ssa.Function)
			if !ok {
				continue
			}

			if callee_fn.Name() == "Fatal" || callee_fn.Name() == "Fatalf" || callee_fn.Name() == "FailNow" ||
				callee_fn.Name() == "Skip" || callee_fn.Name() == "Skipf" || callee_fn.Name() == "SkipNow" {
				if strings.Contains(callee_fn.Pkg.String(), "package testing") == false {
					continue
				}
				report(inst, parent)
			}

		}
	}
	for _, anony := range target.AnonFuncs {
		find_fatal_in_fn(anony, parent)
	}
}
