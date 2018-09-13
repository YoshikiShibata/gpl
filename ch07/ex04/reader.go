// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.

package main

import "io"

type reader struct {
	bytes []byte
	next  int
}

func NewReader(contents string) io.Reader {
	return &reader{[]byte(contents), 0}
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.next >= len(r.bytes) {
		return 0, io.EOF
	}

	nBytes := len(r.bytes) - r.next
	if nBytes > len(p) {
		nBytes = len(p)
	}

	copy(p, r.bytes[r.next:r.next+nBytes])
	r.next += nBytes
	return nBytes, nil
}
