// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.

// Outline prints the outline of an HTML document tree.
// findElement prints the first HTML element with the specified id attribute
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: findElement url id")
		os.Exit(1)
	}

	if err := findElement(os.Args[1], os.Args[2]); err != nil {
		fmt.Printf("findElement failed: %v\n", err)
	}
}

func findElement(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := ElementByID(doc, id)
	if node == nil {
		fmt.Printf("No Element with \"%s\" attribute Found\n", id)
	} else {
		printNode(node)
	}
	return nil
}

func printNode(n *html.Node) {
	fmt.Printf("<%s", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s='%s'", a.Key, a.Val)
	}
	fmt.Println(">")
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
// The bool return value is used to continue or stop:
// true to continue, false to stop.
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}

// ElementsByID finds the first HTML element with the specified id attribute.
func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == id {
				node = n
				return false
			}
		}
		return true
	}, nil)

	return node
}
