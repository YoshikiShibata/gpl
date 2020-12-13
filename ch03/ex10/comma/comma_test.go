// Copyright © 2015, 2016, 2017, 2020 Yoshiki Shibata.

package comma_test

import (
	"testing"

	"github.com/YoshikiShibata/gpl/ch03/ex10/comma"
)

var data = [...]struct {
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

func execute(t *testing.T, f func(string) string) {
	for _, d := range data {
		result := f(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}

func TestComma(t *testing.T) {
	execute(t, comma.Comma)
}

func TestCommaWithDefaultBuffer(t *testing.T) {
	execute(t, comma.WithDefaultBuffer)
}

func TestCommaWithOptimalBuffer_UsingWriteString(t *testing.T) {
	execute(t, comma.WithOptimalBufferUsingWriteString)
}

func TestCommaWithOptimalBuffer_UsingWrite(t *testing.T) {
	execute(t, comma.WithOptimalBufferUsingWrite)
}

func TestCommaWithoutBuffer(t *testing.T) {
	execute(t, comma.WithoutBuffer)
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
			comma.WithoutBuffer(d.input)
		}
	}
}

func BenchmarkCommaWithDefaultBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.WithDefaultBuffer(d.input)
		}
	}
}

func BenchmarkCommaWithOptimalBuffer_UsingWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.WithOptimalBufferUsingWriteString(d.input)
		}
	}
}

func BenchmarkCommaWithOptimalBuffer_UsingWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			comma.WithOptimalBufferUsingWrite(d.input)
		}
	}
}

// BenchmarkComma-8             	 1000000	      1013 ns/op
// BenchmarkCommaWithoutBuffer-8	 1000000	      1459 ns/op
// BenchmarkCommaWithBuffer0-8  	 1000000	      2394 ns/op
// BenchmarkCommaWithBuffer1-8  	 1000000	      2845 ns/op
// BenchmarkCommaWithBuffer2-8  	  500000	      2966 ns/op

/* Go 1.16 tip 2020/12/13 MacBook Pro (M1: Apple silicon)
BenchmarkComma-8                                         4991583               240.1 ns/op
BenchmarkCommaWithoutBuffer-8                            3555876               338.2 ns/op
BenchmarkCommaWithDefaultBuffer-8                        3661735               327.4 ns/op
BenchmarkCommaWithOptimalBuffer_UsingWriteString-8       4701040               255.1 ns/op
BenchmarkCommaWithOptimalBuffer_UsingWrite-8             4663062               257.7 ns/op
*/
