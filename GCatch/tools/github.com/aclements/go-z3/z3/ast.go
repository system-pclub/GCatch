// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"runtime"
	"unsafe"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// An AST is the abstract syntax tree of an expression, sort, etc.
type AST struct {
	*astImpl
	noEq
}

type astImpl struct {
	ctx *Context
	c   C.Z3_ast
}

// wrapAST wraps a C Z3_ast as a Go AST. This must be called with the
// ctx.lock held.
func wrapAST(ctx *Context, c C.Z3_ast) AST {
	impl := &astImpl{ctx, c}
	// Note that, even if c was just returned by an allocation
	// function, we're still responsible for incrementing its
	// reference count. This is weird, but also nice because we
	// can wrap any AST that comes out of the Z3 API, even if
	// we've already wrapped it, and the reference count will
	// protect the underlying object no matter what happens to the
	// Go wrappers.
	//
	// This must be done atomically with the allocating function.
	// If we allocate two objects without incrementing the
	// refcount on the first, Z3 will reclaim the first object!
	C.Z3_inc_ref(ctx.c, c)
	runtime.SetFinalizer(impl, func(impl *astImpl) {
		impl.ctx.do(func() {
			C.Z3_dec_ref(impl.ctx.c, impl.c)
		})
	})
	return AST{impl, noEq{}}
}

// Context returns the Context that created ast.
func (ast AST) Context() *Context {
	if ast.astImpl == nil {
		return nil
	}
	return ast.ctx
}

// Equal returns true if ast and o are identical ASTs.
func (ast AST) Equal(o AST) bool {
	// Sadly, while AST equality is just pointer equality on the
	// underlying C pointers, it's impossible to expose this as Go
	// equality because we need a Go pointer to attach the
	// finalizer to. We can't make *that* pointer 1:1 with the C
	// pointer without making the object permanently live.
	var out bool
	ast.ctx.do(func() {
		out = z3ToBool(C.Z3_is_eq_ast(ast.ctx.c, ast.c, o.c))
	})
	runtime.KeepAlive(ast)
	runtime.KeepAlive(o)
	return out
}

// String returns ast as an S-expression.
func (ast AST) String() string {
	var res string
	ast.ctx.do(func() {
		res = C.GoString(C.Z3_ast_to_string(ast.ctx.c, ast.c))
	})
	runtime.KeepAlive(ast)
	return res
}

// Hash returns a hash of ast. Structurally identical ASTs will have
// the same hash code.
func (ast AST) Hash() uint64 {
	var res uint64
	ast.ctx.do(func() {
		res = uint64(C.Z3_get_ast_hash(ast.ctx.c, ast.c))
	})
	runtime.KeepAlive(ast)
	return res
}

// ID returns the unique identifier for ast. Within a Context, two
// ASTs have the same ID if and only if they are Equal.
func (ast AST) ID() uint64 {
	var res uint64
	ast.ctx.do(func() {
		res = uint64(C.Z3_get_ast_id(ast.ctx.c, ast.c))
	})
	runtime.KeepAlive(ast)
	return res
}

// Translate copies ast into the target Context.
func (ast AST) Translate(target *Context) AST {
	var res AST
	target.do(func() {
		res = wrapAST(target, C.Z3_translate(ast.ctx.c, ast.c, target.c))
	})
	runtime.KeepAlive(ast)
	return res
}

// Kind returns ast's kind.
func (ast AST) Kind() ASTKind {
	var res ASTKind
	ast.ctx.do(func() {
		res = ASTKind(C.Z3_get_ast_kind(ast.ctx.c, ast.c))
	})
	runtime.KeepAlive(ast)
	return res
}

// AsValue returns this AST as a symbolic value.
//
// It panics if ast is not a value expression. That is, ast must have
// Kind ASTKindApp, ASTKindNumeral, ASTKindVar, or ASTKindQuantifier.
func (ast AST) AsValue() Value {
	kind := ast.Kind()
	switch kind {
	case ASTKindApp, ASTKindNumeral, ASTKindVar, ASTKindQuantifier:
		return value{(*valueImpl)(ast.astImpl), noEq{}}.lift(KindUnknown)
	}
	panic("AST has kind " + kind.String() + " which is not a value")
}

// AsSort returns this AST as a sort.
//
// It panics if ast is not a sort expression. That is, ast must have
// Kind ASTKindSort.
func (ast AST) AsSort() Sort {
	if kind := ast.Kind(); kind != ASTKindSort {
		panic("AST has kind " + kind.String() + ", not ASTKindSort")
	}
	// Weirdly, Z3 doesn't provide an API for this. But these are
	// all just casts.
	var sort Sort
	ast.ctx.do(func() {
		csort := C.Z3_sort(unsafe.Pointer(ast.c))
		sort = wrapSort(ast.ctx, csort, KindUnknown)
	})
	runtime.KeepAlive(ast)
	return sort
}

// AsFuncDecl returns this AST as a FuncDecl.
//
// It panics if ast is not a function declaration expression. That is,
// ast must have Kind ASTKindFuncDecl.
func (ast AST) AsFuncDecl() FuncDecl {
	if kind := ast.Kind(); kind != ASTKindFuncDecl {
		panic("AST has kind " + kind.String() + ", not ASTKindFuncDecl")
	}
	var funcdecl FuncDecl
	ast.ctx.do(func() {
		funcdecl = wrapFuncDecl(ast.ctx, C.Z3_to_func_decl(ast.ctx.c, ast.c))
	})
	runtime.KeepAlive(ast)
	return funcdecl
}
