package tempconv_test

import (
	"fmt"
	"testing"

	"ch02/ex01/tempconv"
)

func TestCToF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(tempconv.CToF(tempconv.BoilingC)))
	}
}

func TestKToC(t *testing.T) {
	if tempconv.KToC(0) != tempconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(tempconv.KToC(0)))
	}
}

func TestCToK(t *testing.T) {
	if tempconv.CToK(0) != -tempconv.Kelvin(tempconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(tempconv.CToK(0)))
	}
}
