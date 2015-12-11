// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var aType = flag.String("type", "complex128",
	"arithmetic type: complex64, complex128, float, rat")
var zoom = flag.Int("zoom", 100, "zoom percent")

const usage = `usage: newton [-type=arithemeticType] [-res=resolution]
    type: complex64, complex128, float, rat. Default is complex128
    zoom: percent. Default 100%`

const (
	width, height = 1024, 1024
)

var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
var zoomFactor float64

func main() {
	flag.Parse()
	validateParams()
	fmt.Fprintf(os.Stderr, "type = %s\n", *aType)
	fmt.Fprintf(os.Stderr, "zoom = %d\n", *zoom)

	zoomFactor = 100.0 / float64(*zoom)
	xmin *= zoomFactor
	xmax *= zoomFactor
	ymin *= zoomFactor
	ymax *= zoomFactor

	switch *aType {
	case "complex128":
		mainComplex128()
	case "complex64":
		mainComplex64()
	}
}

func mainComplex64() {
	fmt.Fprintf(os.Stderr, "factor=%g\n", zoomFactor)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*float32((ymax-ymin)) + float32(ymin)
		for px := 0; px < width; px++ {
			x := float32(px)/width*float32((xmax-xmin)) + float32(xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton64(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mainComplex128() {
	fmt.Fprintf(os.Stderr, "factor=%g\n", zoomFactor)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton128(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func validateParams() {
	switch *aType {
	case "complex64", "complex128", "floa", "rat":
	default:
		showUsage()
	}

	if *zoom < 0 {
		showUsage()
	}
}

func showUsage() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton64(z complex64) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.RGBA{255 - contrast*i, contrast * i, 0, 0xff}
		}
	}
	return color.Black
}

func newton128(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255 - contrast*i, contrast * i, 0, 0xff}
		}
	}
	return color.Black
}
