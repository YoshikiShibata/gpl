// Copyright © 2016 Yoshiki Shibata.

package main

import "testing"

var data = []struct {
	s1       string
	s2       string
	expected bool
}{
	{"Hello", "Hello", true},
	{"Hello!", "Hello", true},
	{"Hello", "Ell Ho", true},
	{"Hello!", "Ell Ho", true},
	{"Hello", "Ell Oh", true},
	{"Hello!", "Ell Oh", true},
	{"Hello", "Hello W", false},
	{"Go Programming", "Programming Go", true},
	{"Yoshiki Shibata", "Asia Boyish Kith", true},
}

func TestAnagrams(t *testing.T) {
	for _, d := range data {
		result := areAnagrams(d.s1, d.s2)
		if result != d.expected {
			t.Errorf("Result is %v, want %v: \"%s\", \"%s\"", result, d.expected, d.s1, d.s2)
		}
	}
}
