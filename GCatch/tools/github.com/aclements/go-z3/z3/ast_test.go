// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "testing"

func TestASTEquality(t *testing.T) {
	ctx := NewContext(nil)
	ints := ctx.IntSort()
	x := ctx.Const("x", ints).(Int)
	a := x.Add(ctx.FromInt(2, ints).(Int))
	b := x.Add(ctx.FromInt(1, ints).(Int)).Add(ctx.FromInt(1, ints).(Int))
	as, bs := ctx.Simplify(a, nil).AsAST(), ctx.Simplify(b, nil).AsAST()
	if !as.Equal(bs) {
		t.Errorf("%v != %v", as, bs)
	}
	if h1, h2 := as.Hash(), bs.Hash(); h1 != h2 {
		t.Errorf("hashes differ: %v != %v", h1, h2)
	}
	if i1, i2 := as.ID(), bs.ID(); i1 != i2 {
		t.Errorf("IDs differ: %v != %v", i1, i2)
	}
}

func TestASTAs(t *testing.T) {
	ctx := NewContext(nil)

	b1, b2 := ctx.BoolConst("b1"), ctx.FromBool(true)
	b3 := b1.Eq(b2)
	for _, val := range []Value{b1, b2, b3} {
		ast1 := val.AsAST()
		if !ast1.AsValue().AsAST().Equal(ast1) {
			t.Errorf("failed to round-trip value %v", val)
		}
		ast1.AsValue().(Bool).Not()
	}

	s1, s2 := ctx.IntSort(), ctx.BVSort(32)
	for _, sort := range []Sort{s1, s2} {
		ast1 := sort.AsAST()
		if !ast1.AsSort().AsAST().Equal(ast1) {
			t.Errorf("failed to round-trip sort %v", sort)
		}
		ctx.FreshConst("x", ast1.AsSort())
	}

	f1 := ctx.FuncDecl("f1", []Sort{s1}, s1)
	f2 := ctx.FuncDecl("f2", nil, s1)
	for _, fd := range []FuncDecl{f1, f2} {
		ast1 := fd.AsAST()
		if !ast1.AsFuncDecl().AsAST().Equal(ast1) {
			t.Errorf("failed to round-trip funcdecl %v", fd)
		}
	}
}

func TestASTTranslate(t *testing.T) {
	ctx1, ctx2 := NewContext(nil), NewContext(nil)

	x := ctx1.BoolConst("x")
	x.AsAST().Translate(ctx2).AsValue().(Bool).Eq(ctx2.FromBool(true))
}
