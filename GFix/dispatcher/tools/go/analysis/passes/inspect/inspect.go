// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package inspect defines an Analyzer that provides an AST inspector
// (github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ast/inspect.Inspect) for the syntax trees of a
// package. It is only a building block for other analyzers.
//
// Example of use in another analysis:
//
//	import (
//		"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis"
//		"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis/passes/inspect"
//		"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ast/inspector"
//	)
//
//	var Analyzer = &analysis.Analyzer{
//		...
//		Requires:       reflect.TypeOf(new(inspect.Analyzer)),
//	}
//
// 	func run(pass *analysis.Pass) (interface{}, error) {
// 		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
// 		inspect.Preorder(nil, func(n ast.Node) {
// 			...
// 		})
// 		return nil
// 	}
//
package inspect

import (
	"reflect"

	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/analysis"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:             "inspect",
	Doc:              "optimize AST traversal for later passes",
	Run:              run,
	RunDespiteErrors: true,
	ResultType:       reflect.TypeOf(new(inspector.Inspector)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	return inspector.New(pass.Files), nil
}
