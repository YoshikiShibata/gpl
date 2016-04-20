// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"ch07/ex06/tempflag"
	"flag"
	"fmt"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
