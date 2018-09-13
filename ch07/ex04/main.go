// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if !Parse(os.Args[1]) {
		os.Exit(1)
	}
}

func Parse(contents string) bool {
	doc, err := html.Parse(NewReader(contents))
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
