// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"bytes"
	"fmt"
	"strconv"
	"text/scanner"
)

//+ Exercise 12.9

// A Token is an interface holding one of the token types:
// Symbol, String, Int, StartList, EndList
// It is not hard to support more types such as Bool, Float, Complex, Interface,
// but they are not supported by this version
type Token interface{}

// A Symbol represents a symbol in S-expressions
type Symbol struct {
	Name string
}

// A String represents a string value in S-expressions
type String struct {
	Value string
}

// A Int represents an int value in S-expressions
type Int struct {
	Value int
}

// A StartList represents the start of a list in S-expressions
type StartList struct {
}

// A EndList represetns the end of a list in S-expressions
type EndList struct {
}

type Decoder struct {
	lex *lexer
}

func NewDecoder(data []byte) *Decoder {
	var decoder Decoder
	decoder.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	decoder.lex.scan.Init(bytes.NewReader(data))
	decoder.lex.next()
	return &decoder
}

func (d *Decoder) Token() Token {
	switch d.lex.token {
	case scanner.Ident:
		name := d.lex.text()
		d.lex.next()
		return Symbol{name}

	case scanner.String:
		value, _ := strconv.Unquote(d.lex.text())
		d.lex.next()
		return String{value}

	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text()) // NOTE: ignoring errors
		d.lex.next()
		return Int{i}

	case '(':
		d.lex.next()
		return StartList{}

	case ')':
		d.lex.next()
		return EndList{}

	}
	panic(fmt.Sprintf("token = %d\n", d.lex.token))
}

//- Exercise 12.9
