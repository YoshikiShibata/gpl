// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package main

import (
	"strconv"
	"testing"
)

var bad_prereqs = map[string][]string{
	"linear algebra": {"calculus"}, // bad
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var bad_prereqs2 = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures": {"discrete math"},
	"databases":       {"data structures"},
	"discrete math": {
		"intro to programming",
		"operating systems", /* bad */
	},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var good_prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TestTopoSort(t *testing.T) {
	for i, test := range []struct {
		prereqs     map[string][]string
		errExpected bool
	}{
		{bad_prereqs, true},
		{bad_prereqs2, true},
		{good_prereqs, false},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ts, err := topoSort(test.prereqs)
			t.Logf("ts : %v, err : %v", ts, err)
			if (err != nil) != test.errExpected {
				t.Errorf("(err != nil) is %t, but want %t",
					(err != nil), test.errExpected)
			}
		})

	}
}
