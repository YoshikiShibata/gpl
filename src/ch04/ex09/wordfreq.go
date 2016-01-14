// Copyright (C) 2015 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// wordfreq reports the frequency of each word in an input text file
func main() {
	if len(os.Args) == 1 {
		showUsage()
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		fmt.Printf("=== %s ===\n", f)
		showWordFreq(wordfreq(f))
	}
}

func showWordFreq(wf map[string]int) {
	var names []string
	for name := range wf {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, wf[name])
	}
}

func wordfreq(fileName string) map[string]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}

	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	freq := make(map[string]int)

	for input.Scan() {
		freq[input.Text()] += 1
	}

	return freq
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "usage: wordfreq files\n")
}
