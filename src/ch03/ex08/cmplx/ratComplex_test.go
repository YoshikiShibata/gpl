// Copyright © 2016 Yoshiki Shibata

package cmplx

import (
	"math"
	"testing"
)

func TestInfinity(t *testing.T) {
	for _, test := range []struct {
		value    float64
		expected bool
	}{
		{0.0, false},
		{math.NaN(), false},
		{math.Inf(1), true},
		{math.Inf(-1), true},
	} {
		rf := NewRatFloat(test.value)
		if rf.IsInf() != test.expected {
			t.Fatalf("%v returned, but want %v\n", rf.IsInf(), test.expected)
		}
	}
}
func TestNaN(t *testing.T) {
	for _, test := range []struct {
		value    float64
		expected bool
	}{
		{0.0, false},
		{math.NaN(), true},
		{math.Inf(1), false},
		{math.Inf(-1), false},
	} {
		rf := NewRatFloat(test.value)
		if rf.IsNaN() != test.expected {
			t.Fatalf("%v returned, but want %v\n", rf.IsNaN(), test.expected)
		}
	}
}

func TestMulRat(t *testing.T) {
	for _, test := range []struct {
		r1, i1 float64
		op     rune
		r2, i2 float64
	}{
		{0.0, 0.0, '*', 0.0, 0.0},
		{0.0, 0.0, '+', 0.0, 0.0},
		{0.0, 0.0, '-', 0.0, 0.0},
		{1.0, 1.0, '*', 2.0, 2.0},
		{1.0, 1.0, '+', 2.0, 2.0},
		{1.0, 1.0, '-', 2.0, 2.0},
		{1.0, 1.0, '/', 2.0, 2.0},
		{128.0, 128.0, '*', 256.0, 256.0},
		{128.0, 128.0, '+', 256.0, 256.0},
		{128.0, 128.0, '-', 256.0, 256.0},
		{128.0, 128.0, '/', 256.0, 256.0},
		{1.0, 0.0, '/', 0.0, 0.0},
	} {
		var rc *RatComplex
		var cplx complex128

		rc1 := NewRatComplex(test.r1, test.i1)
		rc2 := NewRatComplex(test.r2, test.i2)
		cplx1 := complex(test.r1, test.i1)
		cplx2 := complex(test.r2, test.i2)

		switch test.op {
		case '*':
			rc = rc1.Mul(rc2)
			cplx = cplx1 * cplx2
		case '+':
			rc = rc1.Add(rc2)
			cplx = cplx1 + cplx2
		case '-':
			rc = rc1.Sub(rc2)
			cplx = cplx1 - cplx2
		case '/':
			rc = rc1.Quo(rc2)
			cplx = cplx1 / cplx2
		default:
			t.Fatalf("Undefined op = %v", test.op)
		}

		verifyRatComplex(t, rc, cplx)
	}
}

func verifyRatComplex(t *testing.T, rc *RatComplex, cplx complex128) {
	rcReal := rc.real().Float64()
	rcImag := rc.imag().Float64()

	if rcReal != real(cplx) {
		t.Errorf("real is %g, want %g", rcReal, real(cplx))
	}

	if rcImag != imag(cplx) {
		t.Errorf("img is %g, want %g", rcImag, imag(cplx))
	}
}
