// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		showUsageAndExit()
	}
	img, err := readImage(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read an image: %v\n", err)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-jpg":
		err = toJPEG(img, os.Stdout)
	case "-gif":
		err = toGIF(img, os.Stdout)
	case "-png":
		err = toPNG(img, os.Stdout)
	default:
		showUsageAndExit()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "conversion error: %v\n", err)
		os.Exit(1)
	}
}

func showUsageAndExit() {
	fmt.Fprintf(os.Stderr, "imageConvert [-gif | -png | -jpg]\n")
	os.Exit(1)
}

func readImage(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{256, nil, nil})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
