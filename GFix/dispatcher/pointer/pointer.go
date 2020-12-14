package pointer

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/mypointer"
)

// Pointer_build_callgraph runs a hacked version of Go's pointer analysis. This function will
// give no queries to pointer analysis, just getting the callgraph
func Pointer_build_callgraph() *mypointer.Result {
	cfg := &mypointer.Config{
		OLDMains:        nil,
		Prog:            global.Prog,
		Reflection:      global.Pointer_consider_reflection,
		BuildCallGraph:  true,
		Queries:         nil,
		IndirectQueries: nil,
		Log:             nil,
	}
	result, err := mypointer.Analyze(cfg, nil)
	if err != nil {
		fmt.Println("Error when building callgraph with nil Queries:\n", err.Error())
		return nil
	}

	return result
}




