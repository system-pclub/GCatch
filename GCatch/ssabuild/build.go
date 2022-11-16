package ssabuild

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// Our traditional way to build a program. Here strPath should be what's after "/src/", like "github.com/docker/docker"
func BuildWholeProgramTrad(strPath string, boolForce bool, boolShowError bool) (*ssa.Program, []*ssa.Package, bool, string) {
	strMsg := "suc"
	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Tests: true}
	initialPackage, err := packages.Load(cfg, strPath) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		strMsg = "load_err"
		fmt.Println(err)

		return nil, nil, false, strMsg
	}

	if packages.PrintErrors(initialPackage) > 0 { //To ignore building errors, you can comment out a line in this function
		strMsg = "type_err"
		return nil, nil, false, strMsg
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

// A Beta functionality: use go.mod to build a program
// For this, we switched to a newer version of package "golang.org/x/tools/go/packages"
// We are using a copied and slightly modified version downloaded on 11/15/2021
// strModulePath is like go.etcd.io
// strModAbsPath is like /home/you/stubs/etcd
func BuildWholeProgramGoMod(strModulePath string, boolForce bool, boolShowError bool, strModAbsPath string) (*ssa.Program, []*ssa.Package, bool, string) {
	os.Setenv("GO111MODULE", "on")
	strMsg := "suc"
	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Tests: true, Dir: strModAbsPath}
	initialPackage, err := packages.Load(cfg, strModulePath) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		strMsg = "load_err"
		fmt.Println(err)

		return nil, nil, false, strMsg
	}

	if packages.PrintErrors(initialPackage) > 0 { //To ignore building errors, you can comment out a line in this function
		strMsg = "type_err"
		return nil, nil, false, strMsg
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
