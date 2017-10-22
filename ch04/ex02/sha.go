// Copyright © 2016 Yoshiki Shibata

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var shaType = flag.String("sha", "256", `"256" "384" "512"`)

func main() {
	flag.Parse()

	switch *shaType {
	case "256", "384", "512":
		// do nothing
	default:
		fmt.Printf("invalid sha: %s\n", *shaType)
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("SHA type is %s\n", *shaType)

	reader := bufio.NewReader(os.Stdin)

	for {
		bytes, err := reader.ReadBytes('\n')

		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			fmt.Printf("read failed: %v", err)
			os.Exit(1)
		}

		data := bytes[0 : len(bytes)-1]
		switch *shaType {
		case "256":
			sum := sha256.Sum256(data)
			fmt.Printf("%x\n", sum)
		case "384":
			sum := sha512.Sum384(data)
			fmt.Printf("%x\n", sum)
		case "512":
			sum := sha512.Sum512(data)
			fmt.Printf("%x\n", sum)
		}
	}
}
