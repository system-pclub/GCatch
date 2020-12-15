// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package st provides symbolic equivalents of basic Go types.
//
// Every value from this package can be either concrete or symbolic.
// Operations on concrete values always produce concrete values and
// are implemented as the underlying Go operations, so they're
// efficient. If one of the operands is symbolic, the result is
// generally symbolic.
//
// Every type provided by this package has two fields: C and S. If the
// value is concrete, C is its concrete value. Otherwise, S is its
// symbolic value. This makes it convenient to write concrete
// literals. E.g., st.Int{C: 2} is a concrete "2".
//
// Every type also has an "Any" constructor that returns an
// unconstrained symbolic value of that type. This value can be
// thought of as taking on every possible value of that type.
//
// Every Go operation on a type has a corresponding method. Where
// possible, these follow the names from math/big.
//
//	x + y	x.Add(y)
//	x - y	x.Sub(y)
//	x * y	x.Mul(y)
//	x / y	x.Div(y)	(does not panic on symbolic divide by 0)
//	x % y	x.Rem(y)
//
//	x & y	x.And(y)
//	x | y	x.Or(y)
//	x ^ y	x.Xor(y)
//	x << y	x.Lsh(y)
//	x >> y	x.Rsh(y)
//	x &^ y	x.AndNot(y)
//
//	x && y	x.And(y)	(Bool only; *not* short-circuiting)
//	x || y	x.Or(y)		(Bool only; *not* short-circuiting)
//
//	x == y	x.Eq(y)
//	x != y	x.NE(y)
//	x <  y	x.LT(y)
//	x <= y	x.LE(y)
//	x >  y	x.GT(y)
//	x >= y	x.GE(y)
//
//	-x	x.Neg()
//	^x	x.Not()
//	!x	x.Not() 	(Bool only)
//
// For any pair of types T and U that support conversion in Go, T has
// a method ToU() that returns a U value.
//
// TODO: Float, complex, and string types.
package st

// RealApproxDigits is the number of decimal digits an irrational real
// will be approximated to when evaluating it as a *big.Rat.
var RealApproxDigits = 100
