package output

import (
	"fmt"
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/tools/go/callgraph"
)

func PrintCallGraph(graph * callgraph.Graph) {
	for fn, node := range graph.Nodes {

		if fn.Pkg == nil || fn.Pkg.Pkg == nil {
			continue
		}

		if fn.Pkg.Pkg.Path() != config.StrRelativePath {
			continue
		}

		fmt.Println(fn.String())

		for _, edge := range node.Out {
			fmt.Print("-> ")
			fmt.Print(edge.Callee.Func.Name())
			loc := (config.Prog.Fset).Position(edge.Site.Pos())
			if loc.Line > 0 {
				fmt.Print(" at ", loc.Filename, ": ", loc.Line)
			}

			fmt.Println()
		}

		fmt.Println()
		fmt.Println()
	}
}
