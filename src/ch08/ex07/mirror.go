// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
// package links
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	Extract(os.Args[1])
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	host := resp.Request.URL.Host
	fmt.Printf("host = %s\n", host)
	if err := os.Mkdir(host, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s directory exists! Please delete it\n", host)
			return nil, nil
		}

		fmt.Printf("os.Mkdir : %v\n", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	fmt.Printf("%#v\n", resp.Header)

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if a.Val != link.String() {
					fmt.Fprintf(os.Stderr, "[%s] vs [%s]\n", a.Val, link.String())
					n.Attr[i].Val = link.String()
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	f, err := os.Create(host + "/root.html")
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer f.Close()
	html.Render(f, doc)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
