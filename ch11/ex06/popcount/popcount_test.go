// Copyright © 2015, 2017 Yoshiki Shibata

package popcount_test

import (
	"testing"

	"github.com/YoshikiShibata/gpl/ch11/ex06/popcount"
)

func TestZero(t *testing.T) {
	testZero(t, popcount.PopCount)
	testZero(t, popcount.PopCountByShifting)
	testZero(t, popcount.PopCountByClearingBit)
	testZero(t, popcount.BitCount)
}

func testZero(t *testing.T, popCount func(uint64) int) {
	result := popCount(0)

	if result != 0 {
		t.Errorf("PopCount is %d, want 0", result)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, popcount.PopCount)
	testAllBits(t, popcount.PopCountByShifting)
	testAllBits(t, popcount.PopCountByClearingBit)
	testAllBits(t, popcount.BitCount)
}

func testAllBits(t *testing.T, popCount func(uint64) int) {
	result := popCount(0xffffffffffffffff)

	if result != 64 {
		t.Errorf("PopCount is %d, want 64", result)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, popcount.PopCount)
	testEachByte(t, popcount.PopCountByShifting)
	testEachByte(t, popcount.PopCountByClearingBit)
	testEachByte(t, popcount.BitCount)
}

func testEachByte(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 8; i++ {
		var value uint64 = 0xff << (uint(i) * 8)
		result := popCount(value)

		if result != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, result)
		}
	}
}

func Test0x5555(t *testing.T) {
	test0x5555(t, popcount.PopCount)
	test0x5555(t, popcount.PopCountByShifting)
	test0x5555(t, popcount.PopCountByClearingBit)
	test0x5555(t, popcount.BitCount)
}

func test0x5555(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 4; i++ {
		var value uint64 = 0x5555 << (uint(i) * 8)
		result := popCount(value)

		if result != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, result)
		}
	}
}

func TestEachOneBit(t *testing.T) {
	testEachOneBit(t, popcount.PopCount)
	testEachOneBit(t, popcount.PopCountByShifting)
	testEachOneBit(t, popcount.PopCountByClearingBit)
	testEachOneBit(t, popcount.BitCount)
}

func testEachOneBit(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 64; i++ {
		var value uint64 = 1 << uint(i)
		result := popCount(value)

		if result != 1 {
			t.Errorf("PopCount(%x) is %d, want 1", value, result)
		}
	}
}

// Exported (global) variable to store function outputs
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var output int

func benchmarkPopCount(b *testing.B, v uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCount(v)
	}
	output = s
}

func benchmarkPopCountByShifting(b *testing.B, v uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCountByShifting(v)
	}
	output = s
}

func benchmarkPopCountByClearingBit(b *testing.B, v uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.PopCountByClearingBit(v)
	}
	output = s
}

func benchmarkBitCount(b *testing.B, v uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += popcount.BitCount(v)
	}
	output = s
}

const allOnes = 0xFFFFFFFFFFFFFF

func BenchmarkPopCount_0(b *testing.B)        { benchmarkPopCount(b, 0x0) }
func BenchmarkPopCount_FFFF(b *testing.B)     { benchmarkPopCount(b, 0xFFFF) }
func BenchmarkPopCount_FFFFFFFF(b *testing.B) { benchmarkPopCount(b, 0xFFFFFFFF) }
func BenchmarkPopCount_AllOnes(b *testing.B)  { benchmarkPopCount(b, allOnes) }

func BenchmarkPopCountbyShifting_0(b *testing.B)        { benchmarkPopCountByShifting(b, 0x0) }
func BenchmarkPopCountbyShifting_FFFF(b *testing.B)     { benchmarkPopCountByShifting(b, 0xFFFF) }
func BenchmarkPopCountbyShifting_FFFFFFFF(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFFFFFF) }
func BenchmarkPopCountbyShifting_AllOnes(b *testing.B)  { benchmarkPopCountByShifting(b, allOnes) }

func BenchmarkPopCountByClearingBit_0(b *testing.B)    { benchmarkPopCountByClearingBit(b, 0x0) }
func BenchmarkPopCountByClearingBit_FFFF(b *testing.B) { benchmarkPopCountByClearingBit(b, 0xFFFF) }
func BenchmarkPopCountByClearingBit_FFFFFFFF(b *testing.B) {
	benchmarkPopCountByClearingBit(b, 0xFFFFFFFF)
}
func BenchmarkPopCountByClearingBit_AllOnes(b *testing.B) { benchmarkPopCountByClearingBit(b, allOnes) }

