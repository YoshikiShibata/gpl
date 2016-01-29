// Copyright © 2016 Yoshiki Shibata

package main

import (
	"fmt"
	"os"
)

func main() {
	argsLen := len(os.Args)
	if argsLen == 1 {
		fmt.Fprintf(os.Stderr, "Movie is not specified\n")
		fmt.Fprintf(os.Stderr, "poster <movie titles>\n")
		os.Exit(1)
	}

	for _, title := range os.Args[1:] {
		movie, err := getMovie(title)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", title, err)
			continue
		}

		fmt.Printf("%v\n", movie)
		if movie.Poster == "N/A" {
			fmt.Printf("Sorry, the poster for %s is not available\n", title)
			continue
		}

		fmt.Printf("Fetching the poster for %s ... ", title)
		file, err := fetchJPEG(title, movie.Poster)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		fmt.Printf("Saved as %s\n", file)
	}
}
