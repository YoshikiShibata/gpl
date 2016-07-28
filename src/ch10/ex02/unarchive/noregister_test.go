// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package unarchive_test

import (
	"ch10/ex02/unarchive"
	"testing"
)

func TestNoRegistered(t *testing.T) {
	_, err := unarchive.OpenReader("top.zip")

	if err != unarchive.ErrFormat {
		t.Logf("%v\n", err)
		t.Fail()
	}
}
