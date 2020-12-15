// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ops

import "go/token"

type Type struct {
	// StName is the name of the st package wrapper.
	StName string

	// ConType is the concrete Go type for this type.
	ConType string

	// SymType is the Z3 symbolic type for this type.
	SymType string

	Flags Flags

	Bits int
}

type Flags int

const (
	// Kind flags.
	IsBool Flags = 1 << iota
	IsInteger
	IsUnsigned
	IsFloat
	IsComplex
	IsString

	IsBigInt
	IsBigRat

	// Operator type flags.
	OpShift   // Shift operator; right operand is a Uint64
	OpPos     // Unary + operator (identity); no method
	OpCompare // Comparison operator; result is a bool

	// Misc operator flags.
	Z3SignedPrefix // Prefix Z3 method with S/U.

	Comparable = IsBool | IsInteger | IsFloat | IsComplex | IsString | IsBigInt | IsBigRat
	Ordered    = IsInteger | IsFloat | IsString | IsBigInt | IsBigRat
)

var Types = []Type{
	{"Bool", "bool", "Bool", IsBool, 0},
	{"Int", "int", "BV", IsInteger, intBits()},
	{"Int8", "int8", "BV", IsInteger, 8},
	{"Int16", "int16", "BV", IsInteger, 16},
	{"Int32", "int32", "BV", IsInteger, 32},
	{"Int64", "int64", "BV", IsInteger, 64},
	{"Uint", "uint", "BV", IsInteger | IsUnsigned, intBits()},
	{"Uint8", "uint8", "BV", IsInteger | IsUnsigned, 8},
	{"Uint16", "uint16", "BV", IsInteger | IsUnsigned, 16},
	{"Uint32", "uint32", "BV", IsInteger | IsUnsigned, 32},
	{"Uint64", "uint64", "BV", IsInteger | IsUnsigned, 64},
	{"Uintptr", "uintptr", "BV", IsInteger | IsUnsigned, ptrBits()},

	{"Integer", "*big.Int", "Int", IsBigInt, 0},
	{"Real", "*big.Rat", "Real", IsBigRat, 0},
}

func intBits() int {
	n := 0
	for x := ^uint(0); x != 0; x >>= 1 {
		n++
	}
	return n
}

func ptrBits() int {
	n := 0
	for x := ^uintptr(0); x != 0; x >>= 1 {
		n++
	}
	return n
}

type Op struct {
	// Op is the Go syntax for this operation.
	Op string

	// Tok is the token package name for this operation.
	Tok token.Token

	// Method is the method name for this operation.
	//
	// This is also generally the Z3 method on the underlying Z3
	// value, with possible tweaks based on Flags.
	Method string

	// Flags specifies the set of types this operation applies to
	// and may contain flags giving the type of the operator.
	Flags Flags
}

var BinOps = []Op{
	{"+", token.ADD, "Add", IsInteger | IsFloat | IsString | IsBigInt | IsBigRat},
	{"-", token.SUB, "Sub", IsInteger | IsFloat | IsBigInt | IsBigRat},
	{"*", token.MUL, "Mul", IsInteger | IsFloat | IsBigInt | IsBigRat},
	{"/", token.QUO, "Quo", IsInteger | IsFloat | IsBigInt | IsBigRat | Z3SignedPrefix},
	{"%", token.REM, "Rem", IsInteger | IsBigInt | Z3SignedPrefix},

	{"&", token.AND, "And", IsInteger},
	{"|", token.OR, "Or", IsInteger},
	{"^", token.XOR, "Xor", IsInteger},
	{"<<", token.SHL, "Lsh", OpShift | IsInteger},
	{">>", token.SHR, "Rsh", OpShift | IsInteger | Z3SignedPrefix},
	{"&^", token.AND_NOT, "AndNot", IsInteger},

	{"&&", token.LAND, "And", IsBool},
	{"||", token.LOR, "Or", IsBool},

	{"==", token.EQL, "Eq", OpCompare | Comparable},
	{"!=", token.NEQ, "NE", OpCompare | Comparable},
	{"<", token.LSS, "LT", OpCompare | Ordered | Z3SignedPrefix},
	{"<=", token.LEQ, "LE", OpCompare | Ordered | Z3SignedPrefix},
	{">", token.GTR, "GT", OpCompare | Ordered | Z3SignedPrefix},
	{">=", token.GEQ, "GE", OpCompare | Ordered | Z3SignedPrefix},
}

var UnOps = []Op{
	{"+", token.ADD, "", IsInteger | IsFloat | IsBigInt | IsBigRat | OpPos},
	{"-", token.SUB, "Neg", IsInteger | IsFloat | IsBigInt | IsBigRat},
	{"^", token.XOR, "Not", IsInteger},
	{"!", token.NOT, "Not", IsBool},
}
