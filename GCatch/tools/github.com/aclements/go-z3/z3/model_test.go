// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "testing"

func TestModel(t *testing.T) {
	// Create a simple formula with a unique solution.
	ctx := NewContext(nil)
	s := NewSolver(ctx)
	x, y := ctx.BoolConst("x"), ctx.BoolConst("y")
	s.Assert(x.And(y.Not()))

	sat, err := s.Check()
	if err != nil {
		t.Fatalf("failed to compute satisfiability: %s", err)
	} else if !sat {
		t.Fatalf("formula not satisfiable")
	}

	m := s.Model()
	x, y = m.Eval(x, false).(Bool), m.Eval(y, false).(Bool)
	xval, xok := x.AsBool()
	yval, yok := y.AsBool()
	if !(xval && xok && !yval && yok) {
		t.Fatalf("expected x -> true, y -> false; got\n%s", m)
	}
}
