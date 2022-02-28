// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, err := zip.OpenReader("top.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		fmt.Printf("Contents of %s:\n", f.FileInfo().Name())
		if f.Mode().IsDir() {
			fmt.Printf("Directory\n")
		} else {
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.CopyN(os.Stdout, rc, 1024*1024*256)
			if err != nil {
				log.Fatal(err)
			}
			rc.Close()
		}
		fmt.Println()

	}
}
