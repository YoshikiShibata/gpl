// Copyright © 2016 Yoshiki Shibata

package main

// This function eliminate adjacent duplicates from s.
func eliminateAdjacentDuplicates(s []string) []string {
	if len(s) == 0 {
		return s
	}

	current := 0
	for i := 0; i < len(s)-1; i++ {
		if s[current] != s[i+1] {
			s[current+1] = s[i+1]
			current++
			continue
		}
	}
	return s[:current+1]
}
