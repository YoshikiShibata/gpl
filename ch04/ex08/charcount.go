// Copyright (C) 2015 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Control = "Control"
	Digit   = "Digit"
	Graphic = "Graphic"
	Letter  = "Letter"
	Lower   = "Lower"
	Mark    = "Mark"
	Number  = "Number"
	Print   = "Print"
	Punct   = "Punctuation"
	Space   = "Space"
	Symbol  = "Symbol"
	Title   = "Title"
	Upper   = "Upper"
)

var isFunctions = map[string]func(rune) bool{
	Control: unicode.IsControl,
	Digit:   unicode.IsDigit,
	Graphic: unicode.IsGraphic,
	Letter:  unicode.IsLetter,
	Lower:   unicode.IsLower,
	Mark:    unicode.IsMark,
	Number:  unicode.IsNumber,
	Print:   unicode.IsPrint,
	Punct:   unicode.IsPunct,
	Space:   unicode.IsSpace,
	Symbol:  unicode.IsSymbol,
	Title:   unicode.IsTitle,
	Upper:   unicode.IsUpper,
}

func main() {
	counts := make(map[rune]int)       // counts of Unicode chars
	categories := make(map[string]int) // counts of categories
	var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
	invalid := 0                       // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		c, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if c == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[c]++
		utflen[n]++

		for k, f := range isFunctions {
			if f(c) {
				categories[k]++
			}
		}
	}

	fmt.Println("rune\tcount")
	for c, n := range counts {
		fmt.Printf("%c\t%d\n", c, n)
	}

	fmt.Println()
	fmt.Println("len\tcount")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Println()
	fmt.Println("   category\tcount")
	for k, n := range categories {
		fmt.Printf("%11s\t%d\n", k, n)
	}
}
