// ============================================================================
// = rat.go																	  =
// = 	Description: wrapper functions for big.Rat							  =
// = 	Date: December 13, 2021												  =
// ============================================================================

package gobig

import "math/big"

type rat = big.Rat

// ##################### BIG.RAT #####################
// ### INITIALIZATION SECTION

// ZeroRat() returns the ratio of 0 / 1
func ZeroRat() *rat {
	return big.NewRat(0, 1)
}

// NewRat(num, den) returns the ratio of num / den
func NewRat(num, den int64) *rat {
	return big.NewRat(num, den)
}

// ### COMPARISONS

// EqualsRat(a, b) returns true if a == b, false otherwise.
func EqualsRat(a, b *rat) bool {
	return a.Cmp(b) == 0
}

// Less(a, b) returns true if a < b, false otherwise.
func LessRat(a, b *rat) bool {
	return a.Cmp(b) == -1
}

// Greater(a, b) returns true if a > b, false otherwise.
func GreaterRat(a, b *rat) bool {
	return a.Cmp(b) == 1
}

// LessEqual(a, b) returns true if a <= b, false otherwise.
func LessEqualRat(a, b *rat) bool {
	return LessRat(a, b) || EqualsRat(a, b)
}

// GreaterEqual(a, b) returns true if a a >= b, false otherwise.
func GreaterEqualRat(a, b *rat) bool {
	return GreaterRat(a, b) || EqualsRat(a, b)
}

// ### MATHEMATICAL OPERATIONS

// AbsRat(a) returns the absolute value of a
func AbsRat(a *rat) *rat {
	return ZeroRat().Abs(a)
}

// AddRat(a, b) returns a + b
func AddRat(a, b *rat) *rat {
	return ZeroRat().Add(a, b)
}

// DivRat(a, b) returns a / b
func DivRat(a, b *rat) *rat {
	return ZeroRat().Quo(a, b)
}

// Inv(a) returns 1/a
func Inv(a *rat) *rat {
	return ZeroRat().Inv(a)
}

// Mul(a, b) returns a * b
func MulRat(a, b *rat) *rat {
	return ZeroRat().Mul(a, b)
}

// Sub(a, b) returns a - b
func SubRat(a, b *rat) *rat {
	return ZeroRat().Sub(a, b)
}
