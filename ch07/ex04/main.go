// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if !Parse(os.Args[1]) {
		os.Exit(1)
	}
}

func Parse(contents string) bool {
	doc, err := html.Parse(newReader(contents))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		return false
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	return true
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type reader struct {
	bytes []byte
	next  int
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.next >= len(r.bytes) {
		return 0, io.EOF
	}

	nBytes := len(r.bytes) - r.next
	if nBytes > len(p) {
		nBytes = len(p)
	}

	copy(p, r.bytes[r.next:r.next+nBytes])
	r.next += nBytes
	return nBytes, nil
}

func newReader(contents string) io.Reader {
	return &reader{[]byte(contents), 0}
}
