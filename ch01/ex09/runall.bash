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

go build -o fetch fetch.go
if [ $? != 0 ]
then 
	exit 1
fi
./fetch gopl.io
rm fetch
