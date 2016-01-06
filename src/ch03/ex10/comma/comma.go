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

func CommaWithBuffer(s string) string {
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

func CommaWithoutRecursion0(s string) string {
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

func CommaWithoutRecursion1(s string) string {
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

func CommaWithoutRecursion2(s string) string {
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

// BenchmarkComma-8                 	 3000000	       466 ns/op
// BenchmarkCommaWithBuffer-8       	 1000000	      1710 ns/op
// BenchmarkCommaWithoutRecursion0-8	 2000000	       870 ns/op
// BenchmarkCommaWithoutRecursion1-8	 1000000	      1897 ns/op
// BenchmarkCommaWithoutRecursion2-8	 1000000	      1877 ns/op
