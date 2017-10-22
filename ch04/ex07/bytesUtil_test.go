// Copyright © 2016 Yoshiki Shibata

package main

import "testing"

func TestReverseUtf8(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			"abcdefgh",
			"hgfedcba"},
		{
			"Hello World:こんにちは、世界",
			"界世、はちにんこ:dlroW olleH"},
	}

	for _, d := range data {
		b := []byte(d.s)
		ReverseUtf8(b)
		result := string(b)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}
