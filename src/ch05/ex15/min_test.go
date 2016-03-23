// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestMinWithError(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		valid    bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3, 4, 5}, true, 1},
		{[]int{5, 4, 3, 2, 1}, true, 1},
		{[]int{5, 2, 1, 3, 4}, true, 1},
	} {
		m, err := minWithError(test.vals...)
		if !test.valid {
			if err == nil {
				t.Errorf("err != nil expected for %v\n", test.vals)
			}
			continue
		}
		if m != test.expected {
			t.Errorf("m is %d, but want %d\n", m, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{5, 2, 1, 3, 4}, 1},
	} {
		m := min(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("m is %d, but want %d\n", m, test.expected)
		}
	}
}
