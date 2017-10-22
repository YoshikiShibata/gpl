// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestWordCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 2},
		{"Hello My World", 3},
		{"Hello My World ", 3},
		{"Hello World! こんにちは　世界", 4},
		{"Hello World!\nこんにちは　世界", 4},
	}

	var c WordCounter
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

		if c != WordCounter(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}
	}
}
