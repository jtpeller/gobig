// ============================================================================
// = int.go																	  =
// = 	Description: wrapper functions for big.Int							  =
// = 	Date: December 09, 2021												  =
// ============================================================================

package gobig

import (
	"math/big"
	"math/rand"
)

type bint = big.Int

// ##################### BIG.INT #####################
// ### INITIALIZATION SECTION

// Zero() returns a *bint with a value of 0
func Zero() *bint {
	r := big.NewInt(0)
	return r
}

// New() returns a *bint with a value of i
func New(i int64) *bint {
	r := big.NewInt(i)
	return r
}

// Copy() will copy the value of a and return it
func Copy(a *bint) *bint {
	return Add(Zero(), a)
}

// FromString() will convert provided string to a *bint
func FromString(s string, base int) *bint {
	n := Zero()
	n, _ = n.SetString(s, base)
	return n
}

// ### COMPARISONS

// Equals(a, b) returns true if a == b, false otherwise.
func Equals(a, b *bint) bool {
	return a.Cmp(b) == 0
}

// Less(a, b) returns true if a < b, false otherwise.
func Less(a, b *bint) bool {
	return a.Cmp(b) == -1
}

// Greater(a, b) returns true if a > b, false otherwise.
func Greater(a, b *bint) bool {
	return a.Cmp(b) == 1
}

// LessEqual(a, b) returns true if a <= b, false otherwise.
func LessEqual(a, b *bint) bool {
	return Less(a, b) || Equals(a, b)
}

// GreaterEqual(a, b) returns true if a a >= b, false otherwise.
func GreaterEqual(a, b *bint) bool {
	return Greater(a, b) || Equals(a, b)
}

// ### BITWISE OPERATIONS

// And(a, b) returns a & b
func And(a, b *bint) *bint {
	return Zero().And(a, b)
}

// AndNot(a, b) returns a &^ b
func AndNot(a, b *bint) *bint {
	return Zero().AndNot(a, b)
}

// Bit(a, i) returns the value of the ith bit of a
func Bit(a *bint, i int) uint {
	return Copy(a).Bit(i)
}

// BitLen(a) returns the length of a in bits
func BitLen(a *bint) int {
	return a.BitLen()
}

// Lsh(a, n) returns a << n
func Lsh(a *bint, n uint) *bint {
	return Zero().Lsh(a, uint(n))
}

// Not(a) returns ^a
func Not(a *bint) *bint {
	return Zero().Not(a)
}

// Or(a, b) returns a | b
func Or(a, b *bint) *bint {
	return Zero().Or(a, b)
}

// Rsh(a, n) returns a >> n
func Rsh(a *bint, n uint) *bint {
	return Zero().Rsh(a, n)
}

// Xor(a, b) returns a ^ b
func Xor(a, b *bint) *bint {
	return Zero().Xor(a, b)
}

// ### MATHEMATICAL OPERATIONS

// Abs(a) returns |a|, or the absolute value of a
func Abs(a *bint) *bint {
	return Zero().Abs(a)
}

// Add(a, b) returns a + b
func Add(a, b *bint) *bint {
	return Zero().Add(a, b)
}

// Sub(a, b) returns a - b
func Sub(a, b *bint) *bint {
	return Zero().Sub(a, b)
}

// Mul(a, b) returns a * b
func Mul(a, b *bint) *bint {
	return Zero().Mul(a, b)
}

// Div(a, b) returns a / b
func Div(a, b *bint) *bint {
	return Zero().Div(a, b)
}

// Mod(a, b) returns a % b
func Mod(a, b *bint) *bint {
	return Zero().Mod(a, b)
}

// Incr(a) returns a+1
func Incr(a *bint) *bint {
	return Add(a, New(1))
}

// Decr(a) returns a-1
func Decr(a *bint) *bint {
	return Sub(a, New(1))
}

// Prod(a, b) returns the product of all integers in the range [a, b]
func Prod(a, b int64) *bint {
	return Zero().MulRange(a, b)
}

// Exp(a, b, m) returns a**b mod |m|
func Exp(a, b, m *bint) *bint {
	return Zero().Exp(a, b, m)
}

// Pow(a, e) returns a^e
func Pow(a *bint, e *bint) *bint {
	return Exp(a, e, big.NewInt(0))
}

// Sqrt(a) returns the square root of a
func Sqrt(a *bint) *bint {
	return Zero().Sqrt(a)
}

// ### MISCELLANEOUS

// Binomial(n, k) returns the binomial coefficient of (n, k)
func Binomial(n, k int64) *bint {
	return Zero().Binomial(n, k)
}

// GCD(x, y, a, b) returns the GCD of a & b
// x & y (if not nil) will be set such that z = a*x + b*y (z is the GCD)
func GCD(x, y, a, b *bint) *bint {
	return Zero().GCD(x, y, a, b)
}

// Rand returns a pseudo-random number in the interval [0, n)
func Rand(rng *rand.Rand, n *bint) *bint {
	return Zero().Rand(rng, n)
}
