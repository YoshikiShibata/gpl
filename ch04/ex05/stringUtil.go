// Copyright © 2016, 2020 Yoshiki Shibata

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

func eliminateAdjacentDuplicates2(s []string) []string {
	if len(s) == 0 {
		return s
	}

	current := s[0]
	result := s[0:1]
	for _, next := range s[1:] {
		if current == next {
			continue
		}
		result = append(result, next)
		current = next
	}
	return result
}
