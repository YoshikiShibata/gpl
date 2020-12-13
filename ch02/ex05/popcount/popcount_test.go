// Copyright © 2015, 2017 Yoshiki Shibata

package popcount_test

import (
	"testing"

	"github.com/YoshikiShibata/gpl/ch02/ex05/popcount"
)

func TestZero(t *testing.T) {
	testZero(t, popcount.PopCount)
	testZero(t, popcount.PopCountByShifting)
	testZero(t, popcount.PopCountByClearingBit)
	testZero(t, popcount.BitCount)
	testZero(t, popcount.OnesCount)
}

func testZero(t *testing.T, popCount func(uint64) int) {
	output := popCount(0)

	if output != 0 {
		t.Errorf("PopCount is %d, want 0", output)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, popcount.PopCount)
	testAllBits(t, popcount.PopCountByShifting)
	testAllBits(t, popcount.PopCountByClearingBit)
	testAllBits(t, popcount.BitCount)
	testAllBits(t, popcount.OnesCount)
}

func testAllBits(t *testing.T, popCount func(uint64) int) {
	output := popCount(0xffffffffffffffff)

	if output != 64 {
		t.Errorf("PopCount is %d, want 64", output)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, popcount.PopCount)
	testEachByte(t, popcount.PopCountByShifting)
	testEachByte(t, popcount.PopCountByClearingBit)
	testEachByte(t, popcount.BitCount)
	testEachByte(t, popcount.OnesCount)
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
	test0x5555(t, popcount.PopCountByShifting)
	test0x5555(t, popcount.PopCountByClearingBit)
	test0x5555(t, popcount.BitCount)
	test0x5555(t, popcount.OnesCount)
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
	testEachOneBit(t, popcount.PopCountByShifting)
	testEachOneBit(t, popcount.PopCountByClearingBit)
	testEachOneBit(t, popcount.BitCount)
	testEachOneBit(t, popcount.OnesCount)
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

// Exported (global) variable serving as input for some
// of the benchmarks to ensure side-effect free calls
// are not optimized away.
var input uint64 = 0x1234567890ABCDEF

// Exported (global) variable to store function outputs
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var output int

func BenchmarkPopCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCount(input)
	}
	output = s
}

func BenchmarkPopCountByShifting(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCountByShifting(input)
	}
	output = s
}

func BenchmarkPopByClearingBit(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCountByClearingBit(input)
	}
	output = s
}

func BenchmarkBitCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.BitCount(input)
	}
	output = s
}

func BenchmarkOnesCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.OnesCount(input)
	}
	output = s
}

/*
BenchmarkPopCount-8          	200000000	        5.66 ns/op
BenchmarkPopCountByShifting-8	20000000	        67.6 ns/op
BenchmarkPopByClearingBit-8  	50000000	        28.2 ns/op
BenchmarkBitCount-8          	1000000000	        2.29 ns/op
*/

/* Go1.10 beta 2017-12-9
BenchmarkPopCount-4             	300000000	         4.50 ns/op
BenchmarkPopCountByShifting-4   	10000000	       100 ns/op
BenchmarkPopByClearingBit-4     	50000000	        39.6 ns/op
BenchmarkBitCount-4             	500000000	         3.13 ns/op
BenchmarkOnesCount-4            	2000000000	         0.94 ns/op
*/

/* Go 1.16 tip 2020-12-13: MacBook Pro (M1: Apple silicon)
BenchmarkPopCount-8                     707768000                1.696 ns/op
BenchmarkPopCountByShifting-8           33525171                35.82 ns/op
BenchmarkPopByClearingBit-8             100000000               11.93 ns/op
BenchmarkBitCount-8                     773205630                1.553 ns/op
BenchmarkOnesCount-8                    1000000000               0.5530 ns/op
*/
