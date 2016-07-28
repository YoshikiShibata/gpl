// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package unarchive

import "testing"

func TestNoRegistered(t *testing.T) {
	_, err := OpenReader("top.zip")

	if err != ErrFormat {
		t.Logf("%v\n", err)
		t.Fail()
	}
}
