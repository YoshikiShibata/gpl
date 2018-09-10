// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.

// Outline prints the outline of an HTML document tree.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(prettyPrint(doc))

	return nil
}

func prettyPrint(doc *html.Node) string {
	var buf bytes.Buffer

	forEachNode(&buf, doc, startElement, endElement)

	return buf.String()
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(w io.Writer, n *html.Node,
	pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

var depth int

func startElement(w io.Writer, n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		startTextNode(w, n)
		return
	case html.DocumentNode:
	case html.ElementNode:
		startElementNode(w, n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
		startDocktypeNode(w, n)
	}
}

func endElement(w io.Writer, n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		return // don't pop
	case html.DocumentNode:
	case html.ElementNode:
		endElementNode(w, n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
	}
}

func startTextNode(w io.Writer, n *html.Node) {
	fmt.Fprintf(w, "%s", n.Data)
}

func startElementNode(w io.Writer, n *html.Node) {
	depth++
	if n.FirstChild == nil {
		return
	}

	attrs := attributes(n.Attr)
	if attrs == "" {
		fmt.Fprintf(w, "\n%*s<%s>", depth, "", n.Data)
	} else {
		fmt.Fprintf(w, "\n%*s<%s %s>", depth, "", n.Data, attrs)
	}
}

func endElementNode(w io.Writer, n *html.Node) {
	if n.FirstChild == nil {
		attrs := attributes(n.Attr)
		if attrs == "" {
			switch n.Data {
			case "br":
				fmt.Fprintf(w, "<%s/>\n", n.Data)
			default:
				fmt.Fprintf(w, "\n%*s<%s />", depth, "", n.Data)
			}
		} else {
			fmt.Fprintf(w, "\n%*s<%s %s />", depth, "", n.Data, attrs)
		}
	} else {
		switch n.Data {
		case "a", "code", "title", "tt", "h1":
			fmt.Fprintf(w, "</%s>", n.Data)
		default:
			fmt.Fprintf(w, "\n%*s</%s>", depth, "", n.Data)
		}
	}
	depth--
}

func attributes(attr []html.Attribute) string {
	var builder strings.Builder

	for i, a := range attr {
		if i != 0 {
			builder.WriteString(" ")

		}
		if a.Namespace == "" {
			builder.WriteString(a.Key)
			builder.WriteString(`="`)
			builder.WriteString(a.Val)
			builder.WriteString(`"`)
		} else {
			builder.WriteString(a.Namespace)
			builder.WriteString(":")
			builder.WriteString(a.Key)
			builder.WriteString(`="`)
			builder.WriteString(a.Val)
			builder.WriteString(`"`)
		}
	}
	return builder.String()
}

func startDocktypeNode(w io.Writer, n *html.Node) {
	if n.Type != html.DoctypeNode {
		panic("Illegal Argument")
	}

	var builder strings.Builder

	builder.WriteString("<!DOCTYPE ")
	builder.WriteString(n.Namespace)

	for i, a := range n.Attr {
		if i != 0 {
			builder.WriteString(" ")
		}

		if a.Key == "public" {
			builder.WriteString("PUBLIC ")
			builder.WriteString(`"`)
		}
	}

	io.WriteString(w, builder.String())
}
