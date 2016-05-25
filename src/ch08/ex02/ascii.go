package main

import "io"

type asciiText struct {
	w io.Writer
}

func (a *asciiText) Write(p []byte) (int, error) {
	cr := []byte{'\r'}
	start := 0
	total := 0
	for i, b := range p {
		if b == '\n' {
			if start < i {
				n, err := a.w.Write(p[start:i])
				if err != nil {
					return total, err
				}
				total += n
				start = i
			}
			_, err := a.w.Write(cr)
			if err != nil {
				return total, err
			}
		}
	}
	n, err := a.w.Write(p[start:len(p)])
	if err != nil {
		return total, err
	}
	total += n
	return total, nil
}
