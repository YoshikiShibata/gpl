// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package params

import "testing"

//+ Exercise 12.11

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

type notagdata struct {
	Labels     []string
	MaxResults int
	Exact      bool
}

func TestPack(t *testing.T) {
	for _, test := range []struct {
		d      data
		params string
	}{
		{data{Labels: []string{"golang"}, MaxResults: 10, Exact: true},
			"l=golang&max=10&x=true"},
		{data{Labels: []string{"golang", "programming"}, MaxResults: 10, Exact: true},
			"l=golang&l=programming&max=10&x=true"},
	} {
		params := Pack(&test.d)

		if params != test.params {
			t.Error("params = %q, but want %q\n", params, test.params)
		}
	}
}

func TestPackNoTag(t *testing.T) {
	for _, test := range []struct {
		d      notagdata
		params string
	}{
		{notagdata{Labels: []string{"golang", "programming"}, MaxResults: 10, Exact: true},
			"labels=golang&labels=programming&maxresults=10&exact=true"},
	} {
		params := Pack(&test.d)

		if params != test.params {
			t.Error("params = %q, but want %q\n", params, test.params)
		}
	}
}

//- Exercise 12.11
