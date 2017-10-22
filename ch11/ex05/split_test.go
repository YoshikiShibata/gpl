// Copyright © 2016 Yoshiki Shibata. All rights reserved

package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, test := range []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a,b,c,d", ",", 4},
	} {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
		// ...
	}
}
