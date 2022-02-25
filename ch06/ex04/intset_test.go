// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		values   []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 144, 9, 42}, 4},
		{[]int{1, 32, 32 << 1, 32 << 2, 32 << 3, 32 << 8}, 6},
	} {
		var x IntSet
		for _, v := range tc.values {
			x.Add(v)
		}
		if x.Len() != tc.expected {
			t.Errorf("x.Len() is %d, but want %v", x.Len(), tc.expected)
		}
	}

	var x IntSet
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet

	// Removes non-existing value
	x.Remove(1000)
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	for i := 0; i < max; i++ {
		x.Remove(i)
		if x.Has(i) {
			t.Errorf("x.Has(%d) is true, but want false", i)
			continue
		}
		if x.Len() != (max - i - 1) {
			t.Errorf("x.Len() is %d, but want %d", x.Len(), max-i-1)
		}
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	c := x.Copy()
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	for i := 0; i < max; i++ {
		if x.Has(i) {
			t.Errorf("x.Has(%d) is true, but want false", i)
		}
		if !c.Has(i) {
			t.Errorf("c.Has(%d) is false, but want true", i)
		}
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll()
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	x.AddAll(1, 2, 3, 4, 5)
	if x.Len() != 5 {
		t.Errorf("x.Len() is %d, but want 5", x.Len())
	}

	for i := 1; i <= 5; i++ {
		if !x.Has(i) {
			t.Errorf("x.Has(%d) is false, but want true", i)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{}},

		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{4, 5}},
		{[]int{1, 10, 100, 1000, 10000}, []int{1, 10}, []int{1, 10}},
		{[]int{1, 10}, []int{1, 10, 100, 1000, 10000}, []int{1, 10}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.IntersectWith(&y)
		if x.Len() != len(tc.expected) {
			t.Errorf("x.Len() is %d, but want %d", x.Len(), len(tc.expected))
			t.Errorf("t is %v, s is %v", tc.t, tc.s)
		}

		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("x.Has(%d) is false, but want true", value)
			}
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5}},

		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3}},
		{[]int{1, 10, 100, 1000, 10000}, []int{1, 10}, []int{100, 1000, 10000}},
		{[]int{1, 10}, []int{1, 10, 100, 1000, 10000}, []int{}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.DifferenceWith(&y)
		if x.Len() != len(tc.expected) {
			t.Errorf("x.Len() is %d, but want %d", x.Len(), len(tc.expected))
			t.Errorf("t is %v, s is %v", tc.t, tc.s)
		}

		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("x.Has(%d) is false, but want true", value)
			}
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},

		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3, 6, 7, 8}},
		{[]int{1, 10, 100, 1000, 10000}, []int{1, 10}, []int{100, 1000, 10000}},
		{[]int{1, 10}, []int{1, 10, 100, 1000, 10000}, []int{100, 1000, 10000}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.SymmetricDifference(&y)
		if x.Len() != len(tc.expected) {
			t.Errorf("x.Len() is %d, but want %d", x.Len(), len(tc.expected))
			t.Errorf("t is %v, s is %v", tc.t, tc.s)
		}

		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("x.Has(%d) is false, but want true", value)
			}
		}
	}
}

func TestEnums(t *testing.T) {
	for _, tc := range []struct {
		t []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 10, 100, 1000, 10000}},
		{[]int{1, 1 << 4, 1 << 6, 1 << 8, 1 << 10, 1 << 12, 1 << 14, 1 << 16}},
	} {
		var x IntSet
		x.AddAll(tc.t...)

		enums := x.Elems()
		if len(enums) != len(tc.t) {
			t.Errorf("len(enums) is %d, but want %d", len(enums), len(tc.t))
			t.Errorf("enums is %v, t is %v", enums, tc.t)
		}

		for _, value := range enums {
			if !x.Has(value) {
				t.Errorf("x.Has(%d) is false, but want true", value)
			}
		}
	}
}
