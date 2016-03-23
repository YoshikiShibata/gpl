// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "fmt"

func minWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("zero length slice")
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func min(one int, vals ...int) int {
	m := one

	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}
