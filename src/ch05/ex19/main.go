// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "fmt"

func main() {
	fmt.Printf("sqrt(2) = %d\n", sqrt(2))
}

func sqrt(n int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()

	panic(n * n)
}
