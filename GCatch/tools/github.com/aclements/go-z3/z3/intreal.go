// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

// Methods that are common to both Int and Real. This file is passed
// to genwrap.go twice with different default types.

// Add returns the sum l + r[0] + r[1] + ...
//
//wrap:expr Add Z3_mk_add l r...

// Mul returns the product l * r[0] * r[1] * ...
//
//wrap:expr Mul Z3_mk_mul l r...

// Sub returns l - r[0] - r[1] - ...
//
//wrap:expr Sub Z3_mk_sub l r...

// Neg returns -l.
//
//wrap:expr Neg Z3_mk_unary_minus l

// Exp returns lᶠ.
//
//wrap:expr Exp Z3_mk_power l r

// LT returns l < r.
//
//wrap:expr LT:Bool Z3_mk_lt l r

// LE returns l <= r.
//
//wrap:expr LE:Bool Z3_mk_le l r

// GT returns l > r.
//
//wrap:expr GT:Bool Z3_mk_gt l r

// GE returns l >= r.
//
//wrap:expr GE:Bool Z3_mk_ge l r
