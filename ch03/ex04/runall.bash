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

go build -o surface.server surface.go 
if [ $? != 0 ]
then 
	exit 1
fi

killall surface.server
./surface.server &

sleep 5

open -a Safari http://localhost:8000/?help 
