package comma_test

import (
	"ch03/ex11/comma"
	"testing"
)

var inputs = []string{"1", "123", "1234", "123456", "1234567"}
var expecteds = []string{"1", "123", "1,234", "123,456", "1,234,567"}

func TestComma(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		result := comma.Comma(inputs[i])
		if result != expecteds[i] {
			t.Errorf("Result is %s, want %s", result, expecteds[i])
		}
	}
}

func TestCommaWithoutRecursion(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		result := comma.CommaWithoutRecursion(inputs[i])
		if result != expecteds[i] {
			t.Errorf("Result is %s, want %s", result, expecteds[i])
		}
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.Comma(input)
		}
	}
}

func BenchmarkCommaWithoutRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.CommaWithoutRecursion(input)
		}
	}
}
