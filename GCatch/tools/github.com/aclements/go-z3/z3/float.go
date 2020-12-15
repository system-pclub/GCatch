// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z3

import (
	"math/big"
	"runtime"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

// Float is a symbolic value representing a floating-point number with
// IEEE 764-2008 semantics (extended to any possible exponent and
// significand length).
//
// A floating-point number falls into one of five categories: normal
// numbers, subnormal numbers, zero, infinite, or NaN. All of these
// except for NaN can be either positive or negative.
//
// Floating-point numbers are represented as three fields: a single
// bit sign, an exponent ebits wide, and a significand sbits-1 wide.
// If the exponent field is neither all 0s or all 1s, then the value
// is a "normal" number, which is interpreted as
//
//     (-1)^sign * 2^(exp - bias) * (1 + sig / 2^sbits)
//
// where bias = 2^(ebits - 1) - 1, and exp and sig are interpreted as
// unsigned binary values. In particular, the significand is extended
// with a "hidden" most-significant 1 bit and the exponent is
// "biased".
//
// Float implements Value.
type Float value

func init() {
	kindWrappers[KindFloatingPoint] = func(x value) Value {
		return Float(x)
	}
}

// FloatSort returns a floating-point sort with ebits exponent bits
// and sbits significand bits.
//
// Note that sbits counts the "hidden" most-significant bit of the
// significand. Hence, a double-precision floating point number has
// sbits of 53, even though only 52 bits are actually represented.
//
// Common exponent and significand bit counts are:
//
//                      ebits sbits
//     Half precision       5    11
//     Single precision     8    24  (float32)
//     Double precision    11    53  (float64)
//     Quad precision      15   113
func (ctx *Context) FloatSort(ebits, sbits int) Sort {
	var sort Sort
	ctx.do(func() {
		sort = wrapSort(ctx, C.Z3_mk_fpa_sort(ctx.c, C.uint(ebits), C.uint(sbits)), KindFloatingPoint)
	})
	return sort
}

// RoundingMode represents a floating-point rounding mode.
//
// The zero value of RoundingMode is RoundToNearestEven, which is the
// rounding mode in Go and the typical default rounding mode in
// floating-point math.
type RoundingMode int

const (
	// RoundToNearestEven rounds floating-point results to the
	// nearest representable value. If the result is exactly
	// midway between two representable values, the even
	// representable value is chosen.
	RoundToNearestEven RoundingMode = iota

	// RoundToNearestAway rounds floating-point results to the
	// nearest representable value. If the result is exactly
	// midway between two representable values, the value with the
	// largest magnitude (away from 0) is chosen.
	RoundToNearestAway

	// RoundToPositive rounds floating-point results toward
	// positive infinity.
	RoundToPositive

	// RoundToPositive rounds floating-point results toward
	// negative infinity.
	RoundToNegative

	// RoundToZero rounds floating-point results toward zero.
	RoundToZero

	roundingModesNum
)

var roundingModeKey = &[]string{"z3.roundingModeKey"}[0]

func (rm RoundingMode) ast(ctx *Context) value {
	// Ugh. Z3 represents these rounding modes as ASTs. Cache them
	// in the context.
	cache, _ := ctx.Extra(roundingModeKey).([]value)
	if cache == nil {
		cache = make([]value, roundingModesNum)
		cache[RoundToNearestEven] = wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_rne(ctx.c)
		})
		cache[RoundToNearestAway] = wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_rna(ctx.c)
		})
		cache[RoundToPositive] = wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_rtp(ctx.c)
		})
		cache[RoundToNegative] = wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_rtn(ctx.c)
		})
		cache[RoundToZero] = wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_rtz(ctx.c)
		})
		ctx.SetExtra(roundingModeKey, cache)
	}
	return cache[rm]
}

