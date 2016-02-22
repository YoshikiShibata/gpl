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
