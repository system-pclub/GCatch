// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package z3 checks the satisfiability of logical formulas.
//
// This package provides bindings for the Z3 SMT solver
// (https://github.com/Z3Prover/z3). Z3 checks satisfiability of
// logical formulas over a wide range of terms, including booleans,
// integers, reals, bit-vectors, and uninterpreted functions. For a
// good introduction to the concepts of SMT and Z3, see the Z3 guide
// (http://rise4fun.com/z3/tutorialcontent/guide).
//
// This package does not yet support all of the features or types
// supported by Z3, though it supports a reasonably large subset.
//
// The main entry point to the z3 package is type Context. All values
// are created and all solving is done relative to some Context, and
// values from different Contexts cannot be mixed.
//
// Symbolic values implement the Value interface. Every value has a
// type, called a "sort" and represented by type Sort. Sorts fall into
// general categories called "kinds", such as Bool and Int. Each kind
// corresponds to a different concrete type that implements the Value
// interface, since the kind determines the set of operations that
// make sense on a value. A Bool expression is also called a
// "formula".
//
// These concrete value types help with type checking expressions, but
// type checking is ultimately done dynamically by Z3. Attempting to
// create a badly typed value will panic.
//
// Symbolic values are represented as expressions of numerals,
// constants, and uninterpreted functions. A numeral is a literal,
// fixed value like "2". A constant is a term like "x", whose value is
// fixed but unspecified. An uninterpreted function is a function
// whose mapping from arguments to results is fixed but unspecified
// (this is in contrast to an "interpreted function" like + whose
// interpretation is specified to be addition). Functions are pure
// (side-effect-free) like mathematical functions, but unlike
// mathematical functions they are always total. A constant can be
// thought of as a function with zero arguments.
//
// It's possible to go back and forth between a symbolic value and the
// expression representing that value using Value.AsAST and
// AST.AsValue.
//
// Type Solver checks the satisfiability of a set of formulas. If the
// Solver determines that a set of formulas is satisfiable, it can
// construct a Model giving a specific assignment of constants and
// uninterpreted functions that satisfies the set of formulas.
package z3

/*
#include <z3.h>
*/
import "C"

func boolToZ3(b bool) C.Z3_bool {
	if b {
		//return C.Z3_TRUE
		return true
	}
	//return C.Z3_FALSE
	return false
}

func z3ToBool(b C.Z3_bool) bool {
	//return b != C.Z3_FALSE
	return b != false
}
