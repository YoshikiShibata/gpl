// Copyright © 2016 Yoshiki Shibata
package main

import (
	"math/rand"
	"testing"
	"time"
)

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

func TestRotateRandomSize(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	dataSize := int(r.Int31n(1<<24)) + 1   // in case 0
	rotateSize := int(r.Int31n(1<<10)) + 1 // in case 0

	a := make([]int, dataSize)

	for i := 0; i < dataSize; i++ {
		a[i] = i
	}

	RotateCycles(a, rotateSize)
	for i := 0; i < dataSize; i++ {
		j := (i + rotateSize) % dataSize
		if a[j] != i {
			t.Errorf("a[%d] is %d, but want %d", j, a[j], i)
		}
	}
}
