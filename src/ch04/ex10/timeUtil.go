// Copyright © 2016, 2017 Yoshiki Shibata

package main

import (
	"fmt"
	"time"
)

func LessThanMonth(t, now time.Time) bool {
	if t.After(now) {
		panic(fmt.Sprintf("Future is specified: %v", t))
	}

	oneMonth := now.AddDate(0, -1, 0)
	return t.After(oneMonth)
}

func LessThanYear(t, now time.Time) bool {
	if t.After(now) {
		panic(fmt.Sprintf("Future is specified: %v", t))
	}

	oneYear := now.AddDate(-1, 0, 0)
	return t.After(oneYear)
}
