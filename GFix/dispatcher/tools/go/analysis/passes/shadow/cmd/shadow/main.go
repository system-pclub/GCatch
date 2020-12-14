// The shadow command runs the shadow analyzer.
package main

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/passes/shadow"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