// rm returns ctx's current rounding mode, initializing it to
// RoundToNearestEven if it isn't set. The ctx lock must *not* be
// held.
func (ctx *Context) rm() value {
	ctx.lock.Lock()
	rm := ctx.roundingModeAST
	ctx.lock.Unlock()
	if rm.valueImpl == nil {
		// Lazily initialize the rounding mode.
		rm = ctx.roundingMode.ast(ctx)
		ctx.lock.Lock()
		ctx.roundingModeAST = rm
		ctx.lock.Unlock()
	}
	return rm
}

// SetRoundingMode sets ctx's current rounding mode for floating-point
// operations and returns ctx's old rounding mode.
func (ctx *Context) SetRoundingMode(rm RoundingMode) RoundingMode {
	rmv := rm.ast(ctx)
	ctx.lock.Lock()
	old := ctx.roundingMode
	ctx.roundingMode = rm
	ctx.roundingModeAST = rmv
	ctx.lock.Unlock()
	return old
}

// RoundingMode returns ctx's current rounding mode for floating-point
// operations.
func (ctx *Context) RoundingMode() RoundingMode {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	return ctx.roundingMode
}

// FloatNaN returns a floating-point NaN of sort s.
func (ctx *Context) FloatNaN(s Sort) Float {
	val := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_nan(ctx.c, s.c)
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(s)
	return val
}

// FloatInf returns a floating-point infinity of sort s.
func (ctx *Context) FloatInf(s Sort, neg bool) Float {
	val := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_inf(ctx.c, s.c, boolToZ3(neg))
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(s)
	return val
}

// FloatZero returns a (signed) floating-point zero of sort s.
func (ctx *Context) FloatZero(s Sort, neg bool) Float {
	val := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_zero(ctx.c, s.c, boolToZ3(neg))
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(s)
	return val
}

// FloatFromBits constructs a floating-point value from a sign bit,
// exponent bits, and significand bits.
//
// See the description of type Float for how these bits are
// interpreted.
//
// sign must be a 1-bit bit-vector.
func (ctx *Context) FloatFromBits(sign, exp, sig BV) Float {
	out := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_fp(ctx.c, sign.c, exp.c, sig.c)
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(sign)
	runtime.KeepAlive(exp)
	runtime.KeepAlive(sig)
	return out
}

// FromFloat32 constructs a floating-point literal from val. sort must
// be a floating-point sort.
func (ctx *Context) FromFloat32(val float32, sort Sort) Float {
	out := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_numeral_float(ctx.c, C.float(val), sort.c)
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(sort)
	return out
}

// FromFloat64 constructs a floating-point literal from val. sort must
// be a floating-point sort.
func (ctx *Context) FromFloat64(val float64, sort Sort) Float {
	out := Float(wrapValue(ctx, func() C.Z3_ast {
		return C.Z3_mk_fpa_numeral_double(ctx.c, C.double(val), sort.c)
	}))
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(sort)
	return out
}

var testingFloatAlwaysFromBigInt bool

