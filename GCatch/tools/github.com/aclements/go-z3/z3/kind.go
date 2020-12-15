// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import "strconv"

/*
#include <z3.h>
*/
import "C"

// kindWrappers is a map of Value constructors for each sort kind.
var kindWrappers = make(map[Kind]func(x value) Value)

// Kind is a general category of sorts, such as int or array.
type Kind int

const (
	KindUninterpreted = Kind(C.Z3_UNINTERPRETED_SORT)
	KindBool          = Kind(C.Z3_BOOL_SORT)
	KindInt           = Kind(C.Z3_INT_SORT)
	KindReal          = Kind(C.Z3_REAL_SORT)
	KindBV            = Kind(C.Z3_BV_SORT)
	KindArray         = Kind(C.Z3_ARRAY_SORT)
	KindDatatype      = Kind(C.Z3_DATATYPE_SORT)
	KindRelation      = Kind(C.Z3_RELATION_SORT)
	KindFiniteDomain  = Kind(C.Z3_FINITE_DOMAIN_SORT)
	KindFloatingPoint = Kind(C.Z3_FLOATING_POINT_SORT)
	KindRoundingMode  = Kind(C.Z3_ROUNDING_MODE_SORT)
	KindUnknown       = Kind(C.Z3_UNKNOWN_SORT)
)

// String returns k as a string like "KindBool".
func (k Kind) String() string {
	switch k {
	case KindUninterpreted:
		return "KindUninterpreted"
	case KindBool:
		return "KindBool"
	case KindInt:
		return "KindInt"
	case KindReal:
		return "KindReal"
	case KindBV:
		return "KindBV"
	case KindArray:
		return "KindArray"
	case KindDatatype:
		return "KindDatatype"
	case KindRelation:
		return "KindRelation"
	case KindFiniteDomain:
		return "KindFiniteDomain"
	case KindFloatingPoint:
		return "KindFloatingPoint"
	case KindRoundingMode:
		return "KindRoundingMode"
	case KindUnknown:
		return "KindUnknown"
	}
	return "Kind(" + strconv.Itoa(int(k)) + ")"
}
