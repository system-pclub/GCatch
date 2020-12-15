// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"math/big"
	"runtime"
	"unsafe"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Real is a symbolic value representing a real number.
//
// Real implements Value.
type Real value

func init() {
	kindWrappers[KindReal] = func(x value) Value {
		return Real(x)
	}
}

// RealSort returns the real sort.
func (ctx *Context) RealSort() Sort {
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_real_sort(ctx.c), KindReal)
	})
	return sort
}

// RealConst returns a int constant named "name".
func (ctx *Context) RealConst(name string) Real {
	return ctx.Const(name, ctx.RealSort()).(Real)
}

// FromBigRat returns a real literal whose value is val.
func (ctx *Context) FromBigRat(val *big.Rat) Real {
	cstr := C.CString(val.String())
	defer C.free(unsafe.Pointer(cstr))
	sort := ctx.RealSort()
	sval := wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_numeral(ctx.c, cstr, sort.c)
	})
	runtime.KeepAlive(sort)
	return Real(sval)
}

// AsRat returns the value of lit as a numerator and denominator Int
// literals. If lit is not a literal or is not rational, it returns
// false for isLiteralRational. To round an arbitrary real to be
// rational, see method Real.Approx.
func (lit Real) AsRat() (numer, denom Int, isLiteralRational bool) {
	if lit.astKind() != C.Z3_NUMERAL_AST {
		// Algebraic literals do not count as Z3_NUMERAL_AST,
		// so this gets all the cases we need.
		return Int{}, Int{}, false
	}
	numer = Int(wrapValue(lit.ctx, func() C.Z3_ast {
		return C.Z3_get_numerator(lit.ctx.c, lit.c)
	}))
	denom = Int(wrapValue(lit.ctx, func() C.Z3_ast {
		return C.Z3_get_denominator(lit.ctx.c, lit.c)
	}))
	runtime.KeepAlive(lit)
	return numer, denom, true
}

// AsBigRat returns the value of lit as a math/big.Rat. If lit is not
// a literal or is not rational, it returns nil, false.
func (lit Real) AsBigRat() (val *big.Rat, isLiteralRational bool) {
	numer, denom, isLiteralRational := lit.AsRat()
	if !isLiteralRational {
		return nil, false
	}
	var rat big.Rat
	bigNumer, _ := numer.AsBigInt()
	bigDenom, _ := denom.AsBigInt()
	rat.SetFrac(bigNumer, bigDenom)
	return &rat, true
}

// Approx approximates lit as two rational literals, where the
// difference between lower and upper is less than 1/10**precision. If
// lit is not an irrational literal, it returns false for
// isLiteralIrrational.
func (lit Real) Approx(precision int) (lower, upper Real, isLiteralIrrational bool) {
	var isAlgebraicNumber bool
	lit.ctx.do(func() {
		// Despite the name, this really means an *irrational*
		// algebraic number.
		isAlgebraicNumber = z3ToBool(C.Z3_is_algebraic_number(lit.ctx.c, lit.c))
	})
	if !isAlgebraicNumber {
		return Real{}, Real{}, false
	}
	lower = Real(wrapValue(lit.ctx, func() C.Z3_ast {
		return C.Z3_get_algebraic_number_lower(lit.ctx.c, lit.c, C.unsigned(precision))
	}))
	upper = Real(wrapValue(lit.ctx, func() C.Z3_ast {
		return C.Z3_get_algebraic_number_upper(lit.ctx.c, lit.c, C.unsigned(precision))
	}))
	runtime.KeepAlive(lit)
	return lower, upper, true
}

// TODO: AsBigFloat? AsFloat64? AsFloat32? I don't actually know how
// to implement those without potentially double rounding.

//go:generate go run genwrap.go -t Real $GOFILE intreal.go

// Div returns l / r.
//
// If r is 0, the result is unconstrained.
//
//wrap:expr Div Z3_mk_div l r

// ToInt returns the floor of l as sort Int.
//
// Note that this is not truncation. For example, ToInt(-1.3) is -2.
//
//wrap:expr ToInt:Int Z3_mk_real2int l

// IsInt returns a Value that is true if l has no fractional part.
//
//wrap:expr IsInt:Bool Z3_mk_is_int l

// ToFloat converts l into a floating-point number.
//
// If necessary, the result will be rounded according to the current
// rounding mode.
//
//wrap:expr ToFloat:Float l s:Sort : Z3_mk_fpa_to_fp_real @rm l s

// ToFloatExp converts l into a floating-point number l*2^exp.
//
// If necessary, the result will be rounded according to the current
// rounding mode.
//
//wrap:expr ToFloatExp:Float l exp:Int s:Sort : Z3_mk_fpa_to_fp_int_real @rm exp l s
