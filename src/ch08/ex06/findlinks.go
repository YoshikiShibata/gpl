// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type leveledList struct {
	depth int
	lists []string
}

var depthFlag = flag.Int("depth", math.MaxInt32, "depth of links")

func crawl(depth int, url string) *leveledList {
	if depth > *depthFlag {
		return &leveledList{depth + 1, nil}
	}

	fmt.Printf("%3d: %s\n", depth, url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return &leveledList{depth + 1, list}
}

func main() {
	flag.Parse()
	worklist := make(chan *leveledList)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- &leveledList{0, flag.Args()} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		llist := <-worklist
		for _, link := range llist.lists {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(llist.depth, link)
			}
		}
	}
}
