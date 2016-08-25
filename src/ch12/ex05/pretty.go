// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

// This file implements the algorithm described in Derek C. Oppen's
// 1979 Stanford technical report, "Pretty Printing".

import (
	"bytes"
	"fmt"
	"reflect"
)

func MarshalIndent(v interface{}) ([]byte, error) {
	p := printer{width: margin}
	if err := pretty(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

const margin = 80

type token struct {
	kind rune // one of "s ()" (string, blank, start, end)
	str  string
	size int
}

type printer struct {
	tokens []*token // FIFO buffer
	stack  []*token // stack of open ' ' and '(' tokens
	rtotal int      // total number of spaces needed to print stream

	bytes.Buffer
	indents []int
	width   int // remaining space
}

func (p *printer) string(str string) {
	tok := &token{kind: 's', str: str, size: len(str)}
	if len(p.stack) == 0 {
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal += len(str)
	}
}
func (p *printer) pop() (top *token) {
	last := len(p.stack) - 1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}
func (p *printer) push(tok *token) {
	p.stack = append(p.stack, tok)
}
func (p *printer) begin() {
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t := &token{kind: '(', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.push(t)
	p.string("(")
}
func (p *printer) end() {
	p.string(")")
	p.tokens = append(p.tokens, &token{kind: ')'})
	x := p.pop()
	x.size += p.rtotal
	// If x.kind is ' ', then it must be a space between end() and begin().
	// Otherwised, it must be the corresponding "begin"
	if x.kind == ' ' {
		p.pop().size += p.rtotal // corresponding "begin"
	}
	if len(p.stack) == 0 { // this must be the final end()
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}
func (p *printer) space() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == ' ' { // remove the previous space
		x.size += p.rtotal
		p.pop()
	}
	t := &token{kind: ' ', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.push(t)
	p.rtotal++
}

//+ Exercise 12.4
func (p *printer) newline() {
	t := &token{kind: '\n', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.rtotal++ // ??? Not sure how rtotal is used
	fmt.Printf("** rtotal = %d\n", p.rtotal)
}

//- Exercise 12.4

func (p *printer) print(t *token) {
	switch t.kind {
	case 's':
		p.WriteString(t.str)
		p.width -= len(t.str)
	case '(':
		p.indents = append(p.indents, p.width)
	case ')':
		p.indents = p.indents[:len(p.indents)-1] // pop
	case ' ':
		if t.size > p.width {
			p.width = p.indents[len(p.indents)-1] - 1
			fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		} else {
			p.WriteByte(' ')
			p.width--
		}
	//+ Exericse 12.4
	case '\n':
		p.width = p.indents[len(p.indents)-1] - 1
		fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		//- Exercise 12.4
	}
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice: // (value ...)
		p.begin()
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.space()
			}
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
		}
		p.end()

	case reflect.Struct: // ((name value ...)
		p.begin()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				p.newline() // Exercise 12.4
			}
			p.begin()
			p.string(v.Type().Field(i).Name)
			p.space()
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	case reflect.Map: // ((key value ...)
		p.begin()
		for i, key := range v.MapKeys() {
			if i > 0 {
				p.space()
			}
			p.begin()
			if err := pretty(p, key); err != nil {
				return err
			}
			p.space()
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	case reflect.Ptr:
		return pretty(p, v.Elem())

	//+ Exercise 12.3
	case reflect.Bool:
		if v.Bool() {
			p.string("t")
		} else {
			p.string("nil")
		}

	case reflect.Float32, reflect.Float64:
		p.stringf("%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		p.stringf("#C(%f %f)", real(c), imag(c))
	//- Exercise 12.3

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
