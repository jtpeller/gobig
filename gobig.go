// ============================================================================
// = bignum.go																  =
// = 	Description: VERY helpful wrapper functions for big.Int				  =
// = 	Date: December 09, 2021												  =
// ============================================================================

package gobig

import "math/big"

const Version = "v0.1.1"

// ##################### BIG.INT #####################
// ### INITIALIZATION SECTION

// Zero() returns a *big.Int with a value of 0
func Zero() *big.Int {
	r := big.NewInt(0)
	return r
}

// New() returns a *big.Int with a value of i
func New(i int64) *big.Int {
	r := big.NewInt(i)
	return r
}

// Copy() will copy the value of a and return it
func Copy(a *big.Int) *big.Int {
	return Add(Zero(), a)
}

// ### COMPARISON SECTIOn

// Equals(a, b) returns true if a == b, false otherwise.
func Equals(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

// Less(a, b) returns true if a < b, false otherwise.
func Less(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

// Greater(a, b) returns true if a > b, false otherwise.
func Greater(a, b *big.Int) bool {
	return a.Cmp(b) == 1
}

// LessEqual(a, b) returns true if a <= b, false otherwise.
func LessEqual(a, b *big.Int) bool {
	return Less(a, b) || Equals(a, b)
}

// GreaterEqual(a, b) returns true if a a >= b, false otherwise.
func GreaterEqual(a, b *big.Int) bool {
	return Greater(a, b) || Equals(a, b)
}

// ### MATHEMATICAL OPERATIONS

// Abs(a) returns |a|, or the absolute value of a
func Abs(a *big.Int) *big.Int {
	return Zero().Abs(a)
}

// Add(a, b) returns a + b
func Add(a, b *big.Int) *big.Int {
	return Zero().Add(a, b)
}

// Sub(a, b) returns a - b
func Sub(a, b *big.Int) *big.Int {
	return Zero().Sub(a, b)
}

// Mul(a, b) returns a * b
func Mul(a, b *big.Int) *big.Int {
	return Zero().Mul(a,b)
}

// Div(a, b) returns a / b
func Div(a, b *big.Int) *big.Int {
	return Zero().Div(a, b)
}

// Mod(a, b) returns a % b
func Mod(a, b *big.Int) *big.Int {
	return Zero().Mod(a, b)
}

// Incr(a) returns a+1
func Incr(a *big.Int) *big.Int {
	return Add(a, New(1))
}

// Decr(a) returns a-1
func Decr(a *big.Int) *big.Int {
	return Sub(a, New(1))
}

// Pow(a, e) returns a^e
func Pow(a *big.Int, e *big.Int) *big.Int {
	return Zero().Exp(a, e, big.NewInt(0))
}

// Sqrt(a) returns the square root of a
func Sqrt(a *big.Int) *big.Int {
	return Zero().Sqrt(a)
}

// Lsh(a, e) returns the result from shifting a left by e
func Lsh(a *big.Int, e int) *big.Int {
	return Zero().Lsh(a, uint(e))
}

// ##################### BIG.FLOAT ##################### 
// ### INITIALIZATION SECTION

// ZeroFloat() returns a new *big.Float of value 0
func ZeroFloat() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

// NewFloat(a) returns a new *big.Float of value a
func NewFloat(a float64) *big.Float {
	r := big.NewFloat(a)
	r.SetPrec(256)
	return r
}

// ### COMPARISON SECTION

// Equals(a, b) returns true if a == b, false otherwise.
func EqualsFloat(a, b *big.Float) bool {
	return a.Cmp(b) == 0
}

// Less(a, b) returns true if a < b, false otherwise.
func LessFloat(a, b *big.Float) bool {
	return a.Cmp(b) == -1
}

// Greater(a, b) returns true if a > b, false otherwise.
func GreaterFloat(a, b *big.Float) bool {
	return a.Cmp(b) == 1
}

// LessEqual(a, b) returns true if a <= b, false otherwise.
func LessEqualFloat(a, b *big.Float) bool {
	return LessFloat(a, b) || EqualsFloat(a, b)
}

// GreaterEqual(a, b) returns true if a a >= b, false otherwise.
func GreaterEqualFloat(a, b *big.Float) bool {
	return GreaterFloat(a, b) || EqualsFloat(a, b)
}


// ### MATHEMATICAL OPERATIONS

// AbsFloat(a) returns |a| or the absolute value of a
func AbsFloat(a *big.Float) *big.Float {
	return ZeroFloat().Abs(a)
}

// AddFloat(a, b) returns a + b
func AddFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Add(a, b)
}

// SubFloat(a, b) returns a - b
func SubFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Sub(a, b)
}

// MulFloat(a, b) returns a * b
func MulFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Mul(a,b)
}

// DivFloat(a, b) returns a / b
func DivFloat(a, b *big.Float) *big.Float {
	return ZeroFloat().Quo(a, b)
}

// PowFloat(a, e) returns a^e
func PowFloat(a *big.Float, e int64) *big.Float {
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
func SqrtFloat(a *big.Float) *big.Float {
	return ZeroFloat().Sqrt(a)
}

// Floor(a) returns the floor of a as a *big.Int
func Floor(a *big.Float) *big.Int {
	f, _ := a.Int(nil)
	return f
}

// ToFloat(a) converts a (of type *big.Int) to type *big.Float
func ToFloat(a *big.Int) *big.Float {
	r := new(big.Float)
	r.SetInt(a)
	r.SetPrec(256)
	return r
}

// Round(a) rounds a, and returns it as a *big.Int
func Round(a *big.Float) *big.Int {
	a.Add(a, NewFloat(0.5))
	r, _ := a.Int(nil)
	return r
}