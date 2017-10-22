// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// From: Alan Donovan <adonovan@google.com>
// Date: 2015-12-01 23:25 GMT+09:00
//
// You can do it computing the average of four calls such as:
//
//   mandelbrot(x+xdelta, y+ydelta)
//   mandelbrot(x+xdelta, y-ydelta)
//   mandelbrot(x-xdelta, y+ydelta)
//   mandelbrot(x-xdelta, y-ydelta)
//
// where xdelta is the difference between successive x values in
// the original program.

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dx := float64(1.0) / width * (xmax - xmin)
	dy := float64(1.0) / height * (ymax - ymin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, supersample(x, y, dx, dy))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
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

func supersample(x, y, dx, dy float64) color.Color {
	grays := [...]color.Color{
		mandelbrot(complex(x+dx, y+dy)),
		mandelbrot(complex(x+dx, y-dy)),
		mandelbrot(complex(x-dx, y+dy)),
		mandelbrot(complex(x-dx, y-dy))}
	var sumR, sumG, sumB uint32
	for _, rgba := range grays {
		r, g, b, _ := rgba.RGBA()
		// 16 bits consists of two same 8-bit values
		sumR += r >> 8
		sumG += g >> 8
		sumB += b >> 8
	}
	return color.RGBA{uint8(sumR / 4), uint8(sumG / 4), uint8(sumB / 4), 0xff}
}
