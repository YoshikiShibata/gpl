#!/bin/bash

# Copyright (C) 2020 Yoshiki Shibata. All rights reserved.

go vet
if [ $? != 0 ]
then 
	exit 1
fi

golint
if [ $? != 0 ]
then 
	exit 1
fi

go run mandelbrot.go > mandelbrot.png
if [ $? != 0 ]
then 
	exit 1
fi

open -a Safari mandelbrot.png

sleep 5
rm mandelbrot.png
