// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "fmt"

func main() {
	fmt.Printf("square(2) = %d\n", square(2))
}

func square(n int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()

	panic(n * n)
}
