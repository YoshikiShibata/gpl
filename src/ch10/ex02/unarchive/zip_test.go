// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package unarchive

import (
	"io"
	"testing"
)

func TestZip(t *testing.T) {
	r, err := OpenReader("top.zip")

	if err != nil {
		t.Logf("%v\n", err)
		t.Fail()
		return
	}

	for {
		f, err := r.Next()
		if err != nil {
			if err != io.EOF {
				t.Fatalf("%v", err)
			}
			break
		}
		t.Logf("Name: %s\n", f.Name())
		t.Logf("Name: %s\n", f.FileInfo().Name())
	}
}
