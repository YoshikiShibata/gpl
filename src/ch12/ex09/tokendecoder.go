// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"bytes"
	"text/scanner"
)

// A Token is an interface holding one of the token types:
// Symbol, String, Int, StartList, EndList
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
	switch lex.toekn {
	case scanner.Ident:

	case scanner.String:

	case scanner.Int:

	case '(':

	case ')':
	}
}
