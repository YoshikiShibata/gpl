// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"./eval"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("type in an expression: ")
		expression, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			fmt.Println()
			return
		}
		expression = expression[:len(expression)-1]

		fmt.Printf("Expression = %s\n", expression)
		expr, err := eval.Parse(expression)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%g\n", expr.Eval(eval.Env{}))
		}
	}
}
