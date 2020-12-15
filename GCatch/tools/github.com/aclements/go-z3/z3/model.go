// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "runtime"

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// A Model is a binding of constants that satisfies a set of formulas.
type Model struct {
	*modelImpl
	noEq
}

type modelImpl struct {
	ctx *Context
	c   C.Z3_model
}

// wrapModel wraps a C Z3_model as a Go Model. This must be called
// with the ctx.lock held.
func wrapModel(ctx *Context, c C.Z3_model) *Model {
	impl := &modelImpl{ctx, c}
	C.Z3_model_inc_ref(ctx.c, c)
	runtime.SetFinalizer(impl, func(impl *modelImpl) {
		impl.ctx.do(func() {
			C.Z3_model_dec_ref(impl.ctx.c, impl.c)
		})
	})
	return &Model{impl, noEq{}}
}

// Eval evaluates val using the concrete interpretations of constants
// and functions in model m.
//
// If completion is true, it will assign interpretations for any
// constants or functions that currently don't have an interpretation
// in m. Otherwise, the resulting value may not be concrete.
//
// Eval returns nil if val cannot be evaluated. This can happen if val
// contains a quantifier or is type-incorrect, or if m is a partial
// model (that is, the option MODEL_PARTIAL was set to true).
func (m *Model) Eval(val Value, completion bool) Value {
	var ok bool
	var ast AST
	m.ctx.do(func() {
		var cast C.Z3_ast
		ok = z3ToBool(C.Z3_model_eval(m.ctx.c, m.c, val.impl().c, boolToZ3(completion), &cast))
		if ok {
			ast = wrapAST(m.ctx, cast)
		}
	})
	runtime.KeepAlive(m)
	runtime.KeepAlive(val)
	if !ok {
		return nil
	}
	return ast.AsValue()
}

// String returns a string representation of m.
func (m *Model) String() string {
	var res string
	m.ctx.do(func() {
		res = C.GoString(C.Z3_model_to_string(m.ctx.c, m.c))
	})
	runtime.KeepAlive(m)
	return res
}

// Sorts returns the uninterpreted sorts that m assigns an
// interpretation to.
//
// Each of these interpretations is a finite set of distinct values
// known as the "universe" of the sort. These values can be retrieved
// with SortUniverse.
func (m *Model) Sorts() []Sort {
	var res []Sort
	m.ctx.do(func() {
		n := C.Z3_model_get_num_sorts(m.ctx.c, m.c)
		res = make([]Sort, n)
		for i := C.uint(0); i < n; i++ {
			csort := C.Z3_model_get_sort(m.ctx.c, m.c, i)
			res[i] = wrapSort(m.ctx, csort, KindUninterpreted)
		}
	})
	runtime.KeepAlive(m)
	return res
}

// SortUniverse returns the interpretation of s in m. s must be in the
// set returned by m.Sorts.
//
// The interpretation of s is a finite set of distinct values of sort
// s.
func (m *Model) SortUniverse(s Sort) []Uninterpreted {
	var cvec C.Z3_ast_vector
	var n C.uint
	m.ctx.do(func() {
		cvec = C.Z3_model_get_sort_universe(m.ctx.c, m.c, s.c)
		C.Z3_ast_vector_inc_ref(m.ctx.c, cvec)
		n = C.Z3_ast_vector_size(m.ctx.c, cvec)
	})
	defer m.ctx.do(func() { C.Z3_ast_vector_dec_ref(m.ctx.c, cvec) })
	res := make([]Uninterpreted, n)
	for i := C.uint(0); i < n; i++ {
		res[i] = Uninterpreted(wrapValue(m.ctx, func() C.Z3_ast {
			return C.Z3_ast_vector_get(m.ctx.c, cvec, i)
		}))
	}
	runtime.KeepAlive(m)
	return res
}
