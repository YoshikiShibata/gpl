// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("top.tar")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := tar.NewReader(f)

	for {
		h, err := r.Next()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Printf("Name = %s\n", h.Name)
		fmt.Printf("Name = %s\n", h.FileInfo().Name())
		fmt.Printf("Mode = %x\n", uint32(h.FileInfo().Mode()))
		fmt.Printf("Mode = %v\n", h.FileInfo().Mode())
		if h.FileInfo().IsDir() {
			fmt.Printf("Directory\n")
		} else {
			fmt.Printf("Plain file\n")
			_, err = io.Copy(os.Stdout, r)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
