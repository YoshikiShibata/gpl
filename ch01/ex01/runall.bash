#!/bin/bash

# Copyright (C) 2017, 2020 Yoshiki Shibata

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

go build -o echo echo.go
if [ $? != 0 ]
then 
	exit 1
fi
./echo Hello World
rm echo
