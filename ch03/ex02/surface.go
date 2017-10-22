// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015, 2016 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var f func(float64, float64) float64 = f0

var funcType = flag.String("type", "", "select function")

func main() {
	flag.Parse()
	selectFunction(*funcType)

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func selectFunction(fType string) {
	switch fType {
	case "1":
		f = f1
	case "2":
		f = f2
	case "3":
		f = f3
	default:
		f = f0
	}
}

func isFinite(f float64) bool {
	return !math.IsInf(f, 0)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f0(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(-x) * math.Pow(1.5, -r)
}

func f2(x, y float64) float64 {
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}

func f3(x, y float64) float64 {
	return math.Sin(x*y/10) / 10
}
