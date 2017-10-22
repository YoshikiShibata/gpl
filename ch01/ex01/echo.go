// Copyright © 2015 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
