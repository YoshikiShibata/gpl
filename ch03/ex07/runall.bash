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

go run newton.go > newton.png
if [ $? != 0 ]
then 
	exit 1
fi

open -a Safari newton.png

sleep 5
rm *.png
