// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	f, err := os.Open("index.html")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	if len(images) != 4 {
		t.Errorf("len(%v) is not 4", images)
	}

	if len(headings) != 18 {
		t.Errorf("len(%v) is not 18", headings)
	}
}
