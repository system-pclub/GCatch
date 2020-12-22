// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"
import "runtime"

// Array is a symbolic value representing an extensional array.
//
// Unlike typical arrays in programming, an extensional array has an
// arbitrary index (domain) sort, in addition to an arbitrary value
// (range) sort. It can also be viewed like a hash table, except that
// all possible keys are always present.
//
// Arrays are "updated" by storing a new value to a particular index.
// This creates a new array that's identical to the old array except
// that that index.
//
// Array implements Value.
type Array value

func init() {
	kindWrappers[KindArray] = func(x value) Value {
		return Array(x)
	}
}

// ArraySort returns a sort for arrays that are indexed by domain and
// have values from range.
func (ctx *Context) ArraySort(domain, range_ Sort) Sort {
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_array_sort(ctx.c, domain.c, range_.c), KindArray)
	})
	runtime.KeepAlive(domain)
	runtime.KeepAlive(range_)
	return sort
}

// ConstArray returns an Array value where every index maps to value.
func (ctx *Context) ConstArray(domain Sort, value Value) Array {
	res := Array(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_const_array(ctx.c, domain.c, value.impl().c)
	}))
	runtime.KeepAlive(domain)
	runtime.KeepAlive(value)
	return res
}

//go:generate go run genwrap.go -t Array $GOFILE

// Select returns the value of array x at index i.
//
// i's sort must match x's domain. The result has the sort of x's
// range.
//
//wrap:expr Select:Value x i:Value : Z3_mk_select x i

// Store returns an array y that's identical to x except that
// y.Select(i) == v.
//
// i's sort must match x's domain and v's sort must match x's range.
// The result has the same sort as x.
//
//wrap:expr Store x i:Value v:Value : Z3_mk_store x i v

// Default returns the default value of an array, for arrays that can
// be represented as finite maps plus a default value.
//
// This is useful for extracting array values interpreted by models.
//
//wrap:expr Default:Value x : Z3_mk_array_default x
