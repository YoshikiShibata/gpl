// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "io"

// CountingWriter returns a new Writer that wraps the given Writer, and a
// pointer to an int64 variable that at any moment contains the number of bytes
// written to the new Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper
	wp.w = w
	return &wp, &(wp.c)
}

type Wrapper struct {
	c int64
	w io.Writer
}

func (wp *Wrapper) Write(b []byte) (n int, err error) {
	n, err = wp.w.Write(b)
	wp.c += int64(n)
	return
}
