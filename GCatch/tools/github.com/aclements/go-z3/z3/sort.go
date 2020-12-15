// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"runtime"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Sort represents the type of a symbolic value. A Sort can be a basic
// type such as "int" or "bool" or a parameterized type such as
// "bit-vector of 30 bits" or "array from int to bool".
type Sort struct {
	*sortImpl
	noEq
}

// sortImpl wraps the underlying C.Z3_sort. This is separate from Sort
// so a finalizer can be attached to this without exposing it to the
// user.
type sortImpl struct {
	ctx  *Context
	c    C.Z3_sort
	kind Kind
}

// wrapSort wraps a C Z3_sort as a Go Sort. This must be called with
// the ctx.lock held.
func wrapSort(ctx *Context, c C.Z3_sort, kind Kind) Sort {
	C.Z3_inc_ref(ctx.c, C.Z3_sort_to_ast(ctx.c, c))
	if kind == KindUnknown {
		kind = Kind(C.Z3_get_sort_kind(ctx.c, c))
	}
	impl := &sortImpl{ctx, c, kind}
	runtime.SetFinalizer(impl, func(impl *sortImpl) {
		impl.ctx.do(func() {
			C.Z3_dec_ref(impl.ctx.c, C.Z3_sort_to_ast(impl.ctx.c, impl.c))
		})
	})
	return Sort{impl, noEq{}}
}

// Context returns the Context that created sort.
func (s Sort) Context() *Context {
	if s.sortImpl == nil {
		return nil
	}
	return s.ctx
}

// String returns a string representation of s.
func (s Sort) String() string {
	var res string
	s.ctx.do(func() {
		res = C.GoString(C.Z3_sort_to_string(s.ctx.c, s.c))
	})
	runtime.KeepAlive(s)
	return res
}

// Kind returns s's kind.
func (s Sort) Kind() Kind {
	return s.kind
}

// BVSize returns the bit size of a bit-vector sort.
func (s Sort) BVSize() int {
	var size int
	s.ctx.do(func() {
		size = int(C.Z3_get_bv_sort_size(s.ctx.c, s.c))
	})
	runtime.KeepAlive(s)
	return size
}

// FloatSize returns the number of exponent and significand bits in s.
func (s Sort) FloatSize() (ebits, sbits int) {
	s.ctx.do(func() {
		ebits = int(C.Z3_fpa_get_ebits(s.ctx.c, s.c))
		sbits = int(C.Z3_fpa_get_sbits(s.ctx.c, s.c))
	})
	runtime.KeepAlive(s)
	return
}

// DomainAndRange returns the domain and range of an array sort.
func (s Sort) DomainAndRange() (domain, range_ Sort) {
	s.ctx.do(func() {
		domain = wrapSort(s.ctx, C.Z3_get_array_sort_domain(s.ctx.c, s.c), KindUnknown)
		range_ = wrapSort(s.ctx, C.Z3_get_array_sort_range(s.ctx.c, s.c), KindUnknown)
	})
	runtime.KeepAlive(s)
	return
}

// AsAST returns the AST representation of s.
func (s Sort) AsAST() AST {
	var ast AST
	s.ctx.do(func() {
		ast = wrapAST(s.ctx, C.Z3_sort_to_ast(s.ctx.c, s.c))
	})
	runtime.KeepAlive(s)
	return ast
}
