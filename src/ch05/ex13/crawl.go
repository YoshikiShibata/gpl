// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Localcopies makes local copies of the pages, starting with the URLs on the command line.
// Localcopies don't make copies of pages that come from a different domain
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if isSameDomain(item) {
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

var initialURL *url.URL

func isSameDomain(item string) bool {
	u, err := url.Parse(item)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	return strings.HasSuffix(u.Host, initialURL.Host)
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.

	if len(os.Args) != 2 {
		fmt.Println("usage: crawl <url>")
		os.Exit(1)
	}
	var err error

	initialURL, err = url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(*initialURL)
	}

	breadthFirst(crawl, []string{os.Args[1]})
}
