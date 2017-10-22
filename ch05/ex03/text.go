// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func printTextNode(stack []string, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		stack = append(stack, n.Data)
	case html.TextNode:
		sLen := len(stack)
		if sLen == 0 {
			panic("Impossible")
		}

		last := stack[sLen-1]
		if last != "script" && last != "style" {
			trimmed := strings.TrimSpace(n.Data)
			if len(trimmed) > 0 {
				fmt.Printf("<%s>%s</%s>\n", last, trimmed, last)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextNode(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text: %v\n", err)
		os.Exit(1)
	}

	printTextNode(nil, doc)
}
