#!/bin/bash 

echo ""
echo "Madelbrot with multiple goroutines"
go build -o mandelbrot mandelbrot.go

echo ""
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
time ./mandelbrot > madelbrot.1.png

echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
time ./mandelbrot > madelbrot.2.png

echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
time ./mandelbrot > madelbrot.4.png

echo ""
export GOMAXPROCS=8
echo "GOMAXPROCS=${GOMAXPROCS}"
time ./mandelbrot > madelbrot.8.png

echo ""
export GOMAXPROCS=16
echo "GOMAXPROCS=${GOMAXPROCS}"
time ./mandelbrot > madelbrot.16.png
