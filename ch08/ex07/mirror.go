// Copyright © 2015 Yoshiki Shibata. All rights reserved

package main

import (
	"fmt"
	"io"
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

	return extractByType(resp, url, host+"/root.html", host, hostURL)
}

func extractAsFile(url, path, host, hostURL string) error {
	fmt.Printf("Contents of %s should be stored as %s\n", url, path)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	return extractByType(resp, url, path, host, hostURL)
}

func extractByType(resp *http.Response, url, path, host, hostURL string) error {
	contentType := extractContentType(resp.Header)
	if contentType[0] != TextHTML {
		f, err := os.Create(path)
		if err != nil {
			return nil
		}
		defer f.Close()
		io.Copy(f, resp.Body)
		resp.Body.Close()
		f.Close()
		return nil
	}

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
					if a.Val != "/" && a.Val != "#" && link.String() != hostURL {
						if a.Val[0] == '/' {
							extractAsFile(link.String(), host+a.Val, host, hostURL)
						} else {
							extractAsFile(link.String(), host+"/"+a.Val, host, hostURL)
						}
					}
				}
			}
		} else if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key != "src" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				extractAsFile(link.String(), host+"/"+a.Val, host, hostURL)
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer f.Close()
	html.Render(f, doc)
	return nil
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
