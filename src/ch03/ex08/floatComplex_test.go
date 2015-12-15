package main

import "testing"

func TestNew(t *testing.T) {
	fc := NewFloatComplex(0.0, 0.0)

	img_ := fc.imag()
	real_ := fc.real()

	f64, _ := real_.Float64()
	if f64 != 0.0 {
		t.Errorf("Result is %g, want 0.0", f64)
	}
	f64, _ = img_.Float64()
	if f64 != 0.0 {
		t.Errorf("Result is %g, want 0.0", f64)
	}
}

func TestMul(t *testing.T) {
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
	} {
		var fc *FloatComplex
		var cplx complex128

		fc1 := NewFloatComplex(test.r1, test.i1)
		fc2 := NewFloatComplex(test.r2, test.i2)
		cplx1 := complex(test.r1, test.i1)
		cplx2 := complex(test.r2, test.i2)

		switch test.op {
		case '*':
			fc = fc1.Mul(fc2)
			cplx = cplx1 * cplx2
		case '+':
			fc = fc1.Add(fc2)
			cplx = cplx1 + cplx2
		case '-':
			fc = fc1.Sub(fc2)
			cplx = cplx1 - cplx2
		case '/':
			fc = fc1.Quo(fc2)
			cplx = cplx1 / cplx2
		default:
			t.Fatalf("Undefined op = %v", test.op)
		}

		verify(t, fc, cplx)
	}
}

func verify(t *testing.T, fc *FloatComplex, cplx complex128) {
	fcReal, _ := fc.real().Float64()
	fcImag, _ := fc.imag().Float64()

	if fcReal != real(cplx) {
		t.Errorf("real is %g, want %g", fcReal, real(cplx))
	}

	if fcImag != imag(cplx) {
		t.Errorf("img is %g, want %g", fcImag, imag(cplx))
	}
}
