// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>

extern void goZ3ErrorHandler(Z3_context c, Z3_error_code e);
*/
import "C"

// Context is an environment for creating symbolic values and checking
// satisfiability.
//
// Nearly all interaction with Z3 is done relative to a Context.
// Values are bound to the Context that created them and cannot be
// combined with Values from other Contexts.
//
// Context is thread-safe. However, most operations block other
// operations (one notable exception is Interrupt). Hence, to do
// things in parallel, it's best to create multiple Contexts.
type Context struct {
	*contextImpl

	syms map[string]C.Z3_symbol

	// roundingMode is the current floating-point rounding mode.
	roundingMode RoundingMode

	// roundingModeAST is AST of roundingMode, or the zero value
	// if no floating-point math has been performed yet. Use rm()
	// to get the rounding mode AST.
	roundingModeAST value

	// extra contains extra values associated with this Context.
	// This must be outside contextImpl so objects that reference
	// Context (e.g., Values and Sorts) can be added to here
	// without creating a cycle and preventing finalization.
	extra map[interface{}]interface{}

	// lock protects AST reference counts and the context's last
	// error. Use Context.do to acquire this around a Z3 operation
	// and panic if the operation has an error status.
	lock sync.Mutex
}

type contextImpl struct {
	c C.Z3_context
}

//export goZ3ErrorHandler
func goZ3ErrorHandler(ctx C.Z3_context, e C.Z3_error_code) {
	msg := C.Z3_get_error_msg(ctx, e)
	// TODO: Lift the Z3 errors to better Go errors. At least wrap
	// the string in a type and consider using the error code to
	// determine which of different error types to use.
	panic(C.GoString(msg))
}

// NewContext returns a new Z3 context with the given configuration.
//
// The config argument must have been created with NewContextConfig.
// If config is nil, the default configuration is used.
func NewContext(config *Config) *Context {
	// Construct the Z3_config.
	cfg := C.Z3_mk_config()
	defer C.Z3_del_config(cfg)
	if config != nil {
		for key, val := range config.m {
			ckey, cval := C.CString(key), C.CString(fmt.Sprint(val))
			defer C.free(unsafe.Pointer(ckey))
			defer C.free(unsafe.Pointer(cval))
			C.Z3_set_param_value(cfg, ckey, cval)
		}
	}
	// Construct the Z3_context.
	impl := &contextImpl{C.Z3_mk_context_rc(cfg)}
	runtime.SetFinalizer(impl, func(impl *contextImpl) {
		C.Z3_del_context(impl.c)
	})
	ctx := &Context{
		impl,
		make(map[string]C.Z3_symbol),
		RoundToNearestEven,
		value{},
		nil,
		sync.Mutex{},
	}
	// Install an error handler that turns errors into Go panics.
	// This error handler is equivalent to a longjmp on the C++
	// side, but Z3 is actually designed to handle that, which is
	// nice because it saves us the trouble of checking the
	// context's error code all over the place.
	C.Z3_set_error_handler(ctx.c, (*C.Z3_error_handler)(C.goZ3ErrorHandler))
	return ctx
}

// NewContextConfig returns *Config for configuring a new Context.
//
// The following are commonly useful parameters:
//
//		timeout           uint    Timeout in milliseconds used for solvers (default: ∞)
//		auto_config       bool    Use heuristics to automatically select solver and configure it (default: true)
//		proof             bool    Enable proof generation (default: false)
//		model             bool    Enable model generation for solvers by default (default: true)
//	     unsat_core        bool    Enable unsat core generation for solvers by default (default: false)
//
// Most of these can be changed after a Context is created using
// Context.Config().
func NewContextConfig() *Config {
	// Based on context_params.cpp:collect_param_descrs.
	// Unfortunately, there's no way to access this from the API.
	return newConfig([]param{
		{"timeout", "uint", "Timeout in milliseconds used for solvers"},
		{"rlimit", "uint", "Resource limit used for solvers"},
		{"well_sorted_check", "bool", "Type checker"},
		{"auto_config", "bool", "Use heuristics to automatically select solver and configure it"},
		{"model_validate", "bool", "Validate models produced by solvers"},
		{"dump_models", "bool", "Dump models whenever check-sat returns sat"},
		{"trace", "bool", "Trace generation for VCC"},
		{"trace_file_name", "string", "Trace out file for VCC traces"},
		{"debug_ref_count", "bool", "Debug support for AST reference counting"},
		{"smtlib2_compliant", "bool", "Enable SMT-LIB 2.0 compliance"},
		// Solver parameters.
		{"proof", "bool", "Enable proof generation"},
		{"model", "bool", "Enable model generation for solvers"},
		{"unsat_core", "bool", "Enable unsat-core generation for solvers"},
	})
}

// Config returns a *Config object for dynamically changing ctx's
// configuration.
func (ctx *Context) Config() *Config {
	cfg := NewContextConfig()
	cfg.set = ctx.setParam
	return cfg
}

func (ctx *Context) setParam(name string, val interface{}) {
	cname, cval := C.CString(name), C.CString(fmt.Sprint(val))
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cval))
	ctx.do(func() {
		C.Z3_update_param_value(ctx.c, cname, cval)
	})
}

// Interrupt stops the current solver, simplifier, or tactic being
// executed by ctx.
func (ctx *Context) Interrupt() {
	C.Z3_interrupt(ctx.c)
	runtime.KeepAlive(ctx)
}

// Extra returns the "extra" data associated with key, or nil if there
// is no data associated with key.
func (ctx *Context) Extra(key interface{}) interface{} {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	return ctx.extra[key]
}

// SetExtra associates key with value in ctx's "extra" data. This can
// be used by other packages to associate other data with ctx, such as
// caches. key must support comparison.
func (ctx *Context) SetExtra(key, value interface{}) {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	if value == nil {
		if ctx.extra != nil {
			delete(ctx.extra, key)
		}
	} else {
		if ctx.extra == nil {
			ctx.extra = make(map[interface{}]interface{})
		}
		ctx.extra[key] = value
	}
}

// do calls f with a per-context lock held.
//
// Unfortunately, we can't just say that Contexts are not thread-safe
// because we can't help but run finalizers asynchronously, which
// means we need to synchronize both reference counts and the
// per-context last error state.
func (ctx *Context) do(f func()) {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	f()
}

// symbol interns name as a Z3 symbol.
func (ctx *Context) symbol(name string) C.Z3_symbol {
	if sym, ok := ctx.syms[name]; ok {
		return sym
	}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var sym C.Z3_symbol
	ctx.do(func() {
		sym = C.Z3_mk_string_symbol(ctx.c, cname)
		ctx.syms[name] = sym
	})
	return sym
}
