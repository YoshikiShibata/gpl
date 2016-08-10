// Copyright © 2016 Yoshiki Shibata. All rights reserved

package intset

import "testing"

func TestAdd(t *testing.T) {
	var x IntSet
	y := make(map[int]bool)

	sets := []int{1, 144, 9, 9, 42}

	for _, i := range sets {
		x.Add(i)
		y[i] = true

		if x.Has(i) != true {
			t.Errorf("x.Has(%d) is false, but want true\n", i)
		}
	}

	for i := 0; i < 200; i++ {
		if x.Has(i) == y[i] {
			continue
		}
		t.Errorf("x.Has(%d) is %t and y[%d] is %t\n", i, x.Has(i), i, y[i])
	}
}
