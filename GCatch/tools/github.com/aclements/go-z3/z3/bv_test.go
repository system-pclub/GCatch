// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"math/big"
	"testing"
)

func TestBVAsBig(t *testing.T) {
	ctx := NewContext(nil)

	val := ctx.FromBigInt(big.NewInt(255), ctx.BVSort(8)).(BV)
	big1, _ := val.AsBigUnsigned()
	if big1.String() != "255" {
		t.Errorf("expected 255, got %s", big1)
	}
	big2, _ := val.AsBigSigned()
	if big2.String() != "-1" {
		t.Errorf("expected -1, got %s", big2)
	}

	val = ctx.FromBigInt(big.NewInt(-1), ctx.BVSort(8)).(BV)
	big1, _ = val.AsBigUnsigned()
	if big1.String() != "255" {
		t.Errorf("expected 255, got %s", big1)
	}
	big2, _ = val.AsBigSigned()
	if big2.String() != "-1" {
		t.Errorf("expected -1, got %s", big2)
	}
}

func TestBVAsInt64(t *testing.T) {
	ctx := NewContext(nil)

	x := ctx.FromInt(255, ctx.BVSort(8)).(BV)
	vu, isConst, ok := x.AsUint64()
	if vu != 255 || !isConst || !ok {
		t.Errorf("255:8 as uint: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vu, isConst, ok)
	}
	vs, isConst, ok := x.AsInt64()
	if vs != -1 || !isConst || !ok {
		t.Errorf("255:8 as int: expected %v, %v, %v; got %v, %v, %v", -1, true, true, vs, isConst, ok)
	}

	x = ctx.FromInt(-1, ctx.BVSort(8)).(BV)
	vu, isConst, ok = x.AsUint64()
	if vu != 255 || !isConst || !ok {
		t.Errorf("-1:8 as uint: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vu, isConst, ok)
	}
	vs, isConst, ok = x.AsInt64()
	if vs != -1 || !isConst || !ok {
		t.Errorf("-1:8 as int: expected %v, %v, %v; got %v, %v, %v", -1, true, true, vs, isConst, ok)
	}

	x = ctx.FromInt(255, ctx.BVSort(9)).(BV)
	vu, isConst, ok = x.AsUint64()
	if vu != 255 || !isConst || !ok {
		t.Errorf("255:9 as uint: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vu, isConst, ok)
	}
	vs, isConst, ok = x.AsInt64()
	if vs != 255 || !isConst || !ok {
		t.Errorf("255:9 as int: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vs, isConst, ok)
	}

	x = ctx.FromInt(255, ctx.BVSort(64)).(BV)
	vu, isConst, ok = x.AsUint64()
	if vu != 255 || !isConst || !ok {
		t.Errorf("255:64 as uint: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vu, isConst, ok)
	}
	vs, isConst, ok = x.AsInt64()
	if vs != 255 || !isConst || !ok {
		t.Errorf("255:64 as int: expected %v, %v, %v; got %v, %v, %v", 255, true, true, vs, isConst, ok)
	}

	x = ctx.FromInt(-1, ctx.BVSort(128)).(BV)
	vu, isConst, ok = x.AsUint64()
	if vu != 0 || !isConst || ok {
		t.Errorf("-1:128 as uint: expected %v, %v, %v; got %v, %v, %v", 0, true, false, vu, isConst, ok)
	}
	vs, isConst, ok = x.AsInt64()
	if vs != -1 || !isConst || !ok {
		t.Errorf("-1:128 as int: expected %v, %v, %v; got %v, %v, %v", -1, true, true, vs, isConst, ok)
	}
}