func (ctx *Context) floatFromInt(val int64, sort Sort) Float {
	if val == 0 {
		return ctx.FloatZero(sort, false)
	}

	if uint64(val)>>62 != 0 || testingFloatAlwaysFromBigInt {
		// It's way too obnoxious to deal with overflow of
		// large numbers below (especially if we have round
		// val). Just fall back to the big.Int path.
		return ctx.floatFromBigInt(big.NewInt(val), sort)
	}

	ebits, sbits := sort.FloatSize()

	// Compute the sign bit.
	neg := false
	if val < 0 {
		neg, val = true, -val
	}

	// If val won't fit in sbits, round it.
	var lost uint // Bits dropped from val.
	if sbits < 64 && val >= 1<<uint(sbits) {
		for xval := val; xval >= 1<<uint(sbits); {
			xval, lost = xval>>1, lost+1
		}
		down := val >> lost
		// Is it still exact after we drop "lost" bits?
		if down<<lost == val {
			val = down
			goto exact
		}
		up := down + 1
		mid := (down << lost) + (1 << (lost - 1))
		switch ctx.RoundingMode() {
		case RoundToNearestEven:
			if val == mid {
				if down&1 == 0 {
					val = down
				} else {
					val = up
				}
			} else if val < mid {
				val = down
			} else {
				val = up
			}
		case RoundToNearestAway:
			// Despite the name, Z3 implements this as
			// "round to nearest odd". Bug in Z3?
			if val == mid {
				if down&1 == 0 {
					val = up
				} else {
					val = down
				}
			} else if val < mid {
				val = down
			} else {
				val = up
			}
		case RoundToPositive:
			if neg {
				val = down
			} else {
				val = up
			}
		case RoundToNegative:
			if neg {
				val = up
			} else {
				val = down
			}
		case RoundToZero:
			val = down
		}
	}
exact:

	// Compute the exponent. This is also the index of the
	// most-significant set bit.
	exp := uint(0)
	for xval := val; xval > 1; {
		xval, exp = xval>>1, exp+1
	}

	// Clear the hidden significand bit.
	val &^= 1 << exp

	if sbits <= 64 {
		// We can use the easy API. This expects an *unbiased*
		// exponent, but a shifted significand with the
		// most-significant bit stripped.
		out := Float(wrapValue(ctx, func() C.Z3_ast {
			return C.Z3_mk_fpa_numeral_int64_uint64(ctx.c, boolToZ3(neg), C.int64_t(exp+lost), C.uint64_t(val<<(uint(sbits)-exp-1)), sort.c)
		}))
		runtime.KeepAlive(ctx)
		return out
	}

	// We have to use the hard API. Build up the bit vectors
	// ourselves. In this case, the exponent is biased. The
	// significand bits behave the same as the easy case; they
	// just don't fit in a uint64.
	var sigBig big.Int
	sigBig.SetInt64(val)
	sigBig.Lsh(&sigBig, uint(sbits)-exp-1)
	exp += 1<<(uint(ebits)-1) - 1 + lost

	var bvSign BV
	if neg {
		bvSign = ctx.FromInt(1, ctx.BVSort(1)).(BV)
	} else {
		bvSign = ctx.FromInt(0, ctx.BVSort(1)).(BV)
	}
	bvExp := ctx.FromInt(int64(exp), ctx.BVSort(ebits)).(BV)
	bvSig := ctx.FromBigInt(&sigBig, ctx.BVSort(sbits-1)).(BV)
	return ctx.FloatFromBits(bvSign, bvExp, bvSig)
}

func (ctx *Context) floatFromBigInt(val *big.Int, sort Sort) Float {
	if val.Sign() == 0 {
		return ctx.FloatZero(sort, false)
	}
	ebits, sbits := sort.FloatSize()

	neg := val.Sign() < 0
	var x big.Int
	x.Abs(val)

	// If val won't fit in sbits, round it.
	var lost int
	if x.BitLen() > sbits {
		lost = x.BitLen() - sbits
		var up, down big.Int
		down.Rsh(&x, uint(lost))
		// Is it still exact after we drop "lost" bits?
		var tmp big.Int
		if tmp.Lsh(&down, uint(lost)).Cmp(&x) == 0 {
			x = down
			goto exact
		}
		up.Add(&down, big.NewInt(1))
		var mid big.Int
		mid.Lsh(&down, uint(lost)).SetBit(&mid, lost-1, 1)
		switch ctx.RoundingMode() {
		case RoundToNearestEven:
			switch x.Cmp(&mid) {
			case 0:
				if down.Bit(0) == 0 {
					x = down
				} else {
					x = up
				}
			case -1:
				x = down
			case 1:
				x = up
			}
		case RoundToNearestAway:
			switch x.Cmp(&mid) {
			case 0:
				if down.Bit(0) == 0 {
					x = up
				} else {
					x = down
				}
			case -1:
				x = down
			case 1:
				x = up
			}
		case RoundToPositive:
			if neg {
				x = down
			} else {
				x = up
			}
		case RoundToNegative:
			if neg {
				x = up
			} else {
				x = down
			}
		case RoundToZero:
			x = down
		}
	}
exact:

	// Compute the exponent. This is also the index of the
	// most-significant set bit.
	exp := x.BitLen() - 1

	// Clear the hidden significand bit.
	x.SetBit(&x, exp, 0)

	// Construct the significand bits.
	x.Lsh(&x, uint(sbits-exp-1))

	// Construct the biased exponent bits.
	exp += 1<<uint(ebits-1) - 1 + lost

	// Construct the bit-vector components.
	var bvSign BV
	if neg {
		bvSign = ctx.FromInt(1, ctx.BVSort(1)).(BV)
	} else {
		bvSign = ctx.FromInt(0, ctx.BVSort(1)).(BV)
	}
	bvExp := ctx.FromInt(int64(exp), ctx.BVSort(ebits)).(BV)
	bvSig := ctx.FromBigInt(&x, ctx.BVSort(sbits-1)).(BV)
	return ctx.FloatFromBits(bvSign, bvExp, bvSig)
}

