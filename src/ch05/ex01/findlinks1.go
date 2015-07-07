package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result
func visit(n *html.Node, links []string) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(n.FirstChild, links)
	links = visit(n.NextSibling, links)

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(doc, nil) {
		fmt.Println(link)
	}
}
