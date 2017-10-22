// Copyright © 2015 Yoshiki Shibata. All rights reserved.

package concat

func Concat(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str // NB: inefficient!
		sep = " "
	}
	return s
}
