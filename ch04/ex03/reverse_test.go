// Copyright © 2016 Yoshiki Shibata

package main

import "testing"

func TestReverse(t *testing.T) {
	var a [Size]int
	for i := 0; i < Size; i++ {
		a[i] = i
	}
	reverse(&a)
	for i := 0; i < Size; i++ {
		if a[i] != (Size - i - 1) {
			t.Errorf("a[%d] is %d, want %d", i, a[i], Size-i-1)
		}
	}
}
