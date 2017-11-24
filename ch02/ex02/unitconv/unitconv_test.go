// Copyright © 2015 Yoshiki Shibata. All rights reserved.

package unitconv_test

import (
	"fmt"
	"testing"

	"github.com/YoshikiShibata/gpl/ch02/ex02/unitconv"
)

func TestCToF(t *testing.T) {
	if unitconv.CToF(unitconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(unitconv.CToF(unitconv.BoilingC)))
	}
}

func TestKToC(t *testing.T) {
	if unitconv.KToC(0) != unitconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(unitconv.KToC(0)))
	}
}

func TestCToK(t *testing.T) {
	if unitconv.CToK(0) != -unitconv.Kelvin(unitconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(unitconv.CToK(0)))
	}
}

func TestFToK(t *testing.T) {
	expected := unitconv.FToC(0) - unitconv.AbsoluteZeroC

	if unitconv.FToK(0) != unitconv.Kelvin(expected) {
		t.Error(fmt.Sprint(unitconv.FToK(0)))
	}
}
