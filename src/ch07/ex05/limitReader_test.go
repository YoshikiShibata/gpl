// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestLimitReader(t *testing.T) {
	for _, test := range []struct {
		data    string
		bufSize int
		limit   int64
	}{
		{"0123456789", 10, 3},
	} {

	}
}
