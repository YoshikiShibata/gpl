// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestLineCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 1},
		{"Hello World\nHello World", 2},
		{"Hello World\nHello World\n", 2},
		{"Hello World\nHello World\n\n", 3},
		{"Hello World! こんにちは　世界", 1},
		{"Hello World!\nこんにちは　世界", 2},
	}

	var c LineCounter
	for _, d := range data {
		c = 0

		bytes := []byte(d.s)
		n, err := c.Write(bytes)

		if err != nil {
			t.Errorf("Unpexected Error : %v", err)
			continue
		}

		if n != len(bytes) {
			t.Errorf("Written bytes is %d, want %d", n, len(bytes))
			continue
		}

		if c != LineCounter(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}
	}
}
