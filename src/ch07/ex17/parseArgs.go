// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type attribute struct {
	name  string
	value string
}

func (a *attribute) String() string {
	return fmt.Sprintf("%s=\"%s\"", a.name, a.value)
}

type element struct {
	name       string
	attributes []attribute
}

func (e *element) String() string {
	var buff bytes.Buffer

	buff.WriteString("<")
	buff.WriteString(e.name)
	for _, attr := range e.attributes {
		buff.WriteByte(' ')
		buff.WriteString(attr.String())
	}
	buff.WriteString(">")
	return buff.String()
}

// parseArgs parses command line arguments and breaks them into a slice
// of elements. Any attribute will be recoginized as "name=value"
func parseArgs() []*element {
	var result []*element

	var elm *element
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") {
			if elm == nil {
				fmt.Printf("No element name is specified: [%s] ignored\n", arg)
				continue
			}
			nameValue := strings.Split(arg, "=")
			if len(nameValue) != 2 {
				fmt.Printf("Illegal format: [%s] ignored\n", arg)
			}
			attr := attribute{nameValue[0], nameValue[1]}
			elm.attributes = append(elm.attributes, attr)
			continue
		}

		if elm != nil {
			result = append(result, elm)
		}
		elm = &element{name: arg}
	}
	result = append(result, elm)
	return result
}
