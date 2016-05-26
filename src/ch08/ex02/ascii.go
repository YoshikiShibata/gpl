package main

import "io"

// asciiText implements io.Writer and insert necessary a CR character
type asciiText struct {
	w io.Writer
}

func (a *asciiText) Write(p []byte) (int, error) {
	buf := make([]byte, 0, len(p))
	var lastB byte

	for _, b := range p {
		if b == '\n' && lastB != '\r' {
			buf = append(buf, '\r')
		}
		buf = append(buf, b)
		lastB = b
	}

	n, err := a.w.Write(buf)
	if n > len(p) {
		return len(p), err
	}
	return n, err
}
