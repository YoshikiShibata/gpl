// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package unarchive_test

import (
	"ch10/ex02/unarchive"
	_ "ch10/ex02/unarchive/zip"
	"io"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	r, err := unarchive.OpenReader("top.zip")

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

		if !f.FileInfo().Mode().IsDir() {
			rc, err := f.Open()
			if err != nil {
				t.Fatalf("%v", err)
			}
			_, err = io.Copy(os.Stdout, rc)
			if err != nil {
				t.Fatalf("%v", err)
			}
			rc.Close()
		}
	}
}
