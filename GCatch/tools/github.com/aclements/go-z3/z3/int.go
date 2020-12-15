// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"
import "math/big"

// Int is a symbolic value representing an integer with infinite precision.
//
// Int implements Value.
type Int value

func init() {
	kindWrappers[KindInt] = func(x value) Value {
		return Int(x)
	}
}

// IntSort returns the integer sort.
func (ctx *Context) IntSort() Sort {
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_int_sort(ctx.c), KindInt)
	})
	return sort
}

// IntConst returns a int constant named "name".
func (ctx *Context) IntConst(name string) Int {
	return ctx.Const(name, ctx.IntSort()).(Int)
}

// AsInt64 returns the value of lit as an int64. If lit is not a
// literal, it returns 0, false, false. If lit is a literal, but its
// value cannot be represented as an int64, it returns 0, true, false.
func (lit Int) AsInt64() (val int64, isLiteral, ok bool) {
	return lit.asInt64()
}

// AsUint64 is like AsInt64, but returns a uint64 and fails if lit
// cannot be represented as a uint64.
func (lit Int) AsUint64() (val uint64, isLiteral, ok bool) {
	return lit.asUint64()
}

// AsBigInt returns the value of lit as a math/big.Int. If lit is not
// a literal, it returns nil, false.
func (lit Int) AsBigInt() (val *big.Int, isConst bool) {
	return lit.asBigInt()
}

//go:generate go run genwrap.go -t Int $GOFILE intreal.go

// Div returns the floor of l / r.
//
// If r is 0, the result is unconstrained.
//
// Note that this differs from Go division: Go rounds toward zero
// (truncated division), whereas this rounds toward -inf.
//
//wrap:expr Div Z3_mk_div l r

// Mod returns modulus of l / r.
//
// The sign of the result follows the sign of r.
//
//wrap:expr Mod Z3_mk_mod l r

// Rem returns remainder of l / r.
//
// The sign of the result follows the sign of l.
//
// Note that this differs subtly from Go's remainder operator because
// this is based floored division rather than truncated division.
//
//wrap:expr Rem Z3_mk_rem l r

// ToReal converts l to sort Real.
//
//wrap:expr ToReal:Real Z3_mk_int2real l

// ToBV converts l to a bit-vector of width bits.
//
//wrap:expr ToBV:BV l bits:int : Z3_mk_int2bv bits:unsigned l
