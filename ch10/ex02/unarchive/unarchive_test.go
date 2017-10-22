// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package unarchive_test

import (
	"ch10/ex02/unarchive"
	_ "ch10/ex02/unarchive/tar"
	_ "ch10/ex02/unarchive/zip"
	"io"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	readArchive(t, "top.zip")
}

func TestTar(t *testing.T) {
	readArchive(t, "top.tar")
}

func readArchive(t *testing.T, name string) {
	r, err := unarchive.OpenReader(name)

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
