package main

import (
	"flag"
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/checkers/bmoc"
	"github.com/system-pclub/GCatch/GCatch/checkers/conflictinglock"
	"github.com/system-pclub/GCatch/GCatch/checkers/doublelock"
	"github.com/system-pclub/GCatch/GCatch/checkers/fatal"
	"github.com/system-pclub/GCatch/GCatch/checkers/structfield"
	"github.com/system-pclub/GCatch/GCatch/ssabuild"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GCatch/util"
	"github.com/system-pclub/GCatch/GCatch/util/genKill"
	"os"
	"strings"
	"time"

	"github.com/system-pclub/GCatch/GCatch/checkers/forgetunlock"
	"github.com/system-pclub/GCatch/GCatch/config"
)


func main() {

	mainStart := time.Now()
	defer func() {
		mainDur := time.Since(mainStart)
		fmt.Println("\n\nTime of main(): seconds", mainDur.Seconds())
	}()

	pProjectPath := flag.String("path","","Full path of the target project")
	pRelativePath := flag.String("include","","Relative path (what's after /src/) of the target project")
	pCheckerName := flag.String("checker", "BMOC", "the checker to be used, divided by \":\"")
	pShowCompileError := flag.Bool("compile-error", false, "If fail to compile a package, show the errors of compilation")
	pExcludePath := flag.String("exclude", "vendor", "Name of directories that you want to ignore, divided by \":\"")
	pRobustMod := flag.Bool("r", false, "If the main package can't pass compiler, check subdirectories one by one")
	pFnPointerAlias := flag.Bool("pointer", true, "Whether alias analysis is used to figure out function pointers")
	pSkipPkg := flag.Int("skip", -1, "Skip the first N packages")
	pExitPkg := flag.Int("exit", 99999, "Exit when meet the Nth packages")
	pPrintMod := flag.String( "print-mod", "", "Print information like the number of channels, divided by \":\"")
	pGoMod := flag.Bool("mod", false, "Beta functionality: Use this flag to indicate GCatch to build a program (and only it) by go.mod")
	pGoModulePath := flag.String("mod-module-path", "", "Beta functionality: The module path of the program you want to check, like google.golang.org/grpc")
	pGoModAbsPath := flag.String("mod-abs-path", "", "Beta functionality: The absolute path of the program you want to check, which contains go.mod")

	flag.Parse()

	strProjectPath := *pProjectPath
	strRelativePath := *pRelativePath
	mapCheckerName := util.SplitStr2Map(*pCheckerName, ":")
	boolShowCompileError := *pShowCompileError
	boolRobustMod := *pRobustMod
	boolFnPointerAlias := *pFnPointerAlias
	intSkipPkg := *pSkipPkg
	intExitPkg := *pExitPkg

	go func(){
		time.Sleep(time.Duration(config.MAX_GCATCH_DDL_SECOND) * time.Second)
		fmt.Println("!!!!")
		fmt.Println("The checker has been running for", config.MAX_GCATCH_DDL_SECOND,"seconds. Now force exit")
		os.Exit(1)
	}()

	for strCheckerName, _ := range mapCheckerName {
		switch strCheckerName {
		case "unlock": forgetunlock.Initialize()
		case "double": doublelock.Initialize()
		case "conflict", "structfield", "fatal", "BMOC": // no need to initialize these checkers
		case "chSafety":
			config.BoolChSafety = true
		default:
			fmt.Println("Warning, a not existing checker is in -checker= flag:", strCheckerName)
		}
	}

	var errMsg string
	var bSucc bool

	config.BoolGoMod = *pGoMod

	if config.BoolGoMod {
		// A beta functionality: using go.mod to build a program
		config.StrModAbsPath = *pGoModAbsPath
		config.StrModulePath = *pGoModulePath

		if config.StrModAbsPath == "" || config.StrModulePath == "" {
			fmt.Println("mod-module-path or mod-abs-path is empty\nPlease use -help to see what they are and provide valid values")
			return
		}

		config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgramGoMod(config.StrModulePath, false, boolShowCompileError, config.StrModAbsPath) // Create SSA packages for the whole program including the dependencies.

		// What about all the other global variables we defined in traditional way and may use later? Let's go through them one by one
		// (0) config.StrEntrancePath: not used later
		// (1) config.StrGOPATH: not used later
		// (2) config.MapExcludePaths: used in IsPathIncluded in config/path.go, should be fine to just use the same code
		//								used in ListAllPkgPaths in config/path.go, let's avoid using it by disabling the -r flag later
		config.MapExcludePaths = util.SplitStr2Map(*pExcludePath, ":")
		// (3) config.StrRelativePath: used in PrintCallGraph in output/callgraph.go; it is not called anyway
		//								used in IsPathIncluded in config/path.go, need to edit this function to use StrModulePath
		//								used in ListAllPkgPaths in config/path.go like (2), let's disable the -r flag later
		// (4) config.StrAbsolutePath: used in multiple functions in config/path.go, but all can be avoided by disabling the -r flag
		//
		// (5) config.BoolDisableFnPointer: should be fine to just use the legacy code
		config.BoolDisableFnPointer = ! boolFnPointerAlias
		// (6) config.MapPrintMod:			should be fine to just use the legacy code
		config.MapPrintMod = util.SplitStr2Map(*pPrintMod, ":")
		// (7) config.MapHashOfCheckedCh:			should be fine to just use the legacy code
		config.MapHashOfCheckedCh = make(map[string]struct{})
	} else {
		// Our recommended way: building the program by GOPATH
		numIndex := strings.LastIndex(strProjectPath, "/src/")
		if numIndex < 0 {
			fmt.Println("The target project is not in a GOPATH, because its path doesn't contain \"/src/\"")
			os.Exit(2)
		}

		config.StrEntrancePath = strProjectPath[numIndex+5:]
		config.StrGOPATH = os.Getenv("GOPATH")
		config.MapExcludePaths = util.SplitStr2Map(*pExcludePath, ":")
		config.StrRelativePath = strRelativePath
		config.StrAbsolutePath = strProjectPath[:numIndex+5]
		config.StrAbsolutePath = strings.ReplaceAll(config.StrAbsolutePath, "//", "/")
		config.BoolDisableFnPointer = ! boolFnPointerAlias
		config.MapPrintMod = util.SplitStr2Map(*pPrintMod, ":")
		config.MapHashOfCheckedCh = make(map[string]struct{})

		/*
			fmt.Println("entrance", config.StrEntrancePath)
			fmt.Println("gopath", config.StrGOPATH)
			fmt.Println("relative", config.StrRelativePath)
			fmt.Println("absolute", config.StrAbsolutePath)
		*/


		if strings.Contains(config.StrGOPATH, strProjectPath[:numIndex]) == false {
			fmt.Println("The input path doesn't match GOPATH. GOPATH of target project:", strProjectPath[:numIndex], "\tGOPATH:", os.Getenv("GOPATH"))
			os.Exit(3)
		}

		config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgramTrad(config.StrEntrancePath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.
	}

	if bSucc && len(config.Prog.AllPackages()) > 0 {
		// Step 2.1, Case 1: built SSA successfully, run the checkers in process()
		fmt.Println("Successfully built whole program. Now running checkers")

		detect(mapCheckerName)

	} else {
		// Step 2.1, Case 2: building SSA failed
		fmt.Println("Failed to build the whole program. The entrance package or its dependencies have error.", errMsg)
	}

	// Step 2.2 If -r is used, continue checking all child packages
	// however, if -mod is used, then ignore -r, and just stop here because we don't know how to walk other files in ListWorthyPaths
	if config.BoolGoMod {
		fmt.Println("Exit. You are using the -mod flag, which only supports building and checking the specified program (and its dependencies)\n")
		return
	}

	if ! boolRobustMod {
		fmt.Println("Exit. If you want to scan subdirectories and use -race, please use -r")
		return
	}

	fmt.Println("Now trying to build unchecked packages separately...")


	// Step 2.3: List paths of packages that contain "Lock" or "<-" in source code, and rank the paths with the number of "Lock" or "<-"
	wPaths := config.ListWorthyPaths()

	for index, wpath := range wPaths {

		//fmt.Println(wpath.StrPath)

		if wpath.NumLock + wpath.NumSend == 0 {
			break
		}

		if index < intSkipPkg || index >= intExitPkg {
			continue
		}


		config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgramTrad(wpath.StrPath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.
		if bSucc {
			fmt.Println("Successful. Package NO.", index, ":", wpath.StrPath, " Num of Lock & <-:", wpath.NumLock + wpath.NumSend)
			detect(mapCheckerName)
			mainDur := time.Since(mainStart)
			fmt.Println("\n\nTime passed for seconds", mainDur.Seconds())
		} else {
			// Step 2.4, Case 2 : building SSA failed; build its children packages
			fmt.Println("Fail. Package NO.", index, ":", wpath.StrPath, " Num of Lock & <-:", wpath.NumLock + wpath.NumSend, " error:", errMsg)
			for j, child := range wpath.VecChildrenPath {

				if child.NumLock + child.NumSend == 0 {
					break
				}

				config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgramTrad(child.StrPath, true, boolShowCompileError) // Force the package to build, at least some dependencies of it are being built and checked
				if bSucc {
					fmt.Println("\tSuccessfully built sub-Package NO.",j,":\t",child.StrPath, " Num of Lock & <-:", child.NumLock + child.NumSend)
					detect(mapCheckerName)
				} else if errMsg == "load_err" {
					fmt.Println("\tFailed to build sub-Package NO.",j,":\t",child.StrPath, " Num of Lock & <-:", child.NumLock + child.NumSend)
				} else if errMsg == "type_err" {
					fmt.Println("\tPartially built sub-Package NO.",j,":\t",child.StrPath, " Num of Lock & <-:", child.NumLock + child.NumSend)
					detect(mapCheckerName)

				}
			}
		}
	}

}

func detect(mapCheckerName map[string]bool) {

	config.Inst2Defers, config.Defer2Insts = genKill.ComputeDeferMap()

	boolNeedCallGraph := mapCheckerName["double"] || mapCheckerName["conflict"] || mapCheckerName["structfield"]  ||
		mapCheckerName["fatal"]  || mapCheckerName["BMOC"]
	if boolNeedCallGraph {
		config.CallGraph = BuildCallGraph()
		if config.CallGraph == nil {
			return
		}
	}

	for strCheckerName, _ := range mapCheckerName {
		switch strCheckerName {
		case "unlock":
			forgetunlock.Detect()
		case "double":
			doublelock.Detect()
		case "conflict":
			conflictinglock.Detect()
		case "structfield":
			structfield.Detect()
		case "fatal":
			fatal.Detect()
		case "BMOC":
			bmoc.Detect()
		}
	}
}


func BuildCallGraph() * callgraph.Graph {
	cfg := & mypointer.Config{
		OLDMains:        nil,
		Prog:            config.Prog,
		Reflection:      config.POINTER_CONSIDER_REFLECTION,
		BuildCallGraph:  true,
		Queries:         nil,
		IndirectQueries: nil,
		Log:             nil,
	}
	result, err := mypointer.Analyze(cfg, nil)
	defer func() {
		cfg = nil
		result = nil
	}()
	if err != nil {
		fmt.Println("Error when building callgraph with nil Queries:\n", err.Error())
		return nil
	}
	graph := result.CallGraph
	return graph
}
