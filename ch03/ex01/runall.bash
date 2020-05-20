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

go run surface.go > output.svg
if [ $? != 0 ]
then 
	exit 1
fi

open -a Safari output.svg
sleep 5
rm output.svg
