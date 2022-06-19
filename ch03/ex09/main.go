// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width  = 1024
	height = 1024
)

type parameters struct {
	renderer func(complex128) color.Color
	w        http.ResponseWriter

	xmin float64
	ymin float64
	xmax float64
	ymax float64

	xstart int
	ystart int
}

func newParameters(w http.ResponseWriter) *parameters {
	var p parameters

	p.renderer = mandelbrot
	p.w = w

	p.xmin = -2
	p.ymin = -2
	p.xmax = +2
	p.ymax = +2

	p.xstart = 0
	p.ystart = 0

	return &p
}

func (p *parameters) setFractal(values url.Values) error {
	const key = "fractal"

	if !containsKey(values, key) {
		return nil
	}

	fractal := values.Get(key)
	switch fractal {
	case "mandelbrot":
		p.renderer = mandelbrot
	case "newton":
		p.renderer = newton
	case "acos":
		p.renderer = acos
	case "sqrt":
		p.renderer = sqrt
	default:
		return fmt.Errorf("unknown fractal : %s\n"+
			`supported fractals are "mandelbrot", "newton", "acos", "sqrt"`, fractal)
	}
	return nil
}

func (p *parameters) setZoom(values url.Values) error {
	const key = "zoom"

	if !containsKey(values, key) {
		return nil
	}

	zoom, err := strconv.Atoi(values.Get(key))
	if err != nil {
		return fmt.Errorf("zoom error: %v", err)
	}

	if zoom <= 0 {
		return fmt.Errorf("invalid zoom %d: must be greater than zero", zoom)
	}

	p.xmin *= 100 / float64(zoom)
	p.xmax *= 100 / float64(zoom)
	p.ymin *= 100 / float64(zoom)
	p.ymax *= 100 / float64(zoom)
	return nil
}

func (p *parameters) setX(values url.Values) error {
	const key = "x"

	if !containsKey(values, key) {
		return nil
	}

	x, err := strconv.Atoi(values.Get(key))
	if err != nil {
		return fmt.Errorf("x error: %v", err)
	}

	p.xstart = x
	return nil
}

func (p *parameters) setY(values url.Values) error {
	const key = "y"

	if !containsKey(values, key) {
		return nil
	}

	y, err := strconv.Atoi(values.Get(key))
	if err != nil {
		return fmt.Errorf("y error: %v", err)
	}

	p.ystart = y
	return nil
}

func containsKey(values url.Values, key string) bool {
	return len(values.Get(key)) != 0
}

func (p *parameters) help(values url.Values) error {
	const key = "help"

	_, ok := values[key]
	if !ok {
		return nil
	}
	return fmt.Errorf(`available options: fractal, zoom, x, and y +
fractal: "mandelbrot" (default), "newton", "acos", "sqrt"
   zoom: positive integer, default is 100 
     x : x coordinate of the center, default is 0
     y : y coordinate of the center, default is 0`)
}

func (p *parameters) setOptions(values url.Values) error {
	setters := []func(url.Values) error{
		p.setFractal, p.setZoom, p.setX, p.setY, p.help}
	for _, setter := range setters {
		if err := setter(values); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := newParameters(w)
		err := p.setOptions(r.URL.Query())
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}

		render(p)
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Printf("%x\n", err)
	}
}

func render(p *parameters) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py+p.ystart)/float64(height)*(p.ymax-p.ymin) + p.ymin
		for px := 0; px < width; px++ {
			x := float64(px+p.xstart)/float64(width)*(p.xmax-p.xmin) + p.xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, p.renderer(z))
		}
	}
	png.Encode(p.w, img) // NOTE: ignoring errors
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

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
