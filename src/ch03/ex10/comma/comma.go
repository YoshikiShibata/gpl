// Copyright © 2015, 2016 Yoshiki Shibata.

package comma

import (
	"bytes"
)

// comma inserts commas in a non-negative decimal integer string
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func CommaWithBuffer0(s string) string {
	var buf bytes.Buffer
	runes := []byte(s)

	// when rc (rune counter) is zero, then insert a comma
	rc := len(runes) % 3
	if rc == 0 {
		rc = 3
	}

	for _, r := range runes {
		if rc == 0 {
			buf.WriteByte(',')
			rc = 3
		}
		buf.WriteByte(r)
		rc--
	}
	return buf.String()
}

func CommaWithoutBuffer(s string) string {
	// Optimization for short string (less than 4 digits)
	n := len(s)
	if n <= 3 {
		return s
	}

	var temp string

	for {
		if n <= 3 {
			return s + temp
		}
		temp = "," + s[n-3:] + temp
		s = s[:n-3]
		n = len(s)
	}

}

func CommaWithBuffer1(s string) string {
	n := len(s)
	start, end := 0, n%3
	if end == 0 {
		start, end = 0, 3
	}

	buf := bytes.NewBuffer(make([]byte, 0, n+(n-1)/3))

	for end <= n {
		buf.WriteString(s[start:end])
		if end < n {
			buf.WriteString(",")
		}
		start, end = end, end+3
	}
	return buf.String()
}

func CommaWithBuffer2(s string) string {
	b := ([]byte)(s)
	n := len(b)
	start, end := 0, n%3
	if end == 0 {
		start, end = 0, 3
	}

	buf := bytes.NewBuffer(make([]byte, 0, n+(n-1)/3))

	for end <= n {
		buf.Write(b[start:end])
		if end < n {
			buf.WriteByte(',')
		}
		start, end = end, end+3
	}
	return buf.String()
}
