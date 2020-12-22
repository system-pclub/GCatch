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
#include <stdlib.h>
*/
import "C"

// FuncDecl is a function declaration.
//
// A FuncDecl can represent either a interpreted function like "+" or
// an uninterpreted function created by Context.FuncDecl.
//
// A FuncDecl can be applied to a set of arguments to create a Value
// representing the result of the function.
type FuncDecl struct {
	*funcDeclImpl
	noEq
}

type funcDeclImpl struct {
	ctx *Context
	c   C.Z3_func_decl
}

// wrapFuncDecl wraps a C Z3_func_decl as a Go FuncDecl. This must be
// called with the ctx.lock held.
func wrapFuncDecl(ctx *Context, c C.Z3_func_decl) FuncDecl {
	impl := &funcDeclImpl{ctx, c}
	C.Z3_inc_ref(ctx.c, C.Z3_func_decl_to_ast(ctx.c, c))
	runtime.SetFinalizer(impl, func(impl *funcDeclImpl) {
		impl.ctx.do(func() {
			C.Z3_dec_ref(impl.ctx.c, C.Z3_func_decl_to_ast(impl.ctx.c, impl.c))
		})
	})
	runtime.KeepAlive(ctx)
	return FuncDecl{impl, noEq{}}
}

// FuncDecl creates an uninterpreted function named "name".
//
// In contrast with an interpreted function like "+", an uninterpreted
// function is only assigned an interpretation in a particular model,
// and different models may assign different interpretations.
func (ctx *Context) FuncDecl(name string, domain []Sort, range_ Sort) FuncDecl {
	sym := ctx.symbol(name)
	cdomain := make([]C.Z3_sort, len(domain))
	for i, sort := range domain {
		cdomain[i] = sort.c
	}
	var funcdecl FuncDecl
	ctx.do(func() {
		var cdp *C.Z3_sort
		if len(cdomain) > 0 {
			cdp = &cdomain[0]
		}
		funcdecl = wrapFuncDecl(ctx, C.Z3_mk_func_decl(ctx.c, sym, C.uint(len(cdomain)), cdp, range_.c))
	})
	runtime.KeepAlive(domain)
	runtime.KeepAlive(range_)
	return funcdecl
}

// FreshFuncDecl creates a fresh uninterpreted function distinct from
// all other functions.
func (ctx *Context) FreshFuncDecl(prefix string, domain []Sort, range_ Sort) FuncDecl {
	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	cdomain := make([]C.Z3_sort, len(domain))
	for i, sort := range domain {
		cdomain[i] = sort.c
	}
	var funcdecl FuncDecl
	ctx.do(func() {
		var cdp *C.Z3_sort
		if len(cdomain) > 0 {
			cdp = &cdomain[0]
		}
		funcdecl = wrapFuncDecl(ctx, C.Z3_mk_fresh_func_decl(ctx.c, cprefix, C.uint(len(cdomain)), cdp, range_.c))
	})
	runtime.KeepAlive(domain)
	runtime.KeepAlive(range_)
	return funcdecl
}

// Context returns the Context that created f.
func (f FuncDecl) Context() *Context {
	if f.funcDeclImpl == nil {
		return nil
	}
	return f.ctx
}

// String returns a string representation of f.
func (f FuncDecl) String() string {
	var res string
	f.ctx.do(func() {
		res = C.GoString(C.Z3_func_decl_to_string(f.ctx.c, f.c))
	})
	runtime.KeepAlive(f)
	return res
}

// AsAST returns the AST representation of f.
func (f FuncDecl) AsAST() AST {
	var ast AST
	f.ctx.do(func() {
		ast = wrapAST(f.ctx, C.Z3_func_decl_to_ast(f.ctx.c, f.c))
	})
	runtime.KeepAlive(f)
	return ast
}

// Apply creates a Value representing the result of applying f to
// args.
//
// The sorts of args must be the domain of f. The sort of the
// resulting value will be f's range.
func (f FuncDecl) Apply(args ...Value) Value {
	cargs := make([]C.Z3_ast, len(args))
	for i, arg := range args {
		cargs[i] = arg.impl().c
	}
	val := wrapValue(f.ctx, func() C.Z3_ast {
		var cap *C.Z3_ast
		if len(cargs) > 0 {
			cap = &cargs[0]
		}
		return C.Z3_mk_app(f.ctx.c, f.c, C.uint(len(cargs)), cap)
	})
	runtime.KeepAlive(f)
	runtime.KeepAlive(cargs)
	return val.lift(KindUnknown)
}

// Map applies f to each value in each of the args array.
//
// Given that f has sort range_1, ..., range_n -> range, args[i] must
// have array sort [domain -> range_i]. The result will have array
// sort [domain -> range].
func (f FuncDecl) Map(args ...Array) Array {
	cargs := make([]C.Z3_ast, len(args))
	for i, arg := range args {
		cargs[i] = arg.impl().c
	}
	val := wrapValue(f.ctx, func() C.Z3_ast {
		var cap *C.Z3_ast
		if len(cargs) > 0 {
			cap = &cargs[0]
		}
		return C.Z3_mk_map(f.ctx.c, f.c, C.uint(len(cargs)), cap)
	})
	runtime.KeepAlive(f)
	runtime.KeepAlive(args)
	return Array(val)
}

// TODO: Lots of accessors
