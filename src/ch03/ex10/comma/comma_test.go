package comma_test

import (
	"ch03/ex10/comma"
	"testing"
)

var inputs = []string{"1", "123", "1234", "123456", "1234567"}
var expecteds = []string{"1", "123", "1,234", "123,456", "1,234,567"}

var data = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"1", "1"},
	{"123", "123"},
	{"1234", "1,234"},
	{"123456", "123,456"},
	{"1234567", "1,234,567"},
	{"0123456789", "0,123,456,789"},
}

func TestComma(t *testing.T) {
	for _, d := range data {
		result := comma.Comma(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}

func execute(t *testing.T, f func(string) string) {
	for _, d := range data {
		result := f(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}

func TestCommaWithBuffer(t *testing.T) {
	execute(t, comma.CommaWithBuffer)
}

func TestCommaWithoutRecursion0(t *testing.T) {
	execute(t, comma.CommaWithoutRecursion0)
}

func TestCommaWithoutRecursion1(t *testing.T) {
	execute(t, comma.CommaWithoutRecursion1)
}

func TestCommaWithoutRecursion2(t *testing.T) {
	execute(t, comma.CommaWithoutRecursion2)
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.Comma(input)
		}
	}
}

func BenchmarkCommaWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			comma.CommaWithBuffer(input)
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