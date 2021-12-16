// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"flag"
	"fmt"

	"github.com/YoshikiShibata/gpl/ch07/ex06/tempflag"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
