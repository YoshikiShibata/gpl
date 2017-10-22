// Copyright © 2016 Yoshiki Shibata

package cmplx

import (
	"fmt"
	"math"
	"math/big"
)

// FloatComplex represent an immutable complex
type FloatComplex struct {
	real_ *big.Float
	imag_ *big.Float
}

func (fc *FloatComplex) String() string {
	return fmt.Sprintf("{%s, %s}", fc.real_.String(), fc.imag_.String())
}

var precision uint = 0

func SetPrecision(prec uint) {
	precision = prec
}

func NewFloatComplex(r, i float64) *FloatComplex {
	var fc FloatComplex

	fc.real_ = big.NewFloat(r)
	fc.imag_ = big.NewFloat(i)
	if precision > 0 {
		fc.real_.SetPrec(precision)
		fc.imag_.SetPrec(precision)
	}
	return &fc
}

func (fc *FloatComplex) real() *big.Float {
	return fc.real_
}

func (fc *FloatComplex) imag() *big.Float {
	return fc.imag_
}

func (fc *FloatComplex) floats() (*big.Float, *big.Float) {
	return fc.real_, fc.imag_
}

func (fc *FloatComplex) IsZero() bool {
	r, _ := fc.real_.Float64()
	i, _ := fc.imag_.Float64()
	return r == 0.0 && i == 0.0
}

func (fc *FloatComplex) Mul(o *FloatComplex) *FloatComplex {
	// (a + bi)(c + di) = (ac - bd)+(bc +ad)i
	a, b := fc.floats()
	c, d := o.floats()

	r1 := new(big.Float)
	r1.Mul(a, c)
	r2 := new(big.Float)
	r2.Mul(b, d)
	r1 = r1.Sub(r1, r2)

	i1 := new(big.Float)
	i1.Mul(b, c)
	i2 := new(big.Float)
	i2.Mul(a, d)
	i1 = i1.Add(i1, i2)

	return &FloatComplex{r1, i1}
}

func (fc *FloatComplex) Add(o *FloatComplex) *FloatComplex {
	// (a + bi) + (c + di) = (a + c) + (b + d)i
	a, b := fc.floats()
	c, d := o.floats()

	r := new(big.Float)
	r.Add(a, c)
	i := new(big.Float)
	i.Add(b, d)

	return &FloatComplex{r, i}
}

func (fc *FloatComplex) Sub(o *FloatComplex) *FloatComplex {
	// (a + bi) - (c + di) = (a - c) + (b - d)i
	a, b := fc.floats()
	c, d := o.floats()

	r := new(big.Float)
	r.Sub(a, c)
	i := new(big.Float)
	i.Sub(b, d)

	return &FloatComplex{r, i}
}

func (fc *FloatComplex) Quo(o *FloatComplex) *FloatComplex {
	if fc.IsZero() && o.IsZero() {
		panic("Cannot Handle This")
	}

	if o.IsZero() {
		inf := math.Inf(0)
		return &FloatComplex{big.NewFloat(inf), big.NewFloat(inf)}
	}

	// (a + bi) / (c + di)
	//     = ((ac + bd)/(c*c + d*d)) + ((bc - ad)/(c*c + d*d))i
	a, b := fc.floats()
	c, d := o.floats()

	// r1 = (ac + ab)
	r1 := new(big.Float)
	r1.Mul(a, c)
	r2 := new(big.Float)
	r2.Mul(b, d)
	r1.Add(r1, r2)

	// r3 = (c*c + d*d)
	r3 := new(big.Float)
	r3.Mul(c, c)
	r4 := new(big.Float)
	r4.Mul(d, d)
	r3.Add(r3, r4)

	// r1 = ((ac + bd) / (c*c + d*d))
	r1.Quo(r1, r3)

	i1 := new(big.Float)
	i1.Mul(b, c)
	i2 := new(big.Float)
	i2.Mul(a, d)
	i1.Sub(i1, i2)

	i1.Quo(i1, r3)

	return &FloatComplex{r1, i1}
}

func (fc *FloatComplex) Abs() float64 {
	r, i := fc.floats()
	rFloat64, _ := r.Float64()
	iFloat64, _ := i.Float64()
	return math.Hypot(rFloat64, iFloat64)
}
