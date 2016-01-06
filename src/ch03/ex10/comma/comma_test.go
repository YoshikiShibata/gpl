// Copyright © 2015, 2016 Yoshiki Shibata.

package comma_test

import (
	"ch03/ex10/comma"
	"testing"
)

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

func TestCommaWithBuffer0(t *testing.T) {
	execute(t, comma.CommaWithBuffer0)
}

func TestCommaWithBuffer1(t *testing.T) {
	execute(t, comma.CommaWithBuffer1)
}

func TestCommaWithBuffer2(t *testing.T) {
	execute(t, comma.CommaWithBuffer2)
}

func TestCommaWithoutBuffer(t *testing.T) {
	execute(t, comma.CommaWithoutBuffer)
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.Comma(d.input)
		}
	}
}

func BenchmarkCommaWithoutBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.CommaWithoutBuffer(d.input)
		}
	}
}

func BenchmarkCommaWithBuffer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.CommaWithBuffer0(d.input)
		}
	}
}

func BenchmarkCommaWithBuffer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.CommaWithBuffer1(d.input)
		}
	}
}

func BenchmarkCommaWithBuffer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.CommaWithBuffer2(d.input)
		}
	}
}

// BenchmarkComma-8             	 1000000	      1013 ns/op
// BenchmarkCommaWithoutBuffer-8	 1000000	      1459 ns/op
// BenchmarkCommaWithBuffer0-8  	 1000000	      2394 ns/op
// BenchmarkCommaWithBuffer1-8  	 1000000	      2845 ns/op
// BenchmarkCommaWithBuffer2-8  	  500000	      2966 ns/op
