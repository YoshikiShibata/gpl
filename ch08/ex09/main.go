// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016, 2021 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walk(root, &wg)
	}
	wg.Wait()
}

func walk(root string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Traverse the root tree
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	n.Add(1)
	go walkDir(root, &n, fileSizes)

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(root, nfiles, nbytes)
		}
	}

	printDiskUsage(root, nfiles, nbytes) // final totals
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files  %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				log.Printf("entry.Info failed: %v", err)
			} else {
				fileSizes <- info.Size()
			}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
