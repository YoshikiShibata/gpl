// Copyright © 2016 Yoshiki Shibata

package cmplx

import (
	"fmt"
	"math"
	"math/big"
)

type RatFloat struct {
	v    *big.Rat // nil represents the value is not finite
	f    float64  // used to hold Infinite and NaN
	fSet bool     // f field is set or not
}

func NewRatFloat(f float64) *RatFloat {
	var rf RatFloat
	rf.f = f
	rf.fSet = true
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
	if !rf.fSet {
		return false
	}
	return math.IsInf(rf.f, 0)
}

func (rf *RatFloat) IsNaN() bool {
	if !rf.fSet {
		return false
	}
	return math.IsNaN(rf.f)
}

func (rf *RatFloat) Mul(o *RatFloat) *RatFloat {
	if rf.IsNaN() || o.IsNaN() ||
		rf.IsInf() || o.IsInf() {
		return NewRatFloat(rf.Float64() * o.Float64())
	}

	var result big.Rat
	result.Mul(rf.v, o.v)
	// f, _ := result.Float64()
	return &RatFloat{&result, 0.0, false}
}

func (rf *RatFloat) Quo(o *RatFloat) *RatFloat {
	if rf.IsNaN() || o.IsNaN() ||
		rf.IsInf() || o.IsInf() {
		return NewRatFloat(rf.Float64() / o.Float64())
	}

	if o.Float64() == 0.0 {
		return NewRatFloat(rf.Float64() / o.Float64())
	}

	var result big.Rat
	result.Quo(rf.v, o.v)
	// f, _ := result.Float64()
	return &RatFloat{&result, 0.0, false}
}

func (rf *RatFloat) Add(o *RatFloat) *RatFloat {
	if rf.IsNaN() || o.IsNaN() ||
		rf.IsInf() || o.IsInf() {
		return NewRatFloat(rf.Float64() + o.Float64())
	}

	var result big.Rat
	result.Add(rf.v, o.v)
	// f, _ := result.Float64()
	return &RatFloat{&result, 0.0, false}
}

func (rf *RatFloat) Sub(o *RatFloat) *RatFloat {
	if rf.IsNaN() || o.IsNaN() ||
		rf.IsInf() || o.IsInf() {
		return NewRatFloat(rf.Float64() - o.Float64())
	}

	var result big.Rat
	result.Sub(rf.v, o.v)
	// f, _ := result.Float64()
	return &RatFloat{&result, 0.0, false}
}

func (rf *RatFloat) String() string {
	if rf.v != nil {
		return rf.v.RatString()
	}
	return fmt.Sprintf("%v", rf.Float64())
}

func (rf *RatFloat) Float64() float64 {
	if rf.fSet {
		return rf.f
	}
	// fmt.Printf("\t\t\t\tRat to float64\n")
	rf.f, _ = rf.v.Float64()
	rf.fSet = true
	return rf.f
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
		if rc.real_.Float64() != 0.0 {
			return &RatComplex{NewRatFloat(math.Inf(0)),
				NewRatFloat(math.NaN())}
		}
		return &RatComplex{NewRatFloat(math.NaN()),
			NewRatFloat(math.Inf(0))}
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

	return math.Hypot(r.Float64(), i.Float64())
}