// AsBigFloat returns the value of lit as a math/big.Float. If lit is
// not a literal, it returns nil, false. If lit is NaN, it returns
// nil, true (because big.Float cannot represent NaN).
func (lit Float) AsBigFloat() (val *big.Float, isLiteral bool) {
	// The Z3_fpa_get_numeral_* functions panic with "invalid
	// argument" if we pass them a non-literal, so we have to be
	// careful.
	_, sbits := lit.Sort().FloatSize()
	var out big.Float
	out.SetPrec(uint(sbits))
	switch {
	case lit.isAppOf(C.Z3_OP_FPA_NUM):
		var sign C.int
		var sig string
		var exp C.int64_t
		lit.ctx.do(func() {
			C.Z3_fpa_get_numeral_sign(lit.ctx.c, lit.c, &sign)
			sig = C.GoString(C.Z3_fpa_get_numeral_significand_string(lit.ctx.c, lit.c))
			//C.Z3_fpa_get_numeral_exponent_int64(lit.ctx.c, lit.c, &exp, C.Z3_FALSE)
			C.Z3_fpa_get_numeral_exponent_int64(lit.ctx.c, lit.c, &exp, false)
		})
		out.Parse(sig, 10)
		if sign > 0 {
			out.Neg(&out)
		}
		out.SetMantExp(&out, int(exp))
	case lit.isAppOf(C.Z3_OP_FPA_PLUS_ZERO):
	case lit.isAppOf(C.Z3_OP_FPA_MINUS_ZERO):
		out.Parse("-0", 10)
	case lit.isAppOf(C.Z3_OP_FPA_PLUS_INF):
		out.Parse("+inf", 10)
	case lit.isAppOf(C.Z3_OP_FPA_MINUS_INF):
		out.Parse("-inf", 10)
	case lit.isAppOf(C.Z3_OP_FPA_NAN):
		return nil, true
	default:
		return nil, false
	}
	return &out, true
}

//go:generate go run genwrap.go -t Float $GOFILE

// Abs returns the absolute value of l.
//
//wrap:expr Abs Z3_mk_fpa_abs l

// Neg returns -l.
//
//wrap:expr Neg Z3_mk_fpa_neg l

// Add returns l+r.
//
// Add uses the current rounding mode.
//
//wrap:expr Add Z3_mk_fpa_add @rm l r

// Sub returns l-r.
//
// Sub uses the current rounding mode.
//
//wrap:expr Sub Z3_mk_fpa_sub @rm l r

// Mul returns l*r.
//
// Mul uses the current rounding mode.
//
//wrap:expr Mul Z3_mk_fpa_mul @rm l r

// Div returns l/r.
//
// Div uses the current rounding mode.
//
//wrap:expr Div Z3_mk_fpa_div @rm l r

