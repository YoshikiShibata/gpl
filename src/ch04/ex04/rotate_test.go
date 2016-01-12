// Copyright © 2016 Yoshiki Shibata
package main

import "testing"

func TestRotates(t *testing.T) {
	data := []struct {
		source     []int
		expected   []int
		rotateSize int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{9, 10, 11, 0, 1, 2, 3, 4, 5, 6, 7, 8},
			3},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6, 7},
			4},
	}

	for _, d := range data {
		RotateCycles(d.source, d.rotateSize)
		for i := 0; i < len(d.source); i++ {
			if d.source[i] != d.expected[i] {
				t.Errorf("source[%d] is %d, but want %d",
					i, d.source[i], d.expected[i])
			}
		}
	}
}
