// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	data := []string{
		"Hello World\n",
		"How are you?\n",
		"The Go Programming Language\n",
		"プログラミング言語Go\n",
	}

	w, c := CountingWriter(os.Stdout)

	var total int64 = 0

	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("count is %d, want %d", *c, total)
		}
	}
}
