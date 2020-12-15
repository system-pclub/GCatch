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

// TODO: Should the various Value types disallow explicit conversion?
// Right now they all have the same underlying type.

// An Value is a symbolic value (a Z3 expression).
//
// This package exports a concrete type for each different kind of
// value, such as Bool, BV, and Int. These concrete types provide
// methods for deriving other values.
//
// Having separate types for each kind separates which methods can be
// applied to which kind of value and provides some level of static
// type safety. However, by no means does this fully capture Z3's type
// system, so dynamic type checking can still fail.
type Value interface {
	// AsAST returns the abstract syntax tree underlying this
	// Value.
	AsAST() AST

	// Sort returns this value's sort.
	Sort() Sort

	// Context returns the Context that created this Value.
	Context() *Context

	// String returns an S-expression representation of this value.
	String() string

	astKind() C.Z3_ast_kind
	impl() *valueImpl
}

type noEq struct {
	_ [0]func()
}

// value is a general wrapper for the Z3_ast type. Expression values
// are implemented as public types corresponding to Z3 sorts that are
// named types for value.
type value struct {
	// *valueImpl is the internal state of the value. This is wrapped
	// and unexported so we can attach a finalizer to the valueImpl
	// object without any possibility of user code copying the
	// underlying wrapper and breaking our tracking.
	*valueImpl

	// noEq prevents user code from directly comparing values for
	// equality.
	noEq
}

type valueImpl astImpl

func wrapValue(ctx *Context, ctor func() C.Z3_ast) value {
	var val value
	ctx.do(func() {
		cast := ctor()
		val = value{(*valueImpl)(wrapAST(ctx, cast).astImpl), noEq{}}
	})
	return val
}

// lift wraps x in the appropriate Value type. kind must be x's kind if
// known or otherwise SortUnknown.
func (x value) lift(kind Kind) Value {
	if kind == KindUnknown {
		kind = x.Sort().Kind()
	}
	wrap, ok := kindWrappers[kind]
	if !ok {
		panic("value has unknown kind " + kind.String())
	}
	return wrap(x)
}

// Const returns a constant named "name" of the given sort. This
// constant will be same as all other constants created with this
// name.
func (ctx *Context) Const(name string, sort Sort) Value {
	sym := ctx.symbol(name)
	val := wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_const(ctx.c, sym, sort.c)
	})
	runtime.KeepAlive(sort)
	return val.lift(sort.Kind())
}

// FreshConst returns a constant that is distinct from all other
// constants. The name will begin with "prefix".
func (ctx *Context) FreshConst(prefix string, sort Sort) Value {
	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	val := wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fresh_const(ctx.c, cprefix, sort.c)
	})
	runtime.KeepAlive(sort)
	return val.lift(sort.Kind())
}

// FromBigInt returns a literal whose value is val. sort must have
// kind int, real, finite-domain, bit-vector, or float.
func (ctx *Context) FromBigInt(val *big.Int, sort Sort) Value {
	if sort.Kind() == KindFloatingPoint {
		return ctx.floatFromBigInt(val, sort)
	}
	cstr := C.CString(val.Text(10))
	defer C.free(unsafe.Pointer(cstr))
	sval := wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_numeral(ctx.c, cstr, sort.c)
	})
	runtime.KeepAlive(sort)
	return sval.lift(sort.Kind())
}

// TODO: FromBigFloat for real and float sorts (or maybe just float
// sorts since we can deduce the sbits from the big.Float?).

// FromInt returns a literal whose value is val. sort must have kind
// int, real, finite-domain, bit-vector, or float.
func (ctx *Context) FromInt(val int64, sort Sort) Value {
	if sort.Kind() == KindFloatingPoint {
		return ctx.floatFromInt(val, sort)
	}
	sval := wrapValue(ctx, func() C.Z3_ast {
		// Z3_mk_int64 doesn't say real sorts are accepted,
		// but the C++ bindings use it for reals.
		//return C.Z3_mk_int64(ctx.c, C.__int64(val), sort.c)
		return C.Z3_mk_int64(ctx.c, C.int64_t(val), sort.c)
	})
	runtime.KeepAlive(sort)
	return sval.lift(sort.Kind())
}

