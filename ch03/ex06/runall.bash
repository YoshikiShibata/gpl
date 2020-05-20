#!/bin/bash

# Copyright (C) 2020 Yoshiki Shibata. All rights reserved.

golint
if [ $? != 0 ]
then 
	exit 1
fi

go run mandelbrot.go supersample.go > mandelbrot_1.png
if [ $? != 0 ]
then 
	exit 1
fi

go run main.go supersample.go  > mandelbrot_2.png
if [ $? != 0 ]
then 
	exit 1
fi
open -a Safari mandelbrot_1.png
open -a Safari mandelbrot_2.png

sleep 5
rm *.png
