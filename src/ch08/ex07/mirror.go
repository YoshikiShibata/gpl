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
	"strings"

	"golang.org/x/net/html"
)

const (
	ContentType = "Content-Type"
	TextHTML    = "text/html"
)

func main() {
	if err := Extract(os.Args[1]); err != nil {
		fmt.Printf("%v\n", err)
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	host := resp.Request.URL.Host
	path := resp.Request.URL.Path
	scheme := resp.Request.URL.Scheme
	hostURL := scheme + "://" + host
	fmt.Printf("scheme = %s, host = %s, path = %s\n", scheme, host, path)
	fmt.Printf("hostURL = %s\n", hostURL)
	fmt.Printf("%#v\n", resp.Header)

	if err := os.Mkdir(host, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s directory exists! Please delete it\n", host)
			return nil
		}

		fmt.Printf("os.Mkdir : %v\n", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	contentType := extractContentType(resp.Header)

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

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
				if strings.HasPrefix(link.String(), hostURL) {
					// replace href because this is the same host
					n.Attr[i].Val = "file:" + a.Val
					if a.Val != "/" && link.String() != hostURL {
						extractAsFile(link.String(), host+"/"+a.Val)
					}
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	if path == "/" && contentType[0] == TextHTML {
		f, err := os.Create(host + "/root.html")
		if err != nil {
			fmt.Printf("%v\n", err)
			return err
		}
		defer f.Close()
		html.Render(f, doc)
	}

	fmt.Printf("path = %s, contentType[0] = %s\n", path, contentType[0])
	return nil
}

func extractAsFile(url, path string) {
	fmt.Printf("Contents of %s should be stored as %s\n", url, path)
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

func extractContentType(header http.Header) []string {
	contentType, ok := header[ContentType]
	if !ok {
		return nil
	}
	return strings.Split(contentType[0], ";")
}
