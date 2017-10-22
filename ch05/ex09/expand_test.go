// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	hahaha := func(s string) string { return "Ha Ha Ha" }
	toUpper := func(s string) string { return strings.ToUpper(s) }

	for _, test := range []struct {
		s        string
		f        func(string) string
		expected string
	}{
		{"$hello world", hahaha, "Ha Ha Ha world"},
		{"$hello,world", hahaha, "Ha Ha Ha,world"},
		{"$hello $world", hahaha, "Ha Ha Ha Ha Ha Ha"},
		{"$hello,$world", hahaha, "Ha Ha Ha,Ha Ha Ha"},
		{"$hello world", toUpper, "HELLO world"},
		{"$hello,world", toUpper, "HELLO,world"},
		{"$hello $world", toUpper, "HELLO WORLD"},
		{"$hello,$world", toUpper, "HELLO,WORLD"},
	} {
		result := expand(test.s, test.f)
		if result != test.expected {
			t.Errorf("%s, but want %s\n", result, test.expected)
		}
	}
}
