package main

import (
	"flag"
	"fmt"
	"github.com/system-pclub/gochecker/ssabuild"
	"os"
	"strings"
	"time"

	"github.com/system-pclub/gochecker/checkers/forgetunlock"
	"github.com/system-pclub/gochecker/config"
)


func main() {

	pProjectPath := flag.String("path","","Full path of the target project")
	pRelativePath := flag.String("include","","Relative path (what's after /src/) of the target project")
	pCheckerName := flag.String("checker", "unlock", "the checker to be used")
	pShowCompileError := flag.Bool("compile-error", false, "If fail to compile a package, show the errors of compilation")
	pExcludePath := flag.String("exclude", "vendor", "Name of directories that you want to ignore, divided by \":\"")
	pRobustMod := flag.Bool("r", false, "If the main package can't pass compiler, check subdirectories one by one")

	flag.Parse()

	strProjectPath := *pProjectPath
	strRelativePath := *pRelativePath
	strCheckerName := *pCheckerName
	strExcludePath := *pExcludePath
	boolShowCompileError := *pShowCompileError
	boolRobustMod := *pRobustMod

	go func(){
		time.Sleep(time.Duration(config.MAX_SECOND) * time.Second)
		fmt.Println("The checker has been running for", config.MAX_SECOND,"seconds. Now force exit")
		os.Exit(1)
	}()


	numIndex := strings.Index(strProjectPath, "/src/")
	if numIndex < 0 {
		fmt.Println("The target project is not in a GOPATH, because its path doesn't contain \"/src/\"")
		os.Exit(2)
	}

	config.StrEntrancePath = strProjectPath[numIndex+5:]
	config.StrGOPATH = os.Getenv("GOPATH")
	config.VecExcludePaths = config.GetExcludePaths(strExcludePath)
	config.StrRelativePath = strRelativePath
	config.StrAbsolutePath = strProjectPath[:numIndex+5]
	config.StrAbsolutePath = strings.ReplaceAll(config.StrAbsolutePath, "//", "/")


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

	if strCheckerName == "unlock" {
		forgetunlock.Initialize()
	}

	var errMsg string
	var bSucc bool


	config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(config.StrEntrancePath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.

	if bSucc && len(config.Prog.AllPackages()) > 0 {
		// Step 2.1, Case 1: built SSA successfully, run the checkers in process()
		fmt.Println("Successfully built whole program. Now running checkers")

		detect(strCheckerName)

	} else {
		// Step 2.1, Case 2: building SSA failed
		fmt.Println("Failed to build the whole program. The entrance package or its dependencies have error.", errMsg)
	}


	// Step 2.2 If -r is used, continue checking all child packages
	if ! boolRobustMod {
		fmt.Println("Exit. If you want to scan subdirectories and use -race, please use -r")
		return
	}

	fmt.Println("Now trying to build unchecked packages separately...")


	// Step 2.3: List paths of packages that contain "Lock" or "<-" in source code, and rank the paths with the number of "Lock" or "<-"
	wPaths := config.ListWorthyPaths()



	//for _, path := range wPaths {
	//	fmt.Println( path.StrPath, path.NumLock, path.NumSend)
	//}

	//vecTestPackage := [] string {"github.com/etcd-io/etcd/mvcc", "github.com/etcd-io/etcd/raft"}

	for index, wpath := range wPaths {


		/*
	for index, strPath := range vecTestPackage {

		var wpath config.PkgPath

		for _, p := range wPaths {
			if strPath == p.StrPath {
				wpath = p
			}
		}



		if index > 6 {
			break
		}

		fmt.Println()
		fmt.Println()
		fmt.Println(wpath.StrPath) */

		// Step 2.4, Case 1 : built SSA successfully; run the checkers in process()
		config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(wpath.StrPath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.
		if bSucc {
			fmt.Println("Successful. Package NO.", index, ":", wpath.StrPath, " Num of Lock & <-:", wpath.NumLock + wpath.NumSend)
			detect(strCheckerName)
		} else {
			// Step 2.4, Case 2 : building SSA failed; build its children packages
			fmt.Println("Fail. Package NO.", index, ":", wpath.StrPath, " Num of Lock & <-:", wpath.NumLock + wpath.NumSend, " error:", errMsg)
			for j, child := range wpath.VecChildrenPath {
				config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(child, true, boolShowCompileError) // Force the package to build, at least some dependencies of it are being built and checked
				if bSucc {
					fmt.Println("\tSuccessfully built sub-Package NO.",j,":\t",child)
					detect(strCheckerName)
				} else if errMsg == "load_err" {
					fmt.Println("\tFailed to build sub-Package NO.",j,":\t",child)
				} else if errMsg == "type_err" {
					fmt.Println("\tPartially built sub-Package NO.",j,":\t",child)
					detect(strCheckerName)

				}
			}
		}


	}

}

func detect(strCheckerName string) {
	if strCheckerName == "unlock" {
		forgetunlock.Detect()
	}
}