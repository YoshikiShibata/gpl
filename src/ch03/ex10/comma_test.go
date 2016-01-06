// Copyright © 2016 Yoshiki Shibata

package main

import "testing"

func TestComma(t *testing.T) {
	data := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"1", "1"},
		{"123", "123"},
		{"1234", "1,234"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"0123456789", "0,123,456,789"},
	}

	for _, d := range data {
		result := comma(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}
