// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		fmt.Printf("[%s]\n=> Words : %d, Images : %d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}

	input := bufio.NewScanner(strings.NewReader(n.Data))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		images += i
		words += w
	}

	return words, images
}