func BenchmarkBitCount_0(b *testing.B)        { benchmarkBitCount(b, 0x0) }
func BenchmarkBitCount_FFFF(b *testing.B)     { benchmarkBitCount(b, 0xFFFF) }
func BenchmarkBitCount_FFFFFFFF(b *testing.B) { benchmarkBitCount(b, 0xFFFFFFFF) }
func BenchmarkBitCount_AllOnes(b *testing.B)  { benchmarkBitCount(b, allOnes) }

/*
BenchmarkPopCount-8          	200000000	        5.66 ns/op
BenchmarkPopCountByShifting-8	20000000	        67.6 ns/op
BenchmarkPopByClearingBit-8  	50000000	        28.2 ns/op
BenchmarkBitCount-8          	1000000000	        2.29 ns/op
*/

/* Go1.10 beta 2017-12-9
goos: darwin
goarch: amd64
pkg: github.com/YoshikiShibata/gpl/ch11/ex06/popcount
BenchmarkPopCount_0-4                       	300000000	         4.52 ns/op
BenchmarkPopCountbyShifting_0-4             	20000000	        86.0 ns/op
BenchmarkPopCountByClearingBit_0-4          	1000000000	         2.48 ns/op
BenchmarkBitCount_0-4                       	500000000	         3.45 ns/op
BenchmarkPopCount_FFFF-4                    	300000000	         4.67 ns/op
BenchmarkPopCountbyShifting_FFFF-4          	20000000	        84.1 ns/op
BenchmarkPopCountByClearingBit_FFFF-4       	100000000	        12.1 ns/op
BenchmarkBitCount_FFFF-4                    	500000000	         3.31 ns/op
BenchmarkPopCount_FFFFFFFF-4                	300000000	         4.57 ns/op
BenchmarkPopCountbyShifting_FFFFFFFF-4      	20000000	        85.4 ns/op
BenchmarkPopCountByClearingBit_FFFFFFFF-4   	100000000	        24.0 ns/op
BenchmarkBitCount_FFFFFFFF-4                	500000000	         3.29 ns/op
BenchmarkPopCount_AllOnes-4                 	300000000	         4.41 ns/op
BenchmarkPopCountbyShifting_AllOnes-4       	20000000	        82.8 ns/op
BenchmarkPopCountByClearingBit_AllOnes-4    	30000000	        53.6 ns/op
BenchmarkBitCount_AllOnes-4                 	500000000	         3.21 ns/op
*/

/* Go1.12 tip 2018-12-30
macOS Mojava version 10.14.1
MacBook (Retina, 12-inch, Early 2015)
1.3GHz Intel Core M
8 GB 1600 MHz DDR3

goos: darwin
goarch: amd64
pkg: github.com/YoshikiShibata/gpl/ch11/ex06/popcount
BenchmarkPopCount_0-4                       	300000000	         4.75 ns/op
BenchmarkPopCount_FFFF-4                    	300000000	         4.71 ns/op
BenchmarkPopCount_FFFFFFFF-4                	300000000	         4.99 ns/op
BenchmarkPopCount_AllOnes-4                 	300000000	         4.79 ns/op
BenchmarkPopCountbyShifting_0-4             	20000000	        75.3 ns/op
BenchmarkPopCountbyShifting_FFFF-4          	20000000	        75.6 ns/op
BenchmarkPopCountbyShifting_FFFFFFFF-4      	20000000	        74.9 ns/op
BenchmarkPopCountbyShifting_AllOnes-4       	20000000	        75.0 ns/op
BenchmarkPopCountByClearingBit_0-4          	1000000000	         2.74 ns/op
BenchmarkPopCountByClearingBit_FFFF-4       	100000000	        12.3 ns/op
BenchmarkPopCountByClearingBit_FFFFFFFF-4   	50000000	        25.0 ns/op
BenchmarkPopCountByClearingBit_AllOnes-4    	30000000	        50.0 ns/op
BenchmarkBitCount_0-4                       	500000000	         3.40 ns/op
BenchmarkBitCount_FFFF-4                    	500000000	         3.41 ns/op
BenchmarkBitCount_FFFFFFFF-4                	500000000	         3.45 ns/op
BenchmarkBitCount_AllOnes-4                 	500000000	         3.42 ns/op
*/
