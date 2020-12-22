// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"math"
	"math/big"
	"testing"
)

func TestRealRational(t *testing.T) {
	ctx := NewContext(nil)
	rat := ctx.FromBigRat(big.NewRat(5, 4))
	numer, denom, isLit := rat.AsRat()
	if !isLit {
		t.Errorf("(%s).AsRat() returned false", rat)
	} else {
		val, isLit, ok := numer.AsInt64()
		if !(val == 5 && isLit && ok) {
			t.Errorf("numerator of %s: wanted 5, true, true; got %v, %v, %v", rat, val, isLit, ok)
		}
		val, isLit, ok = denom.AsInt64()
		if !(val == 4 && isLit && ok) {
			t.Errorf("numerator of %s: wanted 4, true, true; got %v, %v, %v", rat, val, isLit, ok)
		}
	}

	_, _, isLit = rat.Approx(10)
	if isLit {
		// rat is a rational, so Approx should fail.
		t.Errorf("(%s).Approx(10) returned true", rat)
	}
}

func TestRealIrrational(t *testing.T) {
	ctx := NewContext(nil)
	root2 := ctx.Simplify(ctx.FromInt(2, ctx.IntSort()).(Int).ToReal().Exp(ctx.FromBigRat(big.NewRat(1, 2))), nil).(Real)

	_, _, isLit := root2.AsRat()
	if isLit {
		t.Errorf("(%s).AsRat() returned true", root2)
	}

	l, u, isLit := root2.Approx(10)
	if !isLit {
		t.Errorf("(%s).Approx(10) returned false", root2)
	} else {
		t.Logf("(%s).Approx(10) = [%v, %v]", root2, l, u)
		lr, isLit := l.AsBigRat()
		if !isLit {
			t.Fatalf("lower bound %v is not a literal rational", l)
		}
		ur, isLit := u.AsBigRat()
		if !isLit {
			t.Fatalf("upper bound %v is not a literal rational", u)
		}
		const r2 = 1.4142135623730951
		lf, _ := lr.Float64()
		if math.Abs(lf-r2) > 1e-10 {
			t.Errorf("lower bound |%v - %v| > 1e-10", lf, r2)
		}
		uf, _ := ur.Float64()
		if math.Abs(uf-r2) > 1e-10 {
			t.Errorf("upper bound |%v - %v| > 1e-10", uf, r2)
		}
	}
}
