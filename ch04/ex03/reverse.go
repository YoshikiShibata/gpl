// Copyright © 2016 Yoshiki Shibata

package main

const Size = 10 // supported the size of an array

func reverse(a *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
