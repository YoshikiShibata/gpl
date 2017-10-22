// Copyright © 2016 Yoshiki Shibata.

package main

import "strings"

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	// check sign symbol first
	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + comma(s[1:])
	}

	// check if s contains a period
	pIndex := strings.IndexByte(s, '.')
	if pIndex >= 0 {
		return comma(s[:pIndex]) + s[pIndex:]
	}

	// normal insertion of a comma
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
