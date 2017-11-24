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
	result := popCount(0)

	if result != 0 {
		t.Errorf("PopCount is %d, want 0", result)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, popcount.PopCount)
	testAllBits(t, popcount.PopCountWithLoop)
}

func testAllBits(t *testing.T, popCount func(uint64) int) {
	result := popCount(0xffffffffffffffff)

	if result != 64 {
		t.Errorf("PopCount is %d, want 64", result)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, popcount.PopCount)
	testEachByte(t, popcount.PopCountWithLoop)
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
	test0x5555(t, popcount.PopCountWithLoop)
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
	testEachOneBit(t, popcount.PopCountWithLoop)
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

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountWithLoop(0x1234567890ABCDEF)
	}
}

/*
BenchmarkPopCount-8        	200000000	         5.65 ns/op
BenchmarkPopCountWithLoop-8	100000000	        11.4 ns/op
*/
