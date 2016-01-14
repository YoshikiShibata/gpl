// Copyright © 2016 Yoshiki Shibata

package main

import (
	"testing"
	"time"
)

func TestLessThanMonth(t *testing.T) {
	dates := []struct {
		nowYear  int
		nowMonth time.Month
		nowDay   int
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		{2016, time.January, 1, 2016, time.January, 1, true},
		{2016, time.January, 1, 2015, time.December, 31, true},
		{2016, time.January, 1, 2015, time.December, 2, true},
		{2016, time.January, 1, 2015, time.December, 1, false},
		{2016, time.January, 1, 2015, time.November, 30, false},
	}

	for _, d := range dates {
		now := time.Date(d.nowYear, d.nowMonth, d.nowDay, 12, 0, 0, 0, time.UTC)
		tm := time.Date(d.year, d.month, d.day, 12, 0, 0, 0, time.UTC)
		result := LessThanMonth(tm, now)
		if result != d.expected {
			t.Errorf("Result is %b, want %b: %v, %v",
				result, d.expected, tm, now)
		}
	}
}

func TestLessThanYear(t *testing.T) {
	dates := []struct {
		nowYear  int
		nowMonth time.Month
		nowDay   int
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		{2016, time.January, 1, 2016, time.January, 1, true},
		{2016, time.January, 1, 2015, time.January, 2, true},
		{2016, time.January, 1, 2015, time.January, 1, false},
	}

	for _, d := range dates {
		now := time.Date(d.nowYear, d.nowMonth, d.nowDay, 12, 0, 0, 0, time.UTC)
		tm := time.Date(d.year, d.month, d.day, 12, 0, 0, 0, time.UTC)
		result := LessThanYear(tm, now)
		if result != d.expected {
			t.Errorf("Result is %b, want %b: %v, %v",
				result, d.expected, tm, now)
		}
	}
}
