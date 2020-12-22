// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// Config stores a set of configuration parameters. Configs are used
// to configure many different objects in Z3.
type Config struct {
	m   map[string]interface{}
	set func(name string, value interface{})
}

type param struct {
	name, typ, description string
}

func newConfig(desc []param) *Config {
	// TODO: API to access the parameter descriptions.
	return &Config{m: make(map[string]interface{})}
}

func (p *Config) SetBool(name string, value bool) *Config {
	if p.set != nil {
		p.set(name, value)
	} else {
		p.m[name] = value
	}
	return p
}

func (p *Config) SetString(name, value string) *Config {
	if p.set != nil {
		p.set(name, value)
	} else {
		p.m[name] = value
	}
	return p
}

func (p *Config) SetUint(name string, value uint) *Config {
	if p.set != nil {
		p.set(name, value)
	} else {
		p.m[name] = value
	}
	return p
}

func (p *Config) SetFloat(name string, value float64) *Config {
	if p.set != nil {
		p.set(name, value)
	} else {
		p.m[name] = value
	}
	return p
}

func (p *Config) toC(ctx *Context) C.Z3_params {
	var c C.Z3_params
	ctx.do(func() {
		c = C.Z3_mk_params(ctx.c)
		C.Z3_params_inc_ref(ctx.c, c)
	})
	ok := false
	defer func() {
		if !ok {
			C.Z3_params_dec_ref(ctx.c, c)
		}
	}()
	for k, v := range p.m {
		ck := ctx.symbol(k)
		switch v := v.(type) {
		case bool:
			ctx.do(func() {
				C.Z3_params_set_bool(ctx.c, c, ck, boolToZ3(v))
			})
		case string:
			cv := ctx.symbol(v)
			ctx.do(func() {
				C.Z3_params_set_symbol(ctx.c, c, ck, cv)
			})
		case uint:
			ctx.do(func() {
				C.Z3_params_set_uint(ctx.c, c, ck, C.unsigned(v))
			})
		case float64:
			ctx.do(func() {
				C.Z3_params_set_double(ctx.c, c, ck, C.double(v))
			})
		default:
			panic("bad param type")
		}
	}
	ok = true
	return c
}
