// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// FindFiles walks a directory, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
)

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
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func listFiles(path string) []string {
	fmt.Println(path)
	dirInfos := extractFileInfos(path)

	var files []string
	for _, dirInfo := range dirInfos {
		name := dirInfo.Name()
		if name[0] == '.' {
			continue
		}
		files = append(files, path+"/"+dirInfo.Name())
	}
	return files
}

func extractFileInfos(path string) []os.FileInfo {
	f, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Print(err)
		return nil
	}

	if !fileInfo.IsDir() {
		return nil
	}

	dirInfos, err := f.Readdir(0) // all directories
	if err != nil {
		log.Print(err)
		return nil
	}
	return dirInfos
}

func main() {
	// Walk a direct breadth-first
	// starting from the command-line arguments.
	breadthFirst(listFiles, os.Args[1:])
}
