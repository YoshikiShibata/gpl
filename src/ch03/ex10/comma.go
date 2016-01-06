// Copyright © 2016 Yoshiki Shibata

package main

import "bytes"

/* comma func in the text

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
*/

func comma(s string) string {
	var buf bytes.Buffer
	runes := []rune(s)

	// when rc (rune counter) is zero, then insert a comma
	rc := len(runes) % 3
	if rc == 0 {
		rc = 3
	}

	for _, r := range runes {
		if rc == 0 {
			buf.WriteRune(',')
			rc = 3
		}
		buf.WriteRune(r)
		rc--
	}
	return buf.String()
}
