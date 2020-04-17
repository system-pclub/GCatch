// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The vet command is a static checker for Go programs. It has pluggable
// analyzers defined using the github.com/system-pclub/gochecker/tools/go/analysis API, and
// using the github.com/system-pclub/gochecker/tools/go/packages API to load packages in any
// build system.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"github.com/system-pclub/gochecker/tools/go/analysis/multichecker"

	// analysis plug-ins
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/asmdecl"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/assign"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/atomic"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/atomicalign"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/bools"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/buildtag"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/cgocall"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/composite"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/copylock"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/errorsas"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/httpresponse"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/loopclosure"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/lostcancel"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/nilfunc"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/printf"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/shift"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/stdmethods"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/structtag"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/tests"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/unmarshal"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/unreachable"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/unsafeptr"
	"github.com/system-pclub/gochecker/tools/go/analysis/passes/unusedresult"
)

func main() {
	// This suite of analyzers is applied to all code
	// in GOROOT by GOROOT/src/cmd/vet/all. When adding
	// a new analyzer, update the whitelist used by vet/all,
	// or change its vet command to disable the new analyzer.
	multichecker.Main(
		// the traditional vet suite:
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		atomicalign.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		errorsas.Analyzer,
		httpresponse.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		stdmethods.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,

		// for debugging:
		// findcall.Analyzer,
		// pkgfact.Analyzer,

		// uses SSA:
		// nilness.Analyzer,
	)
}
