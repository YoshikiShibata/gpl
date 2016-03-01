// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	"ch09/ex03/memo"
	"ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.SequentialCancel(t, m)
	memotest.Sequential(t, m)
}

func TestConcurrentCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.ConcurrentCancel(t, m)
	memotest.Concurrent(t, m)
}
