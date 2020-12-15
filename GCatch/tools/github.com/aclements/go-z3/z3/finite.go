// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// FiniteDomain is a symbolic value from a finite domain of size n,
// where n depends on the value's sort.
//
// Finite domain values are uninterpreted (see Uninterpreted), but
// represented as numerals. Hence, to construct a specific
// finite-domain value, use methods like Context.FromInt with a value
// in [0, n).
//
// FiniteDomain implements Value.
type FiniteDomain value

func init() {
	kindWrappers[KindFiniteDomain] = func(x value) Value {
		return FiniteDomain(x)
	}
}

// FiniteDomainSort returns a sort for finite domain values with
// domain size n.
//
// Two finite-domain sorts are the same if and only if they have the
// same name.
func (ctx *Context) FiniteDomainSort(name string, n uint64) Sort {
	sym := ctx.symbol(name)
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_finite_domain_sort(ctx.c, sym, C.uint64_t(n)), KindFiniteDomain)
	})
	return sort
}

//go:generate go run genwrap.go -t FiniteDomain $GOFILE
