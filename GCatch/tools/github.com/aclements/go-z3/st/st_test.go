// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package st

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"

	"git.gradebot.org/zxl381/goconcurrencychecker/tools/github.com/aclements/go-z3/z3"
)

// Test that concrete and symbolic operations are equivalent.

func TestEquivBool(t *testing.T) {
	testEquiv(t, reflect.TypeOf(Bool{}), Bool.sym, true, false)
}

func TestEquivInt32(t *testing.T) {
	testEquiv(t, reflect.TypeOf(Int32{}), Int32.sym,
		int32(0), int32(1), int32(-1), int32(31), int32(32),
		int32(math.MaxInt32), int32(math.MinInt32))
}

func TestEquivUint32(t *testing.T) {
	testEquiv(t, reflect.TypeOf(Uint32{}), Uint32.sym,
		uint32(0), uint32(1), uint32(31), uint32(32), uint32(1<<32-1))
}

func TestEquivInteger(t *testing.T) {
	var huge big.Int
	huge.SetString("123456789012345678901234567890", 10)
	testEquiv(t, reflect.TypeOf(Integer{}), Integer.sym,
		big.NewInt(0), big.NewInt(1), big.NewInt(-1), big.NewInt(-3),
		big.NewInt(32), &huge)
}

func TestEquivReal(t *testing.T) {
	testEquiv(t, reflect.TypeOf(Real{}), Real.sym,
		big.NewRat(0, 1), big.NewRat(1, 1), big.NewRat(-1, 1),
		big.NewRat(-3, 1), big.NewRat(-1, 3))
}

func testEquiv(t *testing.T, typ reflect.Type, symMethod interface{}, vals ...interface{}) {
	ctx := z3.NewContext(nil)

	rvals := make([]reflect.Value, len(vals))
	for i, val := range vals {
		rvals[i] = reflect.ValueOf(val)
	}

	for mi := 0; mi < typ.NumMethod(); mi++ {
		m := typ.Method(mi)
		switch m.Name {
		case "IsConcrete", "Eval", "String":
			continue
		case "Lsh", "Rsh":
			// TODO: Test these
			continue
		}
		t.Run(m.Name, func(t *testing.T) {
			inputs := genArgs(rvals, m.Type.NumIn())
			for _, input := range inputs {
				if m.Name == "Quo" || m.Name == "Rem" {
					s := fmt.Sprint(input[1].Interface())
					if s == "0" || s == "0/1" {
						// Avoid divide by zero
						continue
					}
				}

				c, s := wrap(ctx, typ, symMethod, input)

				// Do the operation concretely.
				cres := m.Func.Call(c)[0]

				// Do the operation symbolically.
				sres := m.Func.Call(s)[0]

				// Check that they're equal.
				eq := cres.MethodByName("Eq").Call([]reflect.Value{sres})[0]
				if !toBool(ctx, eq.Interface().(Bool)) {
					t.Errorf("%s(%v) = %v, want %v", m.Name, sliceInterface(c), ctx.Simplify(sres.FieldByName("S").Interface().(z3.Value), nil), cres.FieldByName("C").Interface())
				}
			}
		})
	}
}

// genArgs returns the Cartesian product vals^n.
func genArgs(vals []reflect.Value, n int) [][]reflect.Value {
	if n == 0 {
		return [][]reflect.Value{nil}
	}
	next := genArgs(vals, n-1)
	res := [][]reflect.Value{}
	for _, base := range next {
		for _, val := range vals {
			res = append(res, append([]reflect.Value{val}, base...))
		}
	}
	return res
}

// wrap constructs st values of type typ with the concrete and
// symbolic values given in vals.
func wrap(ctx *z3.Context, typ reflect.Type, symMethod interface{}, vals []reflect.Value) (con, sym []reflect.Value) {
	rcache := reflect.ValueOf(getCache(ctx))
	rsym := reflect.ValueOf(symMethod)
	for _, val := range vals {
		c := reflect.New(typ).Elem()
		c.FieldByName("C").Set(val)
		s := reflect.New(typ).Elem()
		s.FieldByName("S").Set(rsym.Call([]reflect.Value{c, rcache})[0])
		con = append(con, c)
		sym = append(sym, s)
	}
	return
}

func sliceInterface(rs []reflect.Value) []interface{} {
	out := make([]interface{}, len(rs))
	for i, r := range rs {
		out[i] = r.Interface()
	}
	return out
}

func toBool(ctx *z3.Context, b Bool) bool {
	// Since everything is literals, the simplifier should have no
	// trouble getting the answer and is dramatically faster than
	// the solver.
	val, ok := ctx.Simplify(b.S, nil).(z3.Bool).AsBool()
	if !ok {
		panic("failed to simplify to a literal")
	}
	return val
}
