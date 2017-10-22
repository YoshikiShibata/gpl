package main

import (
	"testing"
	"unicode"
)

type runes []rune

func (r runes) Len() int {
	return len(r)
}

func (r runes) Less(i, j int) bool {
	ri := r[i]
	rj := r[j]
	if unicode.IsLetter(ri) {
		ri = unicode.ToLower(ri)
	}
	if unicode.IsLetter(rj) {
		rj = unicode.ToLower(rj)
	}
	return ri < rj
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", false},
		{"Evil I did dwell; lewd did I live.", false},
		{"Able was I ere I saw Elba", true},
		{"ete", true},
		{"Et se resservir, ivresse reste.", false},
		{"palindrome", false}, // 非回文
		{"desserts", false},   // 半回文
	} {
		if got := IsPalindrome(runes([]rune(test.input))); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
