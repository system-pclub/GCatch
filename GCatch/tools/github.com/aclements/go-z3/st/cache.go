// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package st

import "github.com/system-pclub/GCatch/GCatch/tools/github.com/aclements/go-z3/z3"

//go:generate go run gen.go -o types.go

type cache struct {
	z3 *z3.Context

	sorts
}

type cacheKeyType struct{}

var cacheKey interface{} = (*cacheKeyType)(nil)

func getCache(ctx *z3.Context) *cache {
	c, _ := ctx.Extra(cacheKey).(*cache)
	if c == nil {
		c = &cache{z3: ctx}
		initSorts(&c.sorts, ctx)
		ctx.SetExtra(cacheKey, c)
	}
	return c
}
