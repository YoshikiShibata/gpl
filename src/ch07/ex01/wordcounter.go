// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	nBytes := len(p)

	for {
		advance, token, err := bufio.ScanWords(p, true)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}

		if token != nil {
			*c += 1
		}

		p = p[advance:]

		if len(p) == 0 {
			return nBytes, nil
		}
	}
}
