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

go build -o lissajous lissajous.go
if [ $? != 0 ]
then 
	exit 1
fi
./lissajous > lissajous.gif
open -a safari lissajous.gif
sleep 2
rm lissajous lissajous.gif
