package callgraph

import (
	"fmt"
	"github.com/system-pclub/GCatch/tools/go/ssa"
	"strconv"
	"strings"
)

// PruneSoundly is aimed to make the callgraph more precise, without hurting its soundness
// Currently we have 2 strategies to prune the callgraph: 1. When an anonymous function does
// not escape, delete all In edges of it except the real one. 2. ...
func (g *Graph) PruneSoundly() {
	g.pruneAnony()
	//g.pruneFuncPointer()

	g.validate()
}

// PruneSDK removes all In and Out Edges of a Node if its package is in Go SDK or golang.org. The Node is not deleted
func (g *Graph) PruneSDK() {
	for _, node := range g.Nodes {
		if node.Func == nil {
			continue
		}
		if node.Func.Pkg == nil {
			continue
		}
		flagIsSDK := false
		path := node.Func.Pkg.Pkg.Path()
		if strings.HasPrefix(path, "golang.org") || strings.Contains(path, "/golang.org") {
			flagIsSDK = true
		} else {
			for _, SDKPrefix := range SDKList {
				if strings.HasPrefix(path,SDKPrefix+"/") || path == SDKPrefix {
					flagIsSDK = true
					break
				}
			}
		}
		if flagIsSDK {
			node.deleteIns()
			node.deleteOuts()
		}
	}
}

var SDKList = []string  {"encoding", "io", "os", "sort", "errors", "path", "strconv", "expvar", "log", "plugin", "strings",
	"cmd", "flag", "sync", "fmt", "syscall", "archive", "compress", "go", "reflect", "testdata", "container", "hash",
	"math", "regexp", "testing", "bufio", "context", "html", "mime", "text", "crypto", "image", "time", "builtin",
	"database", "index", "unicode", "bytes", "debug", "internal", "net", "runtime", "unsafe"}

// PruneInvoke can make the callgraph more precise regarding interface method invoking, but it makes the callgraph
// no longer sound. For a CallInstruction that is of "invoke" mode, try to only reserve Out edges whose callee is in
// the same package as the caller. If failed to find any, then reserve that in the same package as the one
// defines the interface. If still failed, do nothing.
func (g *Graph) PruneInvoke(prog *ssa.Program) {
	for _, node := range g.Nodes {
		CalleeMap := make(map[ssa.CallInstruction][]*Edge) // Put all Edge share the same Site under the same key
		for _,OutEdge := range node.Out {
			if OutEdge.Site.Common().IsInvoke() == false {
				continue
			}
			CalleeMap[OutEdge.Site] = append(CalleeMap[OutEdge.Site],OutEdge)
		}

		const NotFound string = "NOTFOUND"
		for callInstr, edges := range CalleeMap {

			callerPkg := NotFound
			if callInstr.Parent().Pkg != nil {
				callerPkg = callInstr.Parent().Pkg.Pkg.Path()
			}

			interfacePkg := NotFound
			if callInstr.Common().Method.Pkg() != nil {
				interfacePkg =  callInstr.Common().Method.Pkg().Path()
			}

			var edgesAdopted []*Edge
			// Class1: calleePkg is the same as callerPkg; Class2: calleePkg is the same as interfacePkg
			edgesSameWithCaller := []*Edge{}
			edgesSameWithInterface := []*Edge{}

			for _, edge := range edges {
				calleePkg := NotFound

				if edge.Callee.Func.Pkg != nil { // nil when edge.Callee.Func is a wrapper or error.Error, which we don't care
					calleePkg = edge.Callee.Func.Pkg.Pkg.Path()
				} else {
					fn := edge.Callee.Func
					fn = edge.Callee.Func
					_ = fn
				}

				if calleePkg == callerPkg && calleePkg != NotFound {
					edgesSameWithCaller = append(edgesSameWithCaller, edge)
				}
				if calleePkg == interfacePkg && calleePkg != NotFound {
					edgesSameWithInterface = append(edgesSameWithInterface, edge)
				}
			}

			if len(edgesSameWithCaller) != 0 {
				edgesAdopted = edgesSameWithCaller
			} else if len(edgesSameWithInterface) != 0 {
				edgesAdopted = edgesSameWithInterface
			} else {
				continue
			}

			if callerPkg != NotFound {
				if strings.Contains(callerPkg,"k8s.io") && !strings.Contains(callerPkg,"golang.org") && !strings.Contains(callerPkg,"vendor") {
					fmt.Println("-----------")
					fmt.Print("Call Site:")
					fmt.Println(callInstr)
					target_position := (prog.Fset).Position(callInstr.Pos())
					fmt.Printf("\tFile: %s:%d\n",target_position.Filename,target_position.Line)
					fmt.Println("Origin callee:")
					for _, edge := range edges {
						fmt.Println("\t",edge.Callee.Func.String())
					}
					fmt.Println("New callee:")
					for _, edge := range edgesAdopted {
						fmt.Println("\t",edge.Callee.Func.String())
					}
				}
			}

			for _, edge := range edges {
				flagReserve := false
				for _, edgeAdopted := range edgesAdopted {
					if edgeAdopted == edge {
						flagReserve = true
						break
					}
				}
				if flagReserve == false {
					removeOutEdge(edge)
					removeInEdge(edge)
				}
			}
		}
	}
}

