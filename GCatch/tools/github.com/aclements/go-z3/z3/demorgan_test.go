// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "testing"

func TestDeMorgan(t *testing.T) {
	// Test proving De Morgan's duality law.
	ctx := NewContext(nil)

	x, y := ctx.BoolConst("x"), ctx.BoolConst("y")
	l := x.And(y).Not()
	r := x.Not().Or(y.Not())
	conjecture := l.Iff(r)

	s := NewSolver(ctx)
	s.Assert(conjecture.Not())
	sat, err := s.Check()
	if err != nil {
		t.Errorf("failed to compute satisfiability: %s", err)
	} else if sat {
		t.Errorf("disproved De Morgan's law")
	}
}
