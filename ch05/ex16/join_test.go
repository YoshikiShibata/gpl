// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	a        []string
	sep      string
	expected string
}{
	{nil, "", ""},
	{[]string{"abc"}, " ", "abc"},
	{[]string{"abc", "def", "ghi"}, " ", "abc def ghi"},
	{[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, " ",
		"0 1 2 3 4 5 6 7 8 9"},
}

func TestJoin(t *testing.T) {
	for _, tc := range testCases {
		result := Join(tc.sep, tc.a...)
		if result != tc.expected {
			t.Errorf("Join(%s, %v) = %s, want %s",
				tc.sep, tc.a, result, tc.expected)
		}
	}
}

func TestJoin2(t *testing.T) {
	for _, tc := range testCases {
		result := Join2(tc.sep, tc.a...)
		if result != tc.expected {
			t.Errorf("Join(%s, %v) = %s, want %s",
				tc.sep, tc.a, result, tc.expected)
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Join(tc.sep, tc.a...)
		}
	}
}

func BenchmarkJoin2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Join2(tc.sep, tc.a...)
		}
	}
}

func BenchmarkStrings_Join(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			strings.Join(tc.a, tc.sep)
		}
	}
}

// BenchmarkJoin-8           	 1000000	      1331 ns/op
// BenchmarkJoin2-8          	 3000000	       689 ns/op
// BenchmarkStrings_Join-8   	 2000000	       669 ns/op

// 2018.9.13 1.3GHz Intel Core M . 8GB 1600 MHz DDR3
// BenchmarkJoin-4           	 2000000	       700 ns/op
// BenchmarkJoin2-4          	 3000000	       446 ns/op
// BenchmarkStrings_Join-4   	 5000000	       263 ns/op
