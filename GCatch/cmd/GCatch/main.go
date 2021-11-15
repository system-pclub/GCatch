package main

import (
	"flag"
	"fmt"
	"github.com/system-pclub/GCatch/GCatch/checkers/doublelock"
	"github.com/system-pclub/GCatch/GCatch/ssabuild"
	"github.com/system-pclub/GCatch/GCatch/tools/go/callgraph"
	"github.com/system-pclub/GCatch/GCatch/tools/go/mypointer"
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"github.com/system-pclub/GCatch/GCatch/util"
	"go/types"
	"os"
	"sort"
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

	for strCheckerName, _ := range mapCheckerName {
		switch strCheckerName {
		case "unlock": forgetunlock.Initialize()
		case "double": doublelock.Initialize()
		case "conflict", "structfield", "fatal", "BMOC": // no need to initialize these checkers
		default:
			fmt.Println("Warning, a not existing checker is in -checker= flag:", strCheckerName)
		}
	}

	var errMsg string
	var bSucc bool


	config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(config.StrEntrancePath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.

	if bSucc && len(config.Prog.AllPackages()) > 0 {
		// Step 2.1, Case 1: built SSA successfully, run the checkers in process()
		fmt.Println("Successfully built whole program. Now running checkers")

		detect(mapCheckerName)

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

	for index, wpath := range wPaths {

		//fmt.Println(wpath.StrPath)

		if wpath.NumLock + wpath.NumSend == 0 {
			break
		}

		if index < intSkipPkg || index >= intExitPkg {
			continue
		}


		config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(wpath.StrPath, false, boolShowCompileError) // Create SSA packages for the whole program including the dependencies.
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

				config.Prog, config.Pkgs, bSucc, errMsg = ssabuild.BuildWholeProgram(child.StrPath, true, boolShowCompileError) // Force the package to build, at least some dependencies of it are being built and checked
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

type Inter struct {
	memStr string
	typesInter *types.Interface
	typesStr string
	typesNamed *types.Named
	typesVecMethod []*types.Func
	boolAnyInter bool
}

type NormalType struct {
	memStr string
	memType *ssa.Type
	typesType types.Type
	vecMethodFn []*ssa.Function
	vecImplInter []*Inter

	// fields after this NormalType enters vecTargetType
	vecFnLength []int
	totalFnLength int
}

const UNKNOWNLENGTH int = 123456


func detect(mapCheckerName map[string]bool) {

	// This is a hack to find code that may be rewritten in generics:
	/*
	As in Fig. 2 of the FG paper https://arxiv.org/pdf/2005.11710.pdf, the Equal function of Int actually uses that as Int. However, because it has to implement Eq interface, the type of that must be Eq. Can you write a checker to look for similar cases in docker or k8s?
	Specifically, we can look for patterns like this:
	1) one function implements an interface
	and 2) one of the function's parameter is an interface,
	and 3) it is mostly/all asserted to a struct before every usage.
	*

	 */


	vecInter := []*Inter{}
	vecType  := []*NormalType{}

	for _, pkg := range config.Prog.AllPackages() {
		for memStr, mem := range pkg.Members {
			typeMem := mem.Type()
			if typeMem == nil {
				continue
			}

			// If the member is an interface, record in vecInter
			if typeInter, ok := typeMem.Underlying().(*types.Interface); ok {
				newInter := &Inter{
					memStr: 	memStr,
					typesInter: typeInter,
					typesStr:   typeMem.String(),
					typesNamed: nil,
					typesVecMethod: []*types.Func{},
				}
				newInter.typesNamed, _ = typeMem.(*types.Named)
				for i := 0; i < typeInter.NumMethods(); i++ {
					newInter.typesVecMethod = append(newInter.typesVecMethod, typeInter.Method(i))
				}
				if typeInter.NumMethods() == 0 {
					newInter.boolAnyInter = true
				} else {
					newInter.boolAnyInter = false
				}
				vecInter = append(vecInter, newInter)
			} else {
				if memType, ok := mem.(*ssa.Type);ok {
					// If the member is just a Type, record in vecType
					newNType := &NormalType{
						memStr:    mem.String(),
						memType:  memType,
						typesType: typeMem,
						vecMethodFn: []*ssa.Function{},
						vecImplInter: []*Inter{},
						vecFnLength: []int{},
					}

					methodSet := config.Prog.MethodSets.MethodSet(typeMem)
					for j := 0; j < methodSet.Len(); j++ {
						methodSelection := methodSet.At(j)
						progFnMethod := config.Prog.MethodValue(methodSelection)
						if progFnMethod != nil {
							newNType.vecMethodFn = append(newNType.vecMethodFn, progFnMethod)
						}

					}

					vecType = append(vecType, newNType)
				}

			}


		}
	}

	for _, t := range vecType {
		for _, inter := range vecInter {
			if types.Implements(t.typesType, inter.typesInter) {
				t.vecImplInter = append(t.vecImplInter, inter)
			}
		}
	}

	PrintVecInter(vecInter)


	mapTypesType2Inter := make(map[types.Type]*Inter)
	for _, inter := range vecInter {
		mapTypesType2Inter[inter.typesInter] = inter
	}
	mapTypesType2Type := make(map[types.Type]*NormalType)
	for _, t := range vecType {
		mapTypesType2Type[t.typesType] = t
	}

	// T1: the types that have at least one method who has at least one parameter that is an interface
	//		, ranked by lines of code
	// T2: the types that are simply in the program we are interested
	// 		, ranked by lines of code
	// T3: the types that are simply in the program we are interested
	// 		, ranked by number of interfaces that it implements
	vecTargetTypeT1 := []*NormalType{}
	vecTargetTypeT2 := []*NormalType{}
	vecTargetTypeT3 := []*NormalType{}

	loopType:
	for _, t := range vecType {

		// Has methods
		// Implements some interface
		// In target program
		if len(t.vecMethodFn) == 0 || len(t.vecImplInter) == 0{
			continue
		}
		typesStr := t.typesType.String()
		if ! strings.Contains(typesStr, config.StrRelativePath) {
			continue
		}

		vecTargetTypeT2 = append(vecTargetTypeT2, t)
		vecTargetTypeT3 = append(vecTargetTypeT3, t)

		for _, method := range t.vecMethodFn {
			strMethod := method.String()
			_ = strMethod
			beginPosition := config.Prog.Fset.Position(method.Pos())
			fileName := beginPosition.Filename
			line := beginPosition.Line
			_ = line
			_ = fileName
			for i, para := range method.Params {
				if i == 0 { // Surely the receiver meets our standard, let's skip it
					continue
				}
				if _, ok := mapTypesType2Type[para.Type()];ok {
					vecTargetTypeT1 = append(vecTargetTypeT1, t)
					continue loopType
				}
			}
		}
	}

	for _, t := range vecTargetTypeT2 { // T2 is enough, it covers all the types in T1 T2 and T3
		for i, method := range t.vecMethodFn {

			t.vecFnLength = append(t.vecFnLength, UNKNOWNLENGTH)
			_ = method
			pos := method.Pos()
			//Pos() returns the declaring ast.FuncLit.Type.Func
			//or the position of the ast.FuncDecl.Name, if the function was explicit in the source.
			_ = pos
			ast := method.Syntax()
			// Syntax returns an ast.Node whose Pos/End methods provide the lexical extent of the function
			//if it was defined by Go source code (f.Synthetic==""), or nil otherwise.
			//If f was built with debug information (see Package.SetDebugRef),
			//the result is the *ast.FuncDecl or *ast.FuncLit that declared the function.
			//Otherwise, it is an opaque Node providing only position information; this avoids pinning the AST in memory.
			if ast == nil {
				continue
			}
			beginPos := ast.Pos()
			beginPosition := config.Prog.Fset.Position(beginPos)
			endPos := ast.End()
			endPosition := config.Prog.Fset.Position(endPos)
			///DEBUG:
			//fmt.Printf("FnStr:%s\nBegin:%s:%d\nEnd:%s:%d\n", method.String(), beginPosition.Filename,
			//	beginPosition.Line, endPosition.Filename, endPosition.Line)
			if beginPosition.Filename != "" && beginPosition.Filename == endPosition.Filename &&
				beginPosition.Line != 0 && endPosition.Line != 0 && endPosition.Line >= beginPosition.Line {
				length := endPosition.Line - beginPosition.Line + 1
				t.vecFnLength[i] = length
				t.totalFnLength += length
			} else {
				t.vecFnLength[i] = UNKNOWNLENGTH
			}

		}
	}

	sort.SliceStable(vecTargetTypeT1, func(i, j int) bool {
		return vecTargetTypeT1[i].totalFnLength > vecTargetTypeT1[j].totalFnLength
	})

	sort.SliceStable(vecTargetTypeT2, func(i, j int) bool {
		return vecTargetTypeT2[i].totalFnLength > vecTargetTypeT2[j].totalFnLength
	})

	sort.SliceStable(vecTargetTypeT3, func(i, j int) bool {
		return len(vecTargetTypeT3[i].vecImplInter) > len(vecTargetTypeT3[j].vecImplInter)
	})


	PrintVecType(vecTargetTypeT1, "printInterestingTypeT1_")
	PrintVecType(vecTargetTypeT2, "printInterestingTypeT2_")
	PrintVecType(vecTargetTypeT3, "printInterestingTypeT3_")

	return

}

func PrintVecType(vecType []*NormalType, strFilePre string) {

	f, err := os.Create("/home/ziheng/Go/gcatch/src/github.com/system-pclub/GCatch/GCatch/results/" + strFilePre + strings.ReplaceAll(config.StrRelativePath, "/", "_") + ".txt")
	if err != nil {
		fmt.Println("Filed to create file", err)
		return
	}
	defer f.Close()

	str := ""

	for i, t := range vecType {
		if len(t.vecImplInter) == 0 || len(t.vecMethodFn) == 0 {
			continue
		}
		str += fmt.Sprintf("-----NO.%d\nMemStr:%s\ntypesStr:%s\n", i, t.memStr, t.typesType.String())

		str += fmt.Sprintf("Implements interfaces:\n")
		for j, inter := range t.vecImplInter {
			str += fmt.Sprintf("\tNO.%d_%d:\n", i, j)
			str += "\t\t" + inter.typesStr + "\n"
		}

		methodSet := config.Prog.MethodSets.MethodSet(t.typesType)
		str += fmt.Sprintf("Method set:\n")
		for j := 0; j < methodSet.Len(); j++ {
			strMethod := fmt.Sprintf("\tNO.%d_%d:\n", i, j)
			methodSelection := methodSet.At(j)
			switch methodSelection.Kind() {
			case types.MethodVal:
				strMethod += "\t\tKind:MethodVal"
			case types.MethodExpr:
				strMethod += "\t\tKind:MethodExpr"
			case types.FieldVal:
				strMethod += "\t\tKind:FieldVal"
			}
			strMethod += "\n\t\tTypesStr:" + methodSelection.String()
			progFnMethod := config.Prog.MethodValue(methodSelection)
			if progFnMethod == nil {
				strMethod += "\n\t\tSSAFnStr:Can't find"
			} else {
				strMethod += "\n\t\tSSAFnStr:" + progFnMethod.String()
			}

			str += strMethod + "\n"
		}

	}

	_, err = f.WriteString(str)
	if err!= nil {
		fmt.Println("Failed to print str in file")
		return
	}
}

func PrintVecInter(vecInter []*Inter) {
	f, err := os.Create("/home/ziheng/Go/gcatch/src/github.com/system-pclub/GCatch/GCatch/results/printInter_" + strings.ReplaceAll(config.StrRelativePath, "/", "_") + ".txt")
	if err != nil {
		fmt.Println("Filed to create file", err)
		return
	}
	defer f.Close()

	str := ""



	for i, inter := range vecInter {
		str += fmt.Sprintf("-----NO.%d\nMemStr:%s\ntypesStr:%s\n", i, inter.memStr, inter.typesStr)
		str += fmt.Sprintf("Number of methods:%d\nIs empty:%b\n", len(inter.typesVecMethod), inter.boolAnyInter)

		//Don't use this function, it will just return nil files: methodSet := config.Prog.MethodSets.MethodSet(inter.typesInter)
	}

	_, err = f.WriteString(str)
	if err!= nil {
		fmt.Println("Failed to print str in file")
		return
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
