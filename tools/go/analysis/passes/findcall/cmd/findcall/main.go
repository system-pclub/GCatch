// The findcall command runs the findcall analyzer.
package main

import (
	"github.com/system-pclub/GCatch/tools/go/analysis/passes/findcall"
	"github.com/system-pclub/GCatch/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(findcall.Analyzer) }
