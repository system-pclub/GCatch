// The unmarshal command runs the unmarshal analyzer.
package main

import (
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/unmarshal"
	"github.com/system-pclub/gochecker/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unmarshal.Analyzer) }
