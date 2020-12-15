// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// Uninterpreted is a symbolic value with no interpretation.
// Uninterpreted values have identity—that is, they can be compared
// for equality—but have no inherent meaning otherwise.
//
// Uninterpreted and FiniteDomain are similar, except that
// Uninterpreted values come from an unbounded universe.
//
// Uninterpreted implements Value.
type Uninterpreted value

func init() {
	kindWrappers[KindUninterpreted] = func(x value) Value {
		return Uninterpreted(x)
	}
}

// UninterpretedSort returns a sort for uninterpreted values.
//
// Two uninterpreted sorts are the same if and only if they have the
// same name.
func (ctx *Context) UninterpretedSort(name string) Sort {
	sym := ctx.symbol(name)
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_uninterpreted_sort(ctx.c, sym), KindUninterpreted)
	})
	return sort
}

//go:generate go run genwrap.go -t Uninterpreted $GOFILE
