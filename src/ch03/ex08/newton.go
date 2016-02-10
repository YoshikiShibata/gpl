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
	"runtime"
	"sync"

	mc "ch03/ex08/cmplx"
)

var aType = flag.String("type", "complex128",
	"arithmetic type: complex64, complex128, float, rat")
var zoom = flag.Int("zoom", 100, "zoom percent")

const usage = `usage: newton [-type=arithemeticType] [-zoom=percent]
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
	case "float":
		mainFloat()
	}
}

func mainComplex64() {
	fmt.Fprintf(os.Stderr, "=== complex64 ===\n")
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
	fmt.Fprintf(os.Stderr, "=== complex128 ===\n")
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

func mainFloat() {
	fmt.Fprintf(os.Stderr, "=== Float ===\n")
	fmt.Fprintf(os.Stderr, "factor=%g\n", zoomFactor)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	var mutex sync.Mutex

	limiter := make(chan struct{}, runtime.NumCPU()*2)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			limiter <- struct{}{}
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := mc.NewFloatComplex(x, y)
				// Image point (px, py) represents complex value z.
				nwz := newtonFloat(z)
				mutex.Lock()
				img.Set(px, py, nwz)
				mutex.Unlock()
				<-limiter
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func validateParams() {
	args := flag.Args()
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "invalid arguments: %v\n\n", args)
		showUsage()
	}
	switch *aType {
	case "complex64", "complex128", "float", "rat":
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

func newtonFloat(z *mc.FloatComplex) (c color.Color) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintf(os.Stderr, "Ignore: %v\n", x)
			c = color.Black
		}
	}()

	const iterations = 37
	const contrast = 7
	c1 := mc.NewFloatComplex(1.0, 0)
	c4 := mc.NewFloatComplex(4.0, 0)
	for i := uint8(0); i < iterations; i++ {
		// z -= (z - 1/(z*z*z)) / 4
		z = z.Sub(z.Sub(c1.Quo(z.Mul(z).Mul(z))).Quo(c4))
		// if cmplx.Abs(z*z*z*z-1) < 1e-6 {
		if z.Mul(z).Mul(z).Mul(z).Sub(c1).Abs() < 1e-6 {
			return color.RGBA{255 - contrast*i, contrast * i, 0, 0xff}
		}
	}
	return color.Black
}
