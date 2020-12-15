// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestFloatFromInt(t *testing.T) {
	ctx := NewContext(nil)
	isort := ctx.IntSort()
	for _, sbits := range []int{10, 32, 53, 64, 100} {
		s := ctx.FloatSort(11, sbits)
		for _, x := range []int64{0, 1, -1, 2, -2, 42, -42,
			math.MaxInt32, -math.MaxInt32,
			math.MaxInt64, math.MinInt64,
			math.MaxInt64 - 1, math.MinInt64 + 1,
		} {
			got := ctx.FromInt(x, s).(Float)
			// Create the "want" value by converting the
			// infinite-precision int via Z3.
			want := ctx.FromInt(x, isort).(Int).ToReal().ToFloat(s)
			want = ctx.Simplify(want, nil).(Float)
			if !simplifyBool(t, ctx, got.Eq(want)) {
				t.Errorf("FromInt(%d [%d]) = %s, want %s", x, sbits, got, want)
			}
		}
	}
}

func TestFloatFromBigInt(t *testing.T) {
	t.Skip()
	testingFloatAlwaysFromBigInt = true
	defer func() { testingFloatAlwaysFromBigInt = false }()
	TestFloatFromInt(t)
}

func TestFloatFromIntRounding(t *testing.T) {
	t.Skip()
	// Test that our implementation of rounding of integers to fit
	// in too-small floats matches Z3's.
	ctx := NewContext(nil)
	isort := ctx.IntSort()
	s := ctx.FloatSort(10, 4)
	for _, sign := range []int64{1, -1} {
		for x := int64(15); x < 64; x++ {
			x := x * sign
			for rm := RoundingMode(0); rm < roundingModesNum; rm++ {
				ctx.SetRoundingMode(rm)
				got, _, _ := ctx.Simplify(ctx.FromInt(x, s).(Float).ToReal().ToInt(), nil).(Int).AsInt64()

				// Round it using Z3.
				y := ctx.FromInt(x, isort).(Int).ToReal().ToFloat(s).ToReal().ToInt()
				y = ctx.Simplify(y, nil).(Int)
				want, _, _ := y.AsInt64()

				if got != want {
					t.Errorf("%d as float[10, 4] with rounding mode %v: want %d, got %d", x, rm, want, got)
				}
			}
		}
	}
}

func TestFloatFromBigIntRounding(t *testing.T) {
	testingFloatAlwaysFromBigInt = true
	defer func() { testingFloatAlwaysFromBigInt = false }()
	TestFloatFromIntRounding(t)
}

func TestFloatRound(t *testing.T) {
	ctx := NewContext(nil)
	s := ctx.FloatSort(11, 53)
	type test struct {
		val float64
		res [5]int64
	}
	for _, test := range []test{
		{2.5, [5]int64{2, 3, 3, 2, 2}},
		{-2.5, [5]int64{-2, -3, -2, -3, -2}},
		{2.8, [5]int64{3, 3, 3, 2, 2}},
		{-2.8, [5]int64{-3, -3, -2, -3, -2}},
	} {
		f := ctx.FromFloat64(test.val, s)
		for rm, want := range test.res {
			fr := f.Round(RoundingMode(rm))
			if !simplifyBool(t, ctx, fr.Eq(ctx.FromInt(want, s).(Float))) {
				fmt.Println("Round(A, B) = C, want D. A: ",test.val, " B:", RoundingMode(rm), " C:", fr, " D:", want)
				t.Fatal()
			}
		}
	}
}

func TestFloatAsBigFloat(t *testing.T) {
	ctx := NewContext(nil)
	s := ctx.FloatSort(11, 53)
	for _, test := range []float64{0, math.Copysign(0, -1), 42, -42,
		math.Inf(1), math.Inf(-1),
		math.MaxFloat64, -math.MaxFloat64,
		math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64,
		math.NaN()} {
		f := ctx.FromFloat64(test, s)
		fv, ok := f.AsBigFloat()
		if !ok {
			t.Errorf("%s is not a literal", f)
		} else if math.IsNaN(test) {
			if fv != nil {
				t.Errorf("want %v.AsBigFloat() == nil, got %v", test, fv)
			}
		} else if fv == nil {
			t.Errorf("want %v.AsBigFloat() == %v, got nil", test, test)
		} else if v, acc := fv.Float64(); v != test || acc != big.Exact {
			t.Errorf("want %v.AsBigFloat() == %v, got %v [%v]", test, test, v, acc)
		} else if fv.Prec() != 53 {
			t.Errorf("want fv.Prec() == 53, got %v", fv.Prec())
		}
	}
}
