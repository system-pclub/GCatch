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
import "runtime"

// Bool is a symbolic value representing "true" or "false".
//
// Bool implements Value.
type Bool value

func init() {
	kindWrappers[KindBool] = func(x value) Value {
		return Bool(x)
	}
}

// BoolSort returns the boolean sort.
func (ctx *Context) BoolSort() Sort {
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_bool_sort(ctx.c), KindBool)
	})
	return sort
}

// FromBool returns a boolean value with value val.
func (ctx *Context) FromBool(val bool) Bool {
	return Bool(wrapValue(ctx, func() C.Z3_ast {
		if val {
			return C.Z3_mk_true(ctx.c)
		} else {
			return C.Z3_mk_false(ctx.c)
		}
	}))
}

// BoolConst returns a boolean constant named "name".
func (ctx *Context) BoolConst(name string) Bool {
	return ctx.Const(name, ctx.BoolSort()).(Bool)
}

// AsBool returns the value of l as a Go bool. If l is not a literal,
// AsBool returns false, false.
func (l Bool) AsBool() (val bool, isLiteral bool) {
	var res C.Z3_lbool
	l.ctx.do(func() {
		res = C.Z3_get_bool_value(l.ctx.c, l.c)
	})
	runtime.KeepAlive(l)
	return res == C.Z3_L_TRUE, res != C.Z3_L_UNDEF
}

//go:generate go run genwrap.go -t Bool $GOFILE

// Distinct returns a Value that is true if no two vals are equal.
//
// All Values must have the same sort.
//
//wrap:expr Distinct ctx:*Context vals...:Value : Z3_mk_distinct vals...

// Not returns the boolean negation of l.
//
//wrap:expr Not Z3_mk_not l

// IfThenElse returns a Value equal to cons if cond is true, otherwise
// alt.
//
// cons and alt must have the same sort. The result will have the same
// sort as cons and alt.
//
//wrap:expr IfThenElse:Value cond cons:Value alt:Value : Z3_mk_ite cond cons alt

// Iff returns a Value that is true if l and r are equal (l
// if-and-only-if r).
//
//wrap:expr Iff Z3_mk_iff l r

// Implies returns a Value that is true if l implies r.
//
//wrap:expr Implies Z3_mk_implies l r

// Xor returns a Value that is true if l xor r.
//
//wrap:expr Xor Z3_mk_xor l r

// And returns a Value that is true if l and all arguments are true.
//
//wrap:expr And Z3_mk_and l r...

// Or returns a Value that is true if l or any argument is true.
//
//wrap:expr Or Z3_mk_or l r...
