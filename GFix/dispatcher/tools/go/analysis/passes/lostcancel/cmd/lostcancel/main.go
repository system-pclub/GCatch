// The lostcancel command applies the github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/passes/lostcancel"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