// validate will print an error message if the callgraph is not valid
func (g *Graph) validate() {
	for fn, node := range g.Nodes {
		if fn == nil {
			if node.Func != nil {
				fmt.Println("Nil Fn but Node.Func is", node.Func.String())
			}
			continue
		}

		if fn != node.Func {
			fmt.Println("Fn != Node.Func")
		}

		for _,edge := range node.In {
			if edge.Caller == nil {
				fmt.Println("Edge.Caller == nil: Edge:", edge.Description(),"\tNode:",node.String())
			}
			if edge.Callee != node {
				fmt.Println("Edge.Callee != node; Edge:", edge.Description(), "\tNode:",node.String())
			}
		}

		for _,edge := range node.Out {
			if edge.Callee == nil {
				fmt.Println("Edge.Callee == nil: Edge:", edge.Description(),"\tNode:",node.String())
			}
			if edge.Caller != node {
				fmt.Println("Edge.Caller != node; Edge:", edge.Description(), "\tNode:",node.String())
			}
		}
	}
}

// pruneAnony looks for anonymous functions that do not escape, and
// delete In edges of them except the real one.  A very simple escape analysis is applied.
func (g *Graph) pruneAnony() {
	for fn,node := range g.Nodes {
		if fn == nil {
			continue
		}

		// continue if fn is not an anonymous function. An anonymous function ends with "$N", where N is an integer
		IntStrIndex := strings.LastIndex(fn.Name(),"$")
		if IntStrIndex == -1 {
			continue
		}
		IntStr := fn.Name()[IntStrIndex + 1:]
		if _,err := strconv.Atoi(IntStr); err != nil {
			continue
		}

		var useInstr ssa.Instruction // An anonymous function can only be used by one instruction
		aliasValueList := []ssa.Value{} // a list to store values of alias of the anonymous function
		fnParent := fn.Parent()
		fnStr := fn.String()
		outerLoop:
		for _,bb := range fnParent.Blocks {
			for _,instr := range bb.Instrs {
				for _,operand := range instr.Operands(nil) {
					if operandFn,ok := (*operand).(*ssa.Function); ok {
						if operandFn.String() == fnStr {
							useInstr = instr
							aliasValueList = append(aliasValueList,*operand)
							break outerLoop
						}
					}
				}
			}
		}

		flagEscaped := false // In current implementation, escaped means being stored or passed as an argument/binding
		var realCallSite ssa.CallInstruction // The one and only call site that can call fn. May be nil
		worklist := []ssa.Instruction{useInstr} // An instr is in worklist when it has an operand as an alias of fn

		isAlias := func(fnValue ssa.Value, aliasValueList []ssa.Value) bool {
			for _,aliasValue := range aliasValueList {
				if fnValue == aliasValue {
					return true
				}
			}
			return false
		}

		worklist_loop:
		for len(worklist) > 0 {
			next := worklist[0]

			switch concrete := next.(type) {

			// If an anonymous function is an operand of Defer/Go/Call, it only escapes when it is not CallCommon.Value
			case *ssa.Defer:
				if isAlias(concrete.Call.Value, aliasValueList) == false {
					flagEscaped = true
					break worklist_loop
				}
				realCallSite = concrete
			case *ssa.Go:
				if isAlias(concrete.Call.Value, aliasValueList) == false {
					flagEscaped = true
					break worklist_loop
				}
				realCallSite = concrete
			case *ssa.Call:
				if isAlias(concrete.Call.Value, aliasValueList) == false {
					flagEscaped = true
					break worklist_loop
				}
				realCallSite = concrete

			case *ssa.MakeClosure:
				// If an anonymous function is an operand of MakeClosure, it only escapes when it is not MakeClosure.Fn
				if isAlias(concrete.Fn, aliasValueList) == false {
					flagEscaped = true
					break worklist_loop
				}
				MakeClosureValue := ssa.Value(concrete)
				aliasValueList = append(aliasValueList,MakeClosureValue)
				for _,referrer := range *MakeClosureValue.Referrers() {
					worklist = append(worklist,referrer)
				}

			default:
				// Instructions can't be handled. Consider the function is escaped. E.g., stored in a variable, returned
				flagEscaped = true
				break worklist_loop
			}

			// delete next from worklist
			length := len(worklist)
			for i, elem := range worklist {
				if elem == next {
					worklist[i] = worklist[length-1]
					worklist[length - 1] = nil
					worklist = worklist[:length-1]
					break
				}
			}
		}

		if flagEscaped { // continue if this anonymous function escapes
			continue
		}

		// For a not escaped anonymous function, delete every In edge that is not realCallSite
		var realEdge *Edge
		for _,inEdge := range node.In {
			if inEdge.Site == realCallSite {
				realEdge = inEdge
				continue
			}
			defer func() {
				recover() // removeOutEdge may panic
			}()
			removeOutEdge(inEdge)
			inEdge = nil
		}

		if realEdge != nil {
			node.In = []*Edge {realEdge}
		}
	}
}
