// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "bytes"

// Join concatenates the elements of a to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}

	if len(a) == 1 {
		return a[0]
	}

	totalBytes := 0
	for _, s := range a {
		totalBytes += len(s)
	}
	totalBytes += len(sep) * (len(a) - 1)

	b := bytes.NewBuffer(make([]byte, 0, totalBytes))
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func Join2(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}

	if len(a) == 1 {
		return a[0]
	}

	n := len(sep) * (len(a) - 1)
	for _, s := range a {
		n += len(s)
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
