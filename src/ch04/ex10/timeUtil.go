// Copyright © 2016 Yoshiki Shibata

package main

import (
	"fmt"
	"time"
)

func LessThanMonth(t, now time.Time) bool {
	if t.After(now) {
		panic(fmt.Sprintf("Future is specified: %v", t))
	}

	nowYear, nowMonth, nowDay := now.Date()
	year, month, day := t.Date()

	nowMonths := nowYear*12 + int(nowMonth)
	months := year*12 + int(month)

	if months == nowMonths {
		return true
	}

	if (nowMonths - months) >= 2 {
		return false
	}

	return day > nowDay
}

func LessThanYear(t, now time.Time) bool {
	if t.After(now) {
		panic(fmt.Sprintf("Future is specified: %v", t))
	}

	nowYear, nowMonth, nowDay := now.Date()
	year, month, day := t.Date()

	nowMonths := nowYear*12 + int(nowMonth)
	months := year*12 + int(month)

	if (nowMonths - months) >= 13 {
		return false
	}

	if (nowMonths - months) <= 11 {
		return true
	}

	return day > nowDay
}
