// ============================================================================
// = float.go																  =
// = 	Description: wrapper functions for big.Float						  =
// = 	Date: December 13, 2021												  =
// ============================================================================

package gobig

import "math/big"

type bfloat = big.Float

// ##################### BIG.FLOAT #####################
// ### INITIALIZATION SECTION

// ZeroFloat() returns a new *bfloat of value 0
func ZeroFloat() *bfloat {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

// NewFloat(a) returns a new *bfloat of value a
func NewFloat(a float64) *bfloat {
	r := big.NewFloat(a)
	r.SetPrec(256)
	return r
}

// CopyFloat(a) returns a new *bfloat with the value of a
func CopyFloat(a *bfloat) *bfloat {
	return ZeroFloat().Copy(a)
}

// ### COMPARISON SECTION

// Equals(a, b) returns true if a == b, false otherwise.
func EqualsFloat(a, b *bfloat) bool {
	return a.Cmp(b) == 0
}

// Less(a, b) returns true if a < b, false otherwise.
func LessFloat(a, b *bfloat) bool {
	return a.Cmp(b) == -1
}

// Greater(a, b) returns true if a > b, false otherwise.
func GreaterFloat(a, b *bfloat) bool {
	return a.Cmp(b) == 1
}

// LessEqual(a, b) returns true if a <= b, false otherwise.
func LessEqualFloat(a, b *bfloat) bool {
	return LessFloat(a, b) || EqualsFloat(a, b)
}

// GreaterEqual(a, b) returns true if a a >= b, false otherwise.
func GreaterEqualFloat(a, b *bfloat) bool {
	return GreaterFloat(a, b) || EqualsFloat(a, b)
}

// ### MATHEMATICAL OPERATIONS

// AbsFloat(a) returns |a| or the absolute value of a
func AbsFloat(a *bfloat) *bfloat {
	return ZeroFloat().Abs(a)
}

// AddFloat(a, b) returns a + b
func AddFloat(a, b *bfloat) *bfloat {
	return ZeroFloat().Add(a, b)
}

// SubFloat(a, b) returns a - b
func SubFloat(a, b *bfloat) *bfloat {
	return ZeroFloat().Sub(a, b)
}

// MulFloat(a, b) returns a * b
func MulFloat(a, b *bfloat) *bfloat {
	return ZeroFloat().Mul(a,b)
}

// DivFloat(a, b) returns a / b
func DivFloat(a, b *bfloat) *bfloat {
	return ZeroFloat().Quo(a, b)
}

// PowFloat(a, e) returns a^e
func PowFloat(a *bfloat, e int64) *bfloat {
	if e == 0 {
		return NewFloat(1)
	}
	r := ZeroFloat().Copy(a)
	for i := int64(0); i < e-1; i++ {
		r = MulFloat(r, a)
	}
	if e < 0 {		// negative exponents are reciprocals of positives
		return DivFloat(NewFloat(1), r)
	}
	return r
}

// SqrtFloat(a) returns the square root of a
func SqrtFloat(a *bfloat) *bfloat {
	return ZeroFloat().Sqrt(a)
}

// Floor(a) returns the floor of a as a *bint
func Floor(a *bfloat) *bint {
	f, _ := a.Int(nil)
	return f
}

// ToFloat(a) converts a (of type *bint) to type *bfloat
func ToFloat(a *bint) *bfloat {
	r := new(big.Float)
	r.SetInt(a)
	r.SetPrec(256)
	return r
}

// Round(a) rounds a, and returns it as a *bint
func Round(a *bfloat) *bint {
	a.Add(a, NewFloat(0.5))
	r, _ := a.Int(nil)
	return r
}