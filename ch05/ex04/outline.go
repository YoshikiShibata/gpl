// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// !+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		stack = appendLinkIfAvailable(stack, n)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func appendLinkIfAvailable(stack []string, n *html.Node) []string {
	switch n.Data {
	case "img", "script":
		stack = appendValue(stack, n, "src")
	case "link":
		v, ok := extractValue(n, "type")
		if ok && v == "text/css" {
			stack = appendValue(stack, n, "href")
		}
	case "a":
		stack = appendValue(stack, n, "href")
	}

	return stack
}

func appendValue(stack []string, n *html.Node, key string) []string {
	v, ok := extractValue(n, key)
	if ok {
		stack = append(stack, fmt.Sprintf("(%s)", v))
	}
	return stack
}

func extractValue(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}
