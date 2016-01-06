// Copyright © 2016 Yoshiki Shibata.

package main

import (
	"bytes"
	"sort"
	"strings"
)

func areAnagrams(s1, s2 string) bool {
	bytes1 := toComparableSortedBytes(s1)
	bytes2 := toComparableSortedBytes(s2)

	return bytes.Compare(bytes1, bytes2) == 0
}

type Bytes []byte

func (b Bytes) Len() int           { return len(b) }
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func toComparableSortedBytes(s string) []byte {
	result := make([]byte, 0)

	for _, b := range []byte(strings.ToUpper(s)) {
		switch b {
		case '!', '.', ' ', '?':
			// skip
		default:
			result = append(result, b)
		}
	}
	sort.Sort(Bytes(result))
	return result
}
