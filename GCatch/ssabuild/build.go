package ssabuild

import (
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa/ssautil"
	"github.com/system-pclub/GCatch/GCatch/tools/go/packages"
)

func BuildWholeProgram(strPath string, boolForce bool, boolShowError bool) (*ssa.Program, []*ssa.Package, bool, string) {
	strMsg := "suc"
	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Tests: true, }
	initialPackage, err := packages.Load(cfg, strPath) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		strMsg = "load_err"
		if boolShowError {
			fmt.Println(err)
		}

		return nil, nil, false, strMsg
	}

	if packages.PrintErrors(initialPackage, boolShowError) > 0 { //To ignore building errors, you can comment out a line in this function
		strMsg = "type_err"

		if boolForce == false {
			return nil, nil, false, strMsg
		}
	}

	//if packages.PrintErrors(initialPackage, boolShowError) > 0 && boolForce == false {
	//	return nil, nil, false, strMsg
	//}

	// Create SSA packages for all well-typed packages.
	//fmt.Println("build_mode:", "NaiveForm")                   //|BuildSerially|PrintPackage
	prog, pkgs := ssautil.AllPackages(initialPackage, ssa.NaiveForm) //|ssa.BuildSerially|ssa.PrintPackages

	// Build SSA code for the whole program.
	prog.Build()

	for _, p := range pkgs {
		if p != nil {
			p.Build()
		}
	}

	return prog, pkgs, true, strMsg
}


