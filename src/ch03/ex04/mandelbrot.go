package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v + v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{0} // black
}

func main() {
	const (
		width  = 1024
		height = 1024
		scale  = 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := scale * (float64(py)/height - 0.5)
		for px := 0; px < width; px++ {
			x := scale * (float64(px)/width - 0.5)
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}
