// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestMaxWithError(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		valid    bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3, 4, 5}, true, 5},
		{[]int{5, 4, 3, 2, 1}, true, 5},
		{[]int{1, 2, 5, 3, 4}, true, 5},
	} {
		m, err := maxWithError(test.vals...)
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

func TestMax(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 5},
		{[]int{1, 2, 5, 3, 4}, 5},
	} {
		m := max(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("m is %d, but want %d\n", m, test.expected)
		}
	}
}
