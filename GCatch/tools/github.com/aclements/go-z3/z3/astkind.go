// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "strconv"

/*
#include <z3.h>
*/
import "C"

// ASTKind is a general category of ASTs, such as numerals,
// applications, or sorts.
type ASTKind int

const (
	ASTKindApp        = ASTKind(C.Z3_APP_AST)        // Constant and applications
	ASTKindNumeral    = ASTKind(C.Z3_NUMERAL_AST)    // Numeral constants (excluding real algebraic numbers)
	ASTKindVar        = ASTKind(C.Z3_VAR_AST)        // Bound variables
	ASTKindQuantifier = ASTKind(C.Z3_QUANTIFIER_AST) // Quantifiers
	ASTKindSort       = ASTKind(C.Z3_SORT_AST)       // Sorts
	ASTKindFuncDecl   = ASTKind(C.Z3_FUNC_DECL_AST)  // Function declarations
	ASTKindUnknown    = ASTKind(C.Z3_UNKNOWN_AST)
)

// String returns k as a string like "ASTKindApp".
func (k ASTKind) String() string {
	switch k {
	case ASTKindApp:
		return "ASTKindApp"
	case ASTKindNumeral:
		return "ASTKindNumeral"
	case ASTKindVar:
		return "ASTKindVar"
	case ASTKindQuantifier:
		return "ASTKindQuantifier"
	case ASTKindSort:
		return "ASTKindSort"
	case ASTKindFuncDecl:
		return "ASTKindFuncDecl"
	case ASTKindUnknown:
		return "ASTKindUnknown"
	}
	return "ASTKind(" + strconv.Itoa(int(k)) + ")"
}
