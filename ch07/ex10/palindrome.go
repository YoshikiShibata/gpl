// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "sort"

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}
