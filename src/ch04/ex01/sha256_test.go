// Copyright © 2016 Yoshiki Shibata

package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

var data = []struct {
	d1       []byte
	d2       []byte
	expected int
}{
	{[]byte("x"), []byte("x"), 0},
	{[]byte("X"), []byte("X"), 0},
	{[]byte("Hello"), []byte("Hello"), 0},
}

func TestPopCountDiff(t *testing.T) {
	for _, d := range data {
		sum1 := sha256.Sum256(d.d1)
		sum2 := sha256.Sum256(d.d2)
		result := popCountDiff(sum1, sum2)
		if result != d.expected {
			t.Errorf("Result is %d, want %d", result, d.expected)
		}
	}
}

var data2 = []struct {
	d1       [sha256.Size]byte
	d2       [sha256.Size]byte
	expected int
}{
	{
		toFixed([]byte("01234567890123456789012345678901")),
		toFixed([]byte("01234567890123456789012345678901")),
		0},
	{
		toFixed([]byte("01234567890123456789012345678901")),
		toFixed([]byte("11234567890123456789012345678901")),
		1},
	{
		toFixed([]byte("01234567890123456789012345678901")),
		toFixed([]byte("11234567891123456789112345678911")),
		4},
	{
		toFixed([]byte("01234567890123456789012345678901")),
		toFixed([]byte("31234567893123456789312345678931")),
		8},
}

func toFixed(bytes []byte) (fixed [sha256.Size]byte) {
	if len(bytes) != sha256.Size {
		panic(fmt.Sprintf("len is %d, want %d", len(bytes), sha256.Size))
	}
	for i := 0; i < sha256.Size; i++ {
		fixed[i] = bytes[i]
	}
	return
}

func TestPopCountDiff2(t *testing.T) {
	for _, d := range data2 {
		result := popCountDiff(d.d1, d.d2)
		if result != d.expected {
			t.Errorf("Result is %d, want %d", result, d.expected)
		}
	}
}
