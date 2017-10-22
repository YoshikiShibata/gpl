// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	if doc.Type == html.ElementNode {
		for _, tagName := range name {
			if tagName == doc.Data {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, node := range ElementsByTagName(c, name...) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