// MulAdd returns l*r+a (fused multiply and add).
//
// MulAdd uses the current rounding mode on the result of the whole
// operation.
//
//wrap:expr MulAdd Z3_mk_fpa_fma @rm l r a

// Sqrt returns the square root of l.
//
// Sqrt uses the current rounding mode.
//
//wrap:expr Sqrt Z3_mk_fpa_sqrt @rm l

// Rem returns the remainder of l/r.
//
//wrap:expr Rem Z3_mk_fpa_rem l r

// Round rounds l to an integral floating-point value according to
// rounding mode rm.
//
//wrap:expr Round l rm:RoundingMode : Z3_mk_fpa_round_to_integral rm l

// Min returns the minimum of l and r.
//
//wrap:expr Min Z3_mk_fpa_min l r

// Max returns the maximum of l and r.
//
//wrap:expr Max Z3_mk_fpa_max l r

// IEEEEq returns l == r according to IEEE 754 equality.
//
// This differs from Eq, which is true if l and r are identical. In
// contrast, under IEEE equality, ±0 == ±0, while NaN != NaN and ±inf
// != ±inf.
//
//wrap:expr IEEEEq:Bool Z3_mk_fpa_eq l r

// LT returns l < r.
//
//wrap:expr LT:Bool Z3_mk_fpa_lt l r

// LE returns l <= r.
//
//wrap:expr LE:Bool Z3_mk_fpa_leq l r

// GT returns l > r.
//
//wrap:expr GT:Bool Z3_mk_fpa_gt l r

// GE returns l >= r.
//
//wrap:expr GE:Bool Z3_mk_fpa_geq l r

// IsNormal returns true if l is a normal floating-point number.
//
//wrap:expr IsNormal:Bool Z3_mk_fpa_is_normal l

// IsSubnormal returns true if l is a subnormal floating-point number.
//
//wrap:expr IsSubnormal:Bool Z3_mk_fpa_is_subnormal l

// IsZero returns true if l is ±0.
//
//wrap:expr IsZero:Bool Z3_mk_fpa_is_zero l

// IsInfinite returns true if l is ±∞.
//
//wrap:expr IsInfinite:Bool Z3_mk_fpa_is_infinite l

// IsNaN returns true if l is NaN.
//
//wrap:expr IsNaN:Bool Z3_mk_fpa_is_nan l

// IsNegative returns true if l is negative.
//
//wrap:expr IsNegative:Bool Z3_mk_fpa_is_negative l

// IsPositive returns true if l is positive.
//
//wrap:expr IsPositive:Bool Z3_mk_fpa_is_positive l

// ToFloat converts l into a floating-point number of a different
// floating-point sort.
//
// If necessary, the result will be rounded according to the current
// rounding mode.
//
//wrap:expr ToFloat l s:Sort : Z3_mk_fpa_to_fp_float @rm l s

// ToUBV converts l.Round() into an unsigned bit-vector of size 'bits'.
//
// l is first rounded to an integer using the current rounding mode.
// If the result is not in the range [0, 2^bits-1], the result is
// unspecified.
//
//wrap:expr ToUBV:BV l bits:int : Z3_mk_fpa_to_ubv @rm l bits:unsigned

// ToSBV converts l.Round() into a signed bit-vector of size 'bits'.
//
// l is first rounded to an integer using the current rounding mode.
// If the result is not in the range [-2^(bits-1), 2^(bits-1)-1], the
// result is unspecified.
//
//wrap:expr ToSBV:BV l bits:int : Z3_mk_fpa_to_sbv @rm l bits:unsigned

// ToReal converts l into a real number.
//
// If l is ±inf, or NaN, the result is unspecified.
//
//wrap:expr ToReal:Real Z3_mk_fpa_to_real l

// ToIEEEBV converts l to a bit-vector in IEEE 754-2008 format.
//
// Note that NaN has many possible representations. This conversion
// always uses the same representation.
//
//wrap:expr ToIEEEBV:BV Z3_mk_fpa_to_ieee_bv l