func (expr *valueImpl) impl() *valueImpl {
	return expr
}

// Context returns the Context that created expr.
func (expr *valueImpl) Context() *Context {
	if expr == nil {
		return nil
	}
	return expr.ctx
}

// String returns a string representation of expr.
func (expr *valueImpl) String() string {
	var res string
	expr.ctx.do(func() {
		res = C.GoString(C.Z3_ast_to_string(expr.ctx.c, expr.c))
	})
	runtime.KeepAlive(expr)
	return res
}

// AsAST returns the abstract syntax tree underlying expr.
func (expr *valueImpl) AsAST() AST {
	ast := AST{(*astImpl)(expr), noEq{}}
	runtime.KeepAlive(expr)
	return ast
}

// Sort returns expr's sort.
func (expr *valueImpl) Sort() Sort {
	var sort Sort
	expr.ctx.do(func() {
		sort = wrapSort(expr.ctx, C.Z3_get_sort(expr.ctx.c, expr.c), KindUnknown)
	})
	runtime.KeepAlive(expr)
	return sort
}

func (expr *valueImpl) astKind() C.Z3_ast_kind {
	var ckind C.Z3_ast_kind
	expr.ctx.do(func() {
		ckind = C.Z3_get_ast_kind(expr.ctx.c, expr.c)
	})
	runtime.KeepAlive(expr)
	return ckind
}

func (expr *valueImpl) asBigInt() (val *big.Int, isLiteral bool) {
	switch expr.Sort().Kind() {
	default:
		panic("sort " + expr.Sort().String() + " cannot be represented as a big.Int")
	case KindInt, KindBV:
	}
	if expr.astKind() != C.Z3_NUMERAL_AST {
		return nil, false
	}
	var str string
	expr.ctx.do(func() {
		cstr := C.Z3_get_numeral_string(expr.ctx.c, expr.c)
		str = C.GoString(cstr)
	})
	var v big.Int
	if _, ok := v.SetString(str, 10); !ok {
		panic("failed to parse numeral string")
	}
	return &v, true
}

func (expr *valueImpl) asInt64() (val int64, isLiteral, ok bool) {
	switch expr.Sort().Kind() {
	default:
		panic("sort " + expr.Sort().String() + " cannot be represented as an int64")
	case KindInt, KindBV:
	}
	if expr.astKind() != C.Z3_NUMERAL_AST {
		return 0, false, false
	}
	//var cval C.__int64
	var cval C.int64_t
	expr.ctx.do(func() {
		ok = z3ToBool(C.Z3_get_numeral_int64(expr.ctx.c, expr.c, &cval))
	})
	return int64(cval), true, ok
}

func (expr *valueImpl) asUint64() (val uint64, isLiteral, ok bool) {
	switch expr.Sort().Kind() {
	default:
		panic("sort " + expr.Sort().String() + " cannot be represented as an int64")
	case KindInt, KindBV:
	}
	if expr.astKind() != C.Z3_NUMERAL_AST {
		return 0, false, false
	}
	//var cval C.__uint64
	var cval C.uint64_t
	expr.ctx.do(func() {
		ok = z3ToBool(C.Z3_get_numeral_uint64(expr.ctx.c, expr.c, &cval))
	})
	return uint64(cval), true, ok
}

func (expr *valueImpl) isAppOf(k C.Z3_decl_kind) bool {
	var res bool
	expr.ctx.do(func() {
		res = z3ToBool(C.Z3_is_app(expr.ctx.c, expr.c)) && C.Z3_get_decl_kind(expr.ctx.c, C.Z3_get_app_decl(expr.ctx.c, C.Z3_to_app(expr.ctx.c, expr.c))) == k
	})
	runtime.KeepAlive(expr)
	return res
}
