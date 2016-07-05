#!/bin/bash 

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch3/mandelbrot
go build -o imageConverter main.go
if [ $? != 0 ]
then
	exit 1
fi

echo "create mandelbrot.png mandelbrot.jpg mandelbrot.gif"
./mandelbrot > mandelbrot.png
./imageConverter -jpg < mandelbrot.png > mandelbrot.jpg
if [ $? != 0 ]
then
	echo "imageConverter -jpg failed"
	exit 1
fi

./imageConverter -gif < mandelbrot.png > mandelbrot.gif
if [ $? != 0 ]
then
	echo "imageConverter -gif failed"
	exit 1
fi

echo "converts an JPG into PNG and GIF"
./imageConverter -png < mandelbrot.jpg > mandelbrot2.png
if [ $? != 0 ]
then
	echo "imageConverter -png failed"
	exit 1
fi
./imageConverter -png < mandelbrot.jpg > mandelbrot3.png
if [ $? != 0 ]
then
	echo "imageConverter -png failed"
	exit 1
fi

echo "converts an GIF into PNG and JPG"
./imageConverter -jpg < mandelbrot.gif > mandelbrot2.jpg
if [ $? != 0 ]
then
	echo "imageConverter -jpg failed"
	exit 1
fi
./imageConverter -jpg < mandelbrot.gif > mandelbrot3.jpg
if [ $? != 0 ]
then
	echo "imageConverter -jpg failed"
	exit 1
fi
