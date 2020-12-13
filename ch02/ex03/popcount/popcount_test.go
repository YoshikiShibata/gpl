// Copyright © 2015, 2017 Yoshiki Shibata

package popcount_test

import (
	"testing"

	"github.com/YoshikiShibata/gpl/ch02/ex03/popcount"
)

func TestZero(t *testing.T) {
	testZero(t, popcount.PopCount)
	testZero(t, popcount.PopCountWithLoop)
}

func testZero(t *testing.T, popCount func(uint64) int) {
	output := popCount(0)

	if output != 0 {
		t.Errorf("PopCount is %d, want 0", output)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, popcount.PopCount)
	testAllBits(t, popcount.PopCountWithLoop)
}

func testAllBits(t *testing.T, popCount func(uint64) int) {
	output := popCount(0xffffffffffffffff)

	if output != 64 {
		t.Errorf("PopCount is %d, want 64", output)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, popcount.PopCount)
	testEachByte(t, popcount.PopCountWithLoop)
}

func testEachByte(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 8; i++ {
		var value uint64 = 0xff << (uint(i) * 8)
		output := popCount(value)

		if output != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, output)
		}
	}
}

func Test0x5555(t *testing.T) {
	test0x5555(t, popcount.PopCount)
	test0x5555(t, popcount.PopCountWithLoop)
}

func test0x5555(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 4; i++ {
		var value uint64 = 0x5555 << (uint(i) * 8)
		output := popCount(value)

		if output != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, output)
		}
	}
}

func TestEachOneBit(t *testing.T) {
	testEachOneBit(t, popcount.PopCount)
	testEachOneBit(t, popcount.PopCountWithLoop)
}

func testEachOneBit(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 64; i++ {
		var value uint64 = 1 << uint(i)
		output := popCount(value)

		if output != 1 {
			t.Errorf("PopCount(%x) is %d, want 1", value, output)
		}
	}
}

// Exported (global) variable to store function outputs
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var output int

func BenchmarkPopCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCount(0x1234567890ABCDEF)
	}
	output = s
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCountWithLoop(0x1234567890ABCDEF)
	}
	output = s
}

/*
BenchmarkPopCount-8        	200000000	         5.65 ns/op
BenchmarkPopCountWithLoop-8	100000000	        11.4 ns/op
*/

/* Go 1.10 beta 2017-12-09
BenchmarkPopCount-4           	1000000000	         2.24 ns/op
BenchmarkPopCountWithLoop-4   	50000000	        25.4 ns/op
*/

/* Go 1.16 tip 2020-12-13: MacBook Pro (M1: Apple silicon)
BenchmarkPopCount               631673558                1.899 ns/op
BenchmarkPopCountWithLoop       185293153                6.479 ns/op
*/
