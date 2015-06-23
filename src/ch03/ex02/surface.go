package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)

	maxHeight = 1.0
	minHeight = -0.22

	red  = 0x00ff0000
	blue = 0x000000ff
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // computer surface height z

	// project (x,y,z) isometrically onto 2-D SV canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func color(h1, h2, h3, h4 float64) string {
	height := (h1 + h2 + h3 + h4) / 4
	if height > maxHeight || minHeight > height {
		panic(fmt.Sprintf("illegal height : %g", height))
	}

	delta := uint32((maxHeight - height) / (maxHeight - minHeight) * 255)

	c := (red - delta<<16) + delta
	return prependZeros(fmt.Sprintf("%X", c))
}

func prependZeros(hex string) string {
	result := hex
	for i := len(hex); i < 6; i++ {
		result = "0" + result
	}
	return result
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, h1 := corner(i+1, j)
			bx, by, h2 := corner(i, j)
			cx, cy, h3 := corner(i, j+1)
			dx, dy, h4 := corner(i+1, j+2)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(h1, h2, h3, h4))
		}
	}
	fmt.Println("</svg>")
}
