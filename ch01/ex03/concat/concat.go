// Copyright © 2015, 2020 Yoshiki Shibata. All rights reserved.

package concat

// Concat concates strings with a space.
func Concat(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str // NB: inefficient!
		sep = " "
	}
	return s
}
