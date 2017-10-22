// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
topLoop:
	for _, tc := range []struct {
		data    string
		bufSize int
		limit   int64
	}{
		{"0123456789", 0, 3},
		{"0123456789", 1, 0},
		{"0123456789", 10, 3},
		{"", 10, 3},
	} {
		r1 := strings.NewReader(tc.data)
		r2 := strings.NewReader(tc.data)
		b1 := make([]byte, tc.bufSize)
		b2 := make([]byte, tc.bufSize)

		lr1 := LimitReader(r1, tc.limit)
		lr2 := io.LimitReader(r2, tc.limit)

		for i := 0; i < 2; i++ {
			n1, err1 := lr1.Read(b1)
			n2, err2 := lr2.Read(b2)
			t.Logf("n1 = %d, err1 = %v\n", n1, err1)
			t.Logf("n2 = %d, err2 = %v\n", n2, err2)
			if err1 != err2 {
				t.Errorf("err1=%v, err2=%v", err1, err2)
				continue topLoop
			}
			if n1 != n2 {
				t.Errorf("n1=%d, n2=%d", n1, n2)
				continue topLoop
			}
		}
	}
}
