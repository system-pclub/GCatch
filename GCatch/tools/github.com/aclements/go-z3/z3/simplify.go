// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// Simplify simplifies expression x.
//
// The config argument must have been created with NewSimplifyConfig.
// If config is nil, the default configuration is used.
//
// The resulting expression will have the same sort and value as x,
// but with a simpler AST.
func (ctx *Context) Simplify(x Value, config *Config) Value {
	var cparams C.Z3_params
	if config != nil {
		cparams = config.toC(ctx)
		defer C.Z3_params_dec_ref(ctx.c, cparams)
	}
	return wrapValue(ctx, func() C.Z3_ast {
		if config == nil {
			return C.Z3_simplify(ctx.c, x.impl().c)
		} else {
			return C.Z3_simplify_ex(ctx.c, x.impl().c, cparams)
		}
	}).lift(KindUnknown)
}

// NewSimplifyConfig returns *Config for configuring the simplifier.
func NewSimplifyConfig(ctx *Context) *Config {
	// TODO: Get the Z3_param_descr.
	return newConfig(nil)
}
