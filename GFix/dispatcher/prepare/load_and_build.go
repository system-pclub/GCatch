package prepare

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/packages"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa/ssautil"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//Note: don't use Ssa_build_packages! It will lose a lot of functions, and lose BB in external functions
func Ssa_build_packages(path string) []*ssa.Package {
	// Load, parse, and type-check the initial packages.
	cfg := &packages.Config{Mode: packages.LoadSyntax}
	initial, err := packages.Load(cfg, path) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		fmt.Println("error1: error after packages.Load")
		log.Fatal(err)
	}

	// Stop if any package had errors.
	// This step is optional; without it, the next step
	// will create SSA for only a subset of packages.
	//if packages.PrintErrors(initial) > 0 {  //I have commented out a line in this function
	//	fmt.Println("error2: packages.PrintErrors > 0 !")
	//	log.Fatalf("packages contain errors")
	//}

	// Create SSA packages for all well-typed packages.
	fmt.Println("build_mode:", "NaiveForm")
	_, pkgs := ssautil.Packages(initial, ssa.NaiveForm) //note: pkgs only contains initial packages. The number of initial packages = the number of paths in  packages.Load(cfg, path...)

	for _, pkg := range pkgs {
		pkg.Build()
	}

	return pkgs
}

func Ssa_build_wholeprogam(path string, force bool, show_compile_error bool) (*ssa.Program, []*ssa.Package, string) {

	suc := "suc"
	println(path)
	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Tests: true}
	initial, err := packages.Load(cfg, path) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		suc = "load_err"
		if show_compile_error {
			fmt.Println(err)
		}

		return nil, nil, suc
	}

	if packages.PrintErrors(initial, show_compile_error) > 0 { //To ignore building errors, you can commented out a line in this function
		suc = "type_err"
	}

	if packages.PrintErrors(initial, show_compile_error) > 0 && force == false {
		return nil, nil, suc
		//log.Fatalf("packages contain errors")
	}

	// Create SSA packages for all well-typed packages.
	//fmt.Println("build_mode:", "NaiveForm")                   //|BuildSerially|PrintPackage
	prog, pkgs := ssautil.AllPackages(initial, ssa.NaiveForm) //|ssa.BuildSerially|ssa.PrintPackages

	// Build SSA code for the whole program.
	prog.Build()

	for _, p := range pkgs {
		if p != nil {
			p.Build()
		}
	}

	return prog, pkgs, suc
}

// Ssa_build_one_pkg_x depends on x/tools/go/packages
func Ssa_build_one_pkg_x(path string, show_compile_error bool) (result *ssa.Package, suc bool) {

	result = nil
	suc = false

	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Tests: true}

	start := time.Now()

	initial, err := packages.OriLoad(cfg, "github.com/docker/docker/integration-cli/environment")
	//initial, err := packages.OriLoad(cfg, path) // you can put multiple paths here, but it is unnecessary if you only want one program
	if err != nil {
		if show_compile_error {
			fmt.Println(err)
		}
		return
	}

	time_load := time.Since(start)
	start = time.Now()

	// Create SSA packages for all well-typed packages.
	//fmt.Println("build_mode:", "NaiveForm")                   //|BuildSerially|PrintPackage
	_, pkgs := ssautil.Packages(initial, ssa.NaiveForm) //|ssa.BuildSerially|ssa.PrintPackages

	time_create_ssa := time.Since(start)
	start = time.Now()

	if len(pkgs) != 1 {
		if show_compile_error {
			fmt.Println("Error when building", path, ":\n\tlength of slice \"pkgs\" is:", len(pkgs))
		}
		return
	}

	// Build SSA code for this package.
	result = pkgs[0]

	if result == nil {
		if show_compile_error {
			fmt.Println("Error when building", path, ":\n\tthe package is nil")
		}
		return
	}

	suc = true
	result.Build()

	time_build := time.Since(start)

	fmt.Println("Pkg:", path, "\tTime load:", time_load, "\tTime create ssa:", time_create_ssa, "\tTime build:", time_build)

	return
}

func Ssa_build_one_pkg_compiler(path string, show_compile_error bool) (*ssa.Package, error) {

	absolute_path := global.Absolute_root + path
	files, err := ioutil.ReadDir(absolute_path)
	if err != nil {
		if show_compile_error {
			fmt.Println(err)
		}
		return nil, err
	}

	type go_file struct {
		name     string
		abs_name string
		content  string
	}

	fileset_content := []go_file{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") == false {
			continue
		}
		absolute_name := absolute_path + "/" + file.Name()
		content, err := ioutil.ReadFile(absolute_name)
		if err != nil {
			if show_compile_error {
				fmt.Println(err)
			}
			return nil, err
		}
		new_file := go_file{
			name:     file.Name(),
			abs_name: absolute_name,
			content:  string(content),
		}
		fileset_content = append(fileset_content, new_file)
	}

	start := time.Now()

	// Parse the source files.
	token_fset := token.NewFileSet()
	ast_files := []*ast.File{}
	for _, a_go_file := range fileset_content {
		f, err := parser.ParseFile(token_fset, a_go_file.name, a_go_file.content, 0)
		if err != nil {
			if show_compile_error {
				fmt.Println(err)
			} // parse error
			return nil, err
		}
		ast_files = append(ast_files, f)
	}

	time_parse := time.Since(start)

	// Create the type-checker's package.
	var pkg_name string
	if last_index := strings.LastIndex(path, "/"); last_index > -1 {
		pkg_name = path[last_index+1:]
	} else {
		pkg_name = path
	}
	pkg := types.NewPackage(path, pkg_name)

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	conf := &types.Config{Importer: importer.Default()}
	result, _, err := ssautil.BuildPackage(
		conf, token_fset, pkg, ast_files, ssa.SanityCheckFunctions)
	if err != nil {
		if show_compile_error {
			fmt.Println(err)
		} // type error in some package
		return nil, err
	}

	time_check := time.Since(start)

	//fmt.Println("Pkg:",path,"Time parse:",time_parse,"Time check:",time_check)
	_ = time_parse

	time_check_ms := time_check / time.Millisecond
	time_check_remain := time_check % time.Millisecond
	time_check_ms_float := float64(time_check_ms) + float64(time_check_remain)/1e6
	fmt.Println(time_check_ms_float)
	return result, nil
}

// Get_and_install changes GOPATH, runs "go get XXX", and changes GOPATH back
func Get_and_install(path string) {
	defer func() {
		os.Setenv("GOPATH", global.GOPATH)
	}()

	err := os.Setenv("GOPATH", global.Target_GOPATH)
	if err != nil {
		return
	}

	cmd := exec.Command("bash", "-c", "go get "+path)
	cmd.Dir = global.Absolute_root
	err = cmd.Run()
	if err != nil {
		return
	}
}
