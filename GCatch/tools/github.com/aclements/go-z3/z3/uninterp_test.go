// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"testing"
)

func TestUninterpreted(t *testing.T) {
	ctx := NewContext(nil)
	u1, u2 := ctx.UninterpretedSort("u1"), ctx.UninterpretedSort("u2")
	a, b, c := ctx.Const("a", u1).(Uninterpreted), ctx.Const("b", u1).(Uninterpreted), ctx.Const("c", u1).(Uninterpreted)
	solver := NewSolver(ctx)
	solver.Assert(a.Eq(b))
	solver.Assert(a.NE(c))

	sat, err := solver.Check()
	if err != nil {
		t.Fatal(err)
	} else if !sat {
		t.Fatalf("want sat, got unsat for %v", solver)
	}
	m := solver.Model()

	sorts := m.Sorts()
	if len(sorts) != 1 {
		t.Fatalf("len(sorts) = %d, want 1", len(sorts))
	}
	if sorts[0].String() != "u1" {
		t.Fatalf("want sort u1, got %s", sorts[0])
	}
	u := m.SortUniverse(sorts[0])
	if len(u) != 2 {
		t.Fatalf("want universe of 2, got %s", u)
	}
	ma := m.Eval(a, true).AsAST()
	mb := m.Eval(b, true).AsAST()
	mc := m.Eval(c, true).AsAST()
	// ma and mb should be one of the universe's values and mc
	// should be the other.
	if !ma.Equal(u[0].AsAST()) {
		u[0], u[1] = u[1], u[0]
	}
	if !(ma.Equal(u[0].AsAST()) && mb.Equal(u[0].AsAST()) && mc.Equal(u[1].AsAST())) {
		t.Fatalf("want a=%s b=%s c=%s, got a=%s b=%s c=%s in universe %s", u[0], u[0], u[1], ma, mb, mc, u)
	}

	// If we ask for a sort that's not interpreted by m, we should
	// get an "invalid argument" panic.
	wantPanic(t, "invalid argument", func() { m.SortUniverse(u2) })
}
