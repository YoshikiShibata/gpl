// Copyright © 2016 Yoshiki Shibata

package main

import "fmt"

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("%v\n", data)
	RotateCycles(data, 7)
	fmt.Printf("%v\n", data)
}

//
// This implementation is based on the Section 10.4 Rotate Algorithm
// from "Elements of Programming"
//
func cycleTo(i int, f func(int) int, exchangeValues func(int, int)) {
	k := f(i)
	for k != i {
		exchangeValues(i, k)
		k = f(k)
	}
	return
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// RotateCycles rotate all cycles so that a given slice is rotated right by
// the specified size.
func RotateCycles(data []int, rotateRightSize int) {
	if rotateRightSize <= 0 {
		panic(fmt.Sprintf("rotateSize is %d, must be greater than 0", rotateRightSize))
	}
	rotateRightSize %= len(data)

	n := gcd(len(data), rotateRightSize)
	for i := 0; i < n; i++ {
		cycleTo(i,
			func(i int) int {
				return (i + rotateRightSize) % len(data)
			},
			func(i, j int) {
				data[i], data[j] = data[j], data[i]
			})
	}
}
