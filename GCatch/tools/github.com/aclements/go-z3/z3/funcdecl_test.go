// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "testing"

func TestFuncDecl(t *testing.T) {
	ctx := NewContext(nil)
	s := NewSolver(ctx)
	ints := ctx.IntSort()
	fn := ctx.FuncDecl("f", []Sort{ints}, ints)

	s.Assert(fn.Apply(ctx.FromInt(1, ints).(Int)).(Int).Eq(ctx.FromInt(2, ints).(Int)))
	if sat, err := s.Check(); !sat {
		t.Errorf("%s not satisfiable: %s", s, err)
	}

	s.Assert(fn.Apply(ctx.FromInt(1, ints).(Int)).(Int).Eq(ctx.FromInt(3, ints).(Int)))
	if sat, err := s.Check(); sat {
		t.Errorf("%s satisfiable: %s", s, err)
	}
}
