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

func TestCommaWithoutRecursion0(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		result := comma.CommaWithoutRecursion0(inputs[i])
		if result != expecteds[i] {
			t.Errorf("Result is %s, want %s", result, expecteds[i])
		}
	}
}

func TestCommaWithoutRecursion1(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		result := comma.CommaWithoutRecursion1(inputs[i])
		if result != expecteds[i] {
			t.Errorf("Result is %s, want %s", result, expecteds[i])
		}
	}
}

func TestCommaWithoutRecursion2(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		result := comma.CommaWithoutRecursion2(inputs[i])
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

func BenchmarkCommaWithoutRecursion0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.CommaWithoutRecursion0(input)
		}
	}
}

func BenchmarkCommaWithoutRecursion1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.CommaWithoutRecursion1(input)
		}
	}
}

func BenchmarkCommaWithoutRecursion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.CommaWithoutRecursion1(input)
		}
	}
}
