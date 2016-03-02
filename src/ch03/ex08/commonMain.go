package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	mc "ch03/ex08/cmplx"
)

var floatType = flag.String("type", "complex128",
	"arithmetic type: complex64, complex128, Float, Rat")
var zoom = flag.Int("zoom", 100, "zoom percent")
var precision = flag.Uint("precision", 0, "precision for Float")
var output = flag.String("output", "", "output png to file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var width = flag.Int("w", 1024, "width")
var height = flag.Int("h", 1024, "height")

var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
var zoomFactor float64

func commonMain(f func(io.Writer)) {
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
	f(w)
	end := time.Now()
	fmt.Fprintf(os.Stderr, "Duration := %v\n", end.Sub(start))
}

func showHeader() {
	title := "unknown"
	switch *floatType {
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

func validateParams() {
	args := flag.Args()
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "invalid arguments: %v\n\n", args)
		showUsage()
	}
	switch *floatType {
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
