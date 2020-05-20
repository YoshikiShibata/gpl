// Copyright © 2015, 2016, 2020 Yoshiki Shibata.

package comma

import (
	"bytes"
)

// Comma inserts commas in a non-negative decimal integer string
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// WithoutBuffer doesn't use bytes.Buffer
func WithoutBuffer(s string) string {
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

// WithDefaultBuffer uses a default Buffer.
func WithDefaultBuffer(s string) string {
	var buf bytes.Buffer // default Buffer
	utf8bytes := []byte(s)

	// when bc (byte counter) is zero, then insert a comma
	bc := len(utf8bytes) % 3
	if bc == 0 {
		bc = 3
	}

	for _, r := range utf8bytes {
		if bc == 0 {
			buf.WriteByte(',')
			bc = 3
		}
		buf.WriteByte(r)
		bc--
	}
	return buf.String()
}

// WithOptimalBufferUsingWriteString uses a optimal size of bytes.Buffer
func WithOptimalBufferUsingWriteString(s string) string {
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

// WithOptimalBufferUsingWrite uses a optimal size of bytes.Buffer
func WithOptimalBufferUsingWrite(s string) string {
	b := []byte(s)
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
