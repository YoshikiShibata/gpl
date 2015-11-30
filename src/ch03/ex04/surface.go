// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	angle = math.Pi / 6 // angle of x, y axes (=30°)

	maxHeight = 1.0
	minHeight = -0.22

	redShift   = 16
	greenShift = 8
	blueShift  = 0
)

type constants struct {
	sin30            float64
	cos30            float64
	width            int
	height           int
	cells            int
	xyrange          float64
	xyscale          float64
	zscale           float64
	topColorShift    uint
	bottomColorShift uint
	w                http.ResponseWriter
}

func newConstants(w http.ResponseWriter) *constants {
	var c constants

	c.w = w

	c.sin30 = math.Sin(angle) // sin(30°)
	c.cos30 = math.Cos(angle) // cos(30°)
	c.width = 600             // width of the canvas in pixels
	c.height = 320            // height of the canvas in pixels
	c.cells = 100             // number of grid cells
	c.xyrange = 30.0          // x, y axis range (-xyrange..+xyrange)
	c.topColorShift = redShift
	c.bottomColorShift = blueShift
	c.computeScales()
	return &c
}

func (c *constants) computeScales() {
	c.xyscale = float64(c.width) / 2 / c.xyrange // pixels per x or y unit
	c.zscale = float64(c.height) * 0.4           // pixels per z unit
}

func (c *constants) setWidth(values url.Values) error {
	key := "width"

	if !containsKey(values, key) {
		return nil
	}

	width, err := c.extractPositiveIntValue(values, key)
	if err != nil {
		return err
	}
	c.width = width
	c.computeScales()
	return nil
}

func (c *constants) setHeight(values url.Values) error {
	key := "height"

	if !containsKey(values, key) {
		return nil
	}

	height, err := c.extractPositiveIntValue(values, key)
	if err != nil {
		return err
	}
	c.height = height
	c.computeScales()
	return nil
}

func (c *constants) setTopColor(values url.Values) error {
	key := "topColor"

	if !containsKey(values, key) {
		return nil
	}

	color := values.Get(key)
	shift, err := toColorShift(color)
	c.topColorShift = shift
	if err != nil {
		return err
	}
	return nil
}

func (c *constants) setBottomColor(values url.Values) error {
	key := "bottomColor"

	if !containsKey(values, key) {
		return nil
	}

	color := values.Get(key)
	shift, err := toColorShift(color)
	c.bottomColorShift = shift
	if err != nil {
		return err
	}
	return nil
}

func (c *constants) help(values url.Values) error {
	key := "help"

	_, ok := values[key]
	if !ok {
		return nil
	}
	return fmt.Errorf("available options: topColor, bottomColor, width, height\n" +
		"available colors: RED, GREEN, BLUE\n" +
		"default values: topColor=RED, bottomColor=BLUE, width=320, height=600")
}

func toColorShift(color string) (uint, error) {
	switch color {
	case "red", "RED":
		return redShift, nil
	case "green", "GREEN":
		return greenShift, nil
	case "blue", "BLUE":
		return blueShift, nil
	}
	return 0, errors.New(fmt.Sprintf("unknown color: %s", color))
}

func containsKey(values url.Values, key string) bool {
	return len(values.Get(key)) != 0
}

func (c *constants) extractPositiveIntValue(values url.Values, key string) (int, error) {
	value, err := strconv.Atoi(values.Get(key))

	if err != nil {
		fmt.Fprintf(c.w, "%v", err)
		return -1, err
	}

	if value <= 0 {
		err := errors.New(fmt.Sprintf("invalid %s: %d", key, value))
		fmt.Fprintf(c.w, "%v", err)
		return -1, err
	}
	return value, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func (c *constants) corner(i, j int) (float64, float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := c.xyrange * (float64(i)/float64(c.cells) - 0.5)
	y := c.xyrange * (float64(j)/float64(c.cells) - 0.5)

	z := f(x, y) // computer surface height z

	// project (x,y,z) isometrically onto 2-D SV canvas (sx,sy)
	sx := float64(c.width)/2 + (x-y)*c.cos30*c.xyscale
	sy := float64(c.height)/2 + (x+y)*c.sin30*c.xyscale - z*c.zscale
	return sx, sy, z
}

func (c *constants) color(h1, h2, h3, h4 float64) string {
	height := (h1 + h2 + h3 + h4) / 4
	if height > maxHeight || minHeight > height {
		panic(fmt.Sprintf("illegal height : %g", height))
	}

	delta := uint32((maxHeight - height) / (maxHeight - minHeight) * 255)

	color := (0xff-delta)<<c.topColorShift + delta<<c.bottomColorShift
	return prependZeros(fmt.Sprintf("%X", color))
}

func prependZeros(hex string) string {
	result := hex
	for i := len(hex); i < 6; i++ {
		result = "0" + result
	}
	return result
}

func surface(out io.Writer, c *constants) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", c.width, c.height)
	for i := 0; i < c.cells; i++ {
		for j := 0; j < c.cells; j++ {
			ax, ay, h1 := c.corner(i+1, j)
			bx, by, h2 := c.corner(i, j)
			cx, cy, h3 := c.corner(i, j+1)
			dx, dy, h4 := c.corner(i+1, j+2)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.color(h1, h2, h3, h4))
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func (c *constants) setOptions(values url.Values) error {
	setters := []func(url.Values) error{
		c.setWidth, c.setHeight, c.setTopColor, c.setBottomColor, c.help}
	for _, setter := range setters {
		if err := setter(values); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		c := newConstants(w)
		err := c.setOptions(r.URL.Query())
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}

		w.Header().Add("Content-Type", "image/svg+xml")
		surface(w, c)
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
