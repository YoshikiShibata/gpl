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

func (rf *RatFloat) String() string {
	if rf.v != nil {
		return rf.String()
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
