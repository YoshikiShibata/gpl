package main

import (
	"ch07/ex03/tempflag"
	"flag"
	"fmt"
)

var tempC = tempflag.CelsiusFlag("temp_c", 20.0, "the temperature")
var tempF = tempflag.FahrenheitFlag("temp_f", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*tempC)
	fmt.Println(*tempF)
	fmt.Println("Oop! Kelvin must be implemented")
}
