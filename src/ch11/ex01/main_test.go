// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	for _, test := range []struct {
		bytes   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("こんにちは、世界"),
			map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1, '、': 1, '世': 1, '界': 1},
			[]int{0, 0, 0, 8, 0},
			0},
		{
			[]byte("Hello, World"),
			map[rune]int{'H': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, ' ': 1, 'W': 1, 'r': 1, 'd': 1},
			[]int{0, 12, 0, 0, 0},
			0},
		{
			[]byte("Hello, World\300"), // the last byte is invalid
			map[rune]int{'H': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, ' ': 1, 'W': 1, 'r': 1, 'd': 1},
			[]int{0, 12, 0, 0, 0},
			1},
	} {
		counts, utflen, invalid, err := charcount(bytes.NewReader(test.bytes))

		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}

		// counts
		if len(counts) != len(test.counts) {
			t.Errorf("len(counts) is %d, want %d\n", len(counts), len(test.counts))
			continue
		}

		for k, v := range test.counts {
			count, ok := counts[k]
			if !ok {
				t.Errorf("%c is not included\n", k)
				continue
			}
			if count != v {
				t.Errorf("count for %c is %d, but want %d\n", k, count, v)
				continue
			}
		}

		// utflen
		if len(utflen) != len(test.utflen) {
			t.Errorf("len(utflen) is %d, want %d\n", len(utflen), len(test.utflen))
			continue
		}
		for i := 1; i < len(utflen); i++ {
			if utflen[i] != test.utflen[i] {
				t.Errorf("utflen[%d] is %d, but want %d\n", i, utflen[i], test.utflen[i])
				continue
			}
		}

		// invalid
		if invalid != test.invalid {
			t.Errorf("invalid is %d, but want %d\n", invalid, test.invalid)
		}

	}
}
