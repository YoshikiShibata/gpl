// Copyright © 2015 Yoshiki Shibata. All rights reserved.

// uc vonerts its numeric argument to various units
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ch02/ex02/unitconv"
)

func main() {
	for _, arg := range getArgs() {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		showTemperature(t)
		showLength(t)
		showWeight(t)
		fmt.Println()
	}
}

func getArgs() []string {
	args := []string{}
	if len(os.Args) == 1 {
		// Inputs must be terminated by CTL-D
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			for _, arg := range strings.Split(input.Text(), " ") {
				args = append(args, arg)
			}
		}
	} else {
		args = os.Args[1:]
	}
	return args
}

func showTemperature(t float64) {
	f := unitconv.Fahrenheit(t)
	c := unitconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, unitconv.FToC(f), c, unitconv.CToF(c))
}

func showLength(t float64) {
	f := unitconv.Feet(t)
	m := unitconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, unitconv.FeetToMeter(f), m, unitconv.MeterToFeet(m))
}

func showWeight(t float64) {
	p := unitconv.Pound(t)
	k := unitconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n",
		p, unitconv.PoundToKilogram(p), k, unitconv.KilogramToPound(k))
}
