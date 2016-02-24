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
	"io"
	"math/cmplx"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	mc "ch03/ex08/cmplx"
)

var aType = flag.String("type", "complex128",
	"arithmetic type: complex64, complex128, Float, Rat")
var zoom = flag.Int("zoom", 100, "zoom percent")
var precision = flag.Uint("precision", 0, "precision for Float")
var output = flag.String("output", "", "output png to file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

const (
	width, height = 1024, 1024
)

var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
var zoomFactor float64

func main() {
	flag.Parse()
	validateParams()

	showHeader()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var w io.Writer = os.Stderr

	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		w = f
		defer f.Close()
	}

	zoomFactor = 100.0 / float64(*zoom)
	xmin *= zoomFactor
	xmax *= zoomFactor
	ymin *= zoomFactor
	ymax *= zoomFactor

	if *precision > 0 {
		mc.SetPrecision(*precision)
		fmt.Fprintf(os.Stderr, "precision = %d\n", *precision)
	}

	start := time.Now()
	switch *aType {
	case "complex128":
		mainComplex128(w)
	case "complex64":
		mainComplex64(w)
	case "Float":
		mainFloat(w)
	case "Rat":
		mainRat(w)
	}
	end := time.Now()
	fmt.Fprintf(os.Stderr, "Duration := %v\n", end.Sub(start))
}

func showHeader() {
	title := "unknown"
	switch *aType {
	case "complex64":
		title = "complex64"
	case "complex128":
		title = "complex128"
	case "Float":
		title = "big.Float"
	case "Rat":
		title = "big.Rat"
	}

	fmt.Fprintf(os.Stderr, "\n=== %s ===\n", title)
	fmt.Fprintf(os.Stderr, "zoom = %d\n", *zoom)
}

func mainComplex64(w io.Writer) {
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
	png.Encode(w, img) // NOTE: ignoring errors
}

func mainComplex128(w io.Writer) {
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
	png.Encode(w, img) // NOTE: ignoring errors
}

func mainFloat(w io.Writer) {
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
	png.Encode(w, img) // NOTE: ignoring errors
}

func mainRat(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	var mutex sync.Mutex

	limiter := make(chan struct{}, runtime.NumCPU())

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		fmt.Printf("py = %d\n", py)
		for px := 0; px < width; px++ {
			fmt.Printf("    px = %d\n", px)
			limiter <- struct{}{}
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := mc.NewRatComplex(x, y)
				// Image point (px, py) represents complex value z.
				nwz := newtonRat(z)
				mutex.Lock()
				img.Set(px, py, nwz)
				mutex.Unlock()
				<-limiter
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(w, img) // NOTE: ignoring errors
}

func validateParams() {
	args := flag.Args()
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "invalid arguments: %v\n\n", args)
		showUsage()
	}
	switch *aType {
	case "complex64", "complex128", "Float", "Rat":
	default:
		showUsage()
	}

	if *zoom < 0 {
		showUsage()
	}
}

func showUsage() {
	flag.PrintDefaults()
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

func newtonRat(z *mc.RatComplex) (c color.Color) {
	const iterations = 37
	const contrast = 7
	c1 := mc.NewRatComplex(1.0, 0)
	c4 := mc.NewRatComplex(4.0, 0)
	for i := uint8(0); i < iterations; i++ {
		fmt.Printf("\t\ti : %d %v\n", i, time.Now())
		// z -= (z - 1/(z*z*z)) / 4
		z = z.Sub(z.Sub(c1.Quo(z.Mul(z).Mul(z))).Quo(c4))
		fmt.Printf("\t\t\t: %v\n", time.Now())
		// fmt.Printf("\t\t\t%v\n", z)
		// if cmplx.Abs(z*z*z*z-1) < 1e-6 {
		if z.Mul(z).Mul(z).Mul(z).Sub(c1).Abs() < 1e-6 {
			return color.RGBA{255 - contrast*i, contrast * i, 0, 0xff}
		}
		fmt.Printf("\t\t\t: %v\n", time.Now())
	}
	return color.Black
}
