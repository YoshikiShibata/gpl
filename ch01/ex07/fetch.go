// Copyright © 2015 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
		os.Exit(1)
	}
}
