// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	elem := parseArgs()

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, elem) {
				fmt.Printf("%s: %s\n", joinStack(stack), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(stack []xml.StartElement, arg []*element) bool {
	for len(arg) <= len(stack) {
		if len(arg) == 0 {
			return true
		}
		if stack[0].Name.Local == arg[0].name {
			if containsAllAttributes(stack[0].Attr, arg[0].attributes) {
				arg = arg[1:]
			}
		}
		stack = stack[1:]
	}
	return false
}

func joinStack(stack []xml.StartElement) string {
	var result []string

	for _, elem := range stack {
		result = append(result, elem.Name.Local)
	}
	return strings.Join(result, " ")
}

func containsAllAttributes(stack []xml.Attr, arg []attribute) bool {
	for _, argAttr := range arg {
		matched := false
		for _, stackAttr := range stack {
			if stackAttr.Name.Local == argAttr.name {
				if stackAttr.Value != argAttr.value {
					return false
				} else {
					matched = true
				}
			}
		}
		if !matched {
			return false
		}
	}
	return true
}
