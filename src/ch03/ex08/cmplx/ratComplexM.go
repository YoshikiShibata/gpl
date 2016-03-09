package cmplx

import (
	"math"
	"math/big"
)

// RatComplexM supports a minimum big.Rat operation for Mandelbrot
type RatComplexM struct {
	real_ *big.Rat
	imag_ *big.Rat
}

func NewRatComplexM(r, i float64) *RatComplexM {
	var rcm RatComplexM

	rcm.real_ = &big.Rat{}
	rcm.imag_ = &big.Rat{}
	rcm.real_.SetFloat64(r)
	rcm.imag_.SetFloat64(i)

	return &rcm
}

func (rcm *RatComplexM) Mul(o *RatComplexM) *RatComplexM {
	// (a + bi)(c + di) = (ac - bd)+(bc + ad)i

	a, b := rcm.real_, rcm.imag_
	c, d := o.real_, o.imag_

	r1 := &big.Rat{}
	r1.Mul(a, c) // ac
	r2 := &big.Rat{}
	r2.Mul(b, d) // bd
	r := &big.Rat{}
	r.Sub(r1, r2)

	r1.Mul(b, c) // bc
	r2.Mul(a, d) // ad
	i := &big.Rat{}
	i.Add(r1, r2)

	return &RatComplexM{r, i}
}

func (rcm *RatComplexM) Add(o *RatComplexM) *RatComplexM {
	// (a + bi) + (c + di) = (a + c) + (b + d)i
	a, b := rcm.real_, rcm.imag_
	c, d := o.real_, o.imag_

	r := &big.Rat{}
	r.Add(a, c) // a + c

	i := &big.Rat{}
	i.Add(b, d) // b + d

	return &RatComplexM{r, i}
}

func (rcm *RatComplexM) Abs() float64 {
	r, _ := rcm.real_.Float64()
	i, _ := rcm.imag_.Float64()

	return math.Hypot(r, i)
}
