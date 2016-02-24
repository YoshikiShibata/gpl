// Copyright © 2016 Yoshiki Shibata

package cmplx

import (
	"fmt"
	"math"
	"math/big"
)

type RatFloat struct {
	v *big.Rat // nil represents the value is not finite
	f float64  // used to hold Infinite and NaN
}

func NewRatFloat(f float64) *RatFloat {
	var rf RatFloat
	rf.f = f
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return &rf
	}
	rf.v = &big.Rat{}
	r := rf.v.SetFloat64(f)
	if r == nil {
		panic("Unexpected Nil")
	}
	return &rf
}

func (rf *RatFloat) IsInf() bool {
	return math.IsInf(rf.f, 0)
}

func (rf *RatFloat) IsNaN() bool {
	return math.IsNaN(rf.f)
}

func (rc *RatFloat) Mul(o *RatFloat) *RatFloat {
	if rc.IsNaN() || o.IsNaN() ||
		rc.IsInf() || o.IsInf() {
		return NewRatFloat(rc.f * o.f)
	}

	var result big.Rat
	result.Mul(rc.v, o.v)
	f, _ := result.Float64()
	return &RatFloat{&result, f}
}

func (rc *RatFloat) Quo(o *RatFloat) *RatFloat {
	if rc.IsNaN() || o.IsNaN() ||
		rc.IsInf() || o.IsInf() {
		return NewRatFloat(rc.f / o.f)
	}

	if o.f == 0.0 {
		return NewRatFloat(rc.f / o.f)
	}

	var result big.Rat
	result.Quo(rc.v, o.v)
	f, _ := result.Float64()
	return &RatFloat{&result, f}
}

func (rc *RatFloat) Add(o *RatFloat) *RatFloat {
	if rc.IsNaN() || o.IsNaN() ||
		rc.IsInf() || o.IsInf() {
		return NewRatFloat(rc.f + o.f)
	}

	var result big.Rat
	result.Add(rc.v, o.v)
	f, _ := result.Float64()
	return &RatFloat{&result, f}
}

func (rc *RatFloat) Sub(o *RatFloat) *RatFloat {
	if rc.IsNaN() || o.IsNaN() ||
		rc.IsInf() || o.IsInf() {
		return NewRatFloat(rc.f - o.f)
	}

	var result big.Rat
	result.Sub(rc.v, o.v)
	f, _ := result.Float64()
	return &RatFloat{&result, f}
}

func (rf *RatFloat) String() string {
	if rf.v != nil {
		return rf.v.RatString()
	}
	return fmt.Sprintf("%v", rf.f)
}

func (rf *RatFloat) Float64() float64 {
	if rf.v == nil {
		return rf.f
	}
	f, _ := rf.v.Float64()
	return f
}

type RatComplex struct {
	real_ *RatFloat
	imag_ *RatFloat
}

func (rc *RatComplex) String() string {

	return fmt.Sprintf("{%s, %s}", rc.real_.String(), rc.imag_.String())
}

func NewRatComplex(r, i float64) *RatComplex {
	var rc RatComplex

	rc.real_ = NewRatFloat(r)
	rc.imag_ = NewRatFloat(i)
	return &rc
}

func (rc *RatComplex) real() *RatFloat {
	return rc.real_
}

func (rc *RatComplex) imag() *RatFloat {
	return rc.imag_
}

func (rc *RatComplex) floats() (*RatFloat, *RatFloat) {
	return rc.real_, rc.imag_
}

func (rc *RatComplex) IsZero() bool {
	r := rc.real_.Float64()
	i := rc.imag_.Float64()

	return r == 0.0 && i == 0.0
}

func (rc *RatComplex) Mul(o *RatComplex) *RatComplex {
	// (a + bi)(c + di) = (ac - bd)+(bc + ad)i

	a, b := rc.floats()
	c, d := o.floats()

	r1 := a.Mul(c) // ac
	r2 := b.Mul(d) // bd
	r := r1.Sub(r2)

	i1 := b.Mul(c) // bc
	i2 := a.Mul(d) // ad
	i := i1.Add(i2)

	return &RatComplex{r, i}
}

func (rc *RatComplex) Add(o *RatComplex) *RatComplex {
	// (a + bi) + (c + di) = (a + c) + (b + d)i
	a, b := rc.floats()
	c, d := o.floats()

	r := a.Add(c)
	i := b.Add(d)

	return &RatComplex{r, i}
}

func (rc *RatComplex) Sub(o *RatComplex) *RatComplex {
	// (a + bi) - (c + di) = (a - c) + (b - d)i
	a, b := rc.floats()
	c, d := o.floats()

	r := a.Sub(c)
	i := b.Sub(d)

	return &RatComplex{r, i}
}

func (rc *RatComplex) Quo(o *RatComplex) *RatComplex {
	if rc.IsZero() && o.IsZero() {
		panic("Cannot Handle This")
	}

	if o.IsZero() {
		inf := math.Inf(0)
		return &RatComplex{NewRatFloat(inf), NewRatFloat(inf)}
	}

	// (a + bi) / (c + di)
	//     = ((ac + bd)/(c*c + d*d)) + ((bc -ad)/(c*c + d*d))i
	a, b := rc.floats()
	c, d := o.floats()

	denominator := c.Mul(c).Add(d.Mul(d))
	r := a.Mul(c).Add(b.Mul(d)).Quo(denominator)
	i := b.Mul(c).Sub(a.Mul(d)).Quo(denominator)

	return &RatComplex{r, i}
}

func (rc *RatComplex) Abs() float64 {
	r, i := rc.floats()

	return math.Hypot(r.f, i.f)
}
