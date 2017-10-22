// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package xmltree

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// Node represents either CharData or *Element
type Node interface{}

// CharData represents the contents of xml.CharData
type CharData string

// Element represents an element in a XML
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("<%s", e.Type.Local))
	for _, attr := range e.Attr {
		buf.WriteString(fmt.Sprintf(" %s='%s'", attr.Name.Local, attr.Value))
	}

	if len(e.Children) == 0 {
		buf.WriteString("/>\n")
		return buf.String()
	}

	buf.WriteString(">\n")
	for _, child := range e.Children {
		switch n := child.(type) {
		case *Element:
			buf.WriteString(n.String())
		case CharData:
			buf.WriteString(string(n))
		default:
			panic(fmt.Errorf("Unknown Type: %T", e))
		}
	}

	buf.WriteString(fmt.Sprintf("\n</%s>\n", e.Type.Local))
	return buf.String()
}

func Build(r io.Reader) (*Element, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element

	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			e := &Element{t.Name, t.Attr, nil}
			if len(stack) > 0 {
				index := len(stack) - 1 // index the last Element
				stack[index].Children = append(stack[index].Children, e)
			}
			stack = append(stack, e)

		case xml.EndElement:
			switch {
			case len(stack) == 0:
				panic("Stack is already empty")
			case len(stack) == 1:
				return stack[0], nil
			case len(stack) > 1:
				stack = stack[:len(stack)-1] // pop
			default:
				panic("Impossible")
			}

		case xml.CharData:
			if len(stack) == 0 {
				panic("Empty stack")
			}
			index := len(stack) - 1 // index the last Element
			stack[index].Children = append(stack[index].Children, CharData(t))
		}
	}
	panic("Unreachable")
}
