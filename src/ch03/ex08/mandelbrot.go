// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"runtime"
	"sync"

	mc "ch03/ex08/cmplx"
)

func main() {
	commonMain(executeMandelbrot)
}

func executeMandelbrot(w io.Writer) {
	switch *floatType {
	case "complex128":
		mainComplex128(w)
		/*
			case "complex64":
				mainComplex64(w)
		*/
	case "Float":
		mainFloat(w)
	case "Rat":
		mainRat(w)
	}
}

func mainComplex128(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	for py := 0; py < *height; py++ {
		y := float64(py)/float64(*height)*(ymax-ymin) + ymin
		for px := 0; px < *width; px++ {
			x := float64(px)/float64(*width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mainFloat(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	var wg sync.WaitGroup
	var mutex sync.Mutex

	limiter := make(chan struct{}, runtime.NumCPU()*2)

	for py := 0; py < *height; py++ {
		y := float64(py)/float64(*height)*(ymax-ymin) + ymin
		for px := 0; px < *width; px++ {
			limiter <- struct{}{}
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/float64(*width)*(xmax-xmin) + xmin
				z := mc.NewFloatComplex(x, y)
				// Image point (px, py) represents complex value z.
				mz := mandelbrotFloat(z)
				mutex.Lock()
				img.Set(px, py, mz)
				mutex.Unlock()
				<-limiter
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(w, img) // NOTE: ignoring errors
}

func mainRat(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	var wg sync.WaitGroup
	var mutex sync.Mutex

	limiter := make(chan struct{}, runtime.NumCPU())

	for py := 0; py < *height; py++ {
		fmt.Printf("py = %d\n", py)
		y := float64(py)/float64(*height)*(ymax-ymin) + ymin
		for px := 0; px < *width; px++ {
			fmt.Printf("\tpx = %d\n", px)
			limiter <- struct{}{}
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/float64(*width)*(xmax-xmin) + xmin
				z := mc.NewRatComplexM(x, y)
				// Image point (px, py) represents complex value z.
				mz := mandelbrotRat(z)
				mutex.Lock()
				img.Set(px, py, mz)
				mutex.Unlock()
				<-limiter
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotFloat(z *mc.FloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	v := mc.NewFloatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(z *mc.RatComplexM) color.Color {
	const iterations = 200
	const contrast = 15

	v := mc.NewRatComplexM(0, 0)
	for n := uint8(0); n < iterations; n++ {
		fmt.Printf("\t\tn: %d\n", n)
		// v = v*v + z
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			fmt.Printf("\t\t\tGray[%d]\n", 255-contrast*n)
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
