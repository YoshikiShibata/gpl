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

func CommaWithoutRecursion(s string) string {
	n := len(s)
	start, end := 0, n%3
	if end == 0 {
		start, end = 0, 3
	}

	var buf bytes.Buffer

	for end <= n {
		buf.WriteString(s[start:end])
		if end < n {
			buf.WriteString(",")
		}
		start, end = end, end+3
	}
	return buf.String()
}
