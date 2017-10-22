#!/bin/bash 

echo "Mandelbrot with one goroutine"
go build -o mandelbrot_nogoroutine mandelbrot_nogoroutine.go
time ./mandelbrot_nogoroutine > madelbrot_nogoroutine.png

echo ""
echo "Madelbrot with multiple goroutines"
go build -o mandelbrot_goroutine mandelbrot_goroutine.go
time ./mandelbrot_goroutine > madelbrot_goroutine.png

echo ""
echo "Surface with one goroutine"
go build -o surface_nogoroutine surface_nogoroutine.go
time ./surface_nogoroutine > surface_nogroutine.svg

echo ""
echo "Surface with multiple goroutines"
go build -o surface_goroutine surface_goroutine.go
time ./surface_goroutine > surface_goroutine.svg
