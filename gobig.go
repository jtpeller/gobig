// ============================================================================
// = bignum.go																  =
// = 	Description: VERY helpful wrapper functions for big.Int				  =
// = 	Date: December 09, 2021												  =
// ============================================================================

package gobig

import "math/big"

// INITS
func Zero() *big.Int {
	r := big.NewInt(0)
	return r
}

func New(i int64) *big.Int {
	r := big.NewInt(i)
	return r
}

func Copy(a *big.Int) *big.Int {
	return Add(Zero(), a)
}

// COMPARISONS
func Equals(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func Less(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

func Greater(a, b *big.Int) bool {
	return a.Cmp(b) == 1
}

func LessEqual(a, b *big.Int) bool {
	return Less(a, b) || Equals(a, b)
}

func GreaterEqual(a, b *big.Int) bool {
	return Greater(a, b) || Equals(a, b)
}

// MATHEMATICAL OPERATIONS
func Abs(a *big.Int) *big.Int {
	return Zero().Abs(a)
}

func Add(a, b *big.Int) *big.Int {
	return Zero().Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return Zero().Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return Zero().Mul(a,b)
}

func Div(a, b *big.Int) *big.Int {
	return Zero().Div(a, b)
}

func Mod(a, b *big.Int) *big.Int {
	return Zero().Mod(a, b)
}

func Incr(a *big.Int) *big.Int {
	return Add(a, New(1))
}

func Decr(a *big.Int) *big.Int {
	return Sub(a, New(1))
}

func Pow(a *big.Int, e *big.Int) *big.Int {
	return Zero().Exp(a, e, big.NewInt(0))
}

func Sqrt(a *big.Int) *big.Int {
	return Zero().Sqrt(a)
}

func Lsh(a *big.Int, e int) *big.Int {
	return Zero().Lsh(a, uint(e))
}
