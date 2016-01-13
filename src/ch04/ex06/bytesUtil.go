// Copyright © 2016 Yoshiki Shibata

package main

import (
	"unicode"
	"unicode/utf8"
)

// This function squashes each run of adjacent Unicode spaced
// in a UTF-8-encoded []byte slice into single ASCII space.
// This function is an in-place function.
func squashSpaces(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	spaceBuf := make([]byte, 4)
	spaceSize := utf8.EncodeRune(spaceBuf, ' ')
	spaceBuf = spaceBuf[:spaceSize]
	inSpace := false

	current := 0
	var size int
	var r rune
	for next := 0; next < len(b); next += size {
		r, size = utf8.DecodeRune(b[next:])

		if r == utf8.RuneError {
			panic("Rune Error")
		}

		if unicode.IsSpace(r) {
			if !inSpace {
				copy(b[current:], spaceBuf)
				current += spaceSize
				inSpace = true
			}
			continue
		}

		copy(b[current:], b[next:next+size])
		current += size
		inSpace = false
	}

	return b[:current]
}
