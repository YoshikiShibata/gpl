#!/bin/bash

# Copyright (C) 2017 Yoshiki Shibata

go build -o echo echo.go
if [ $? != 0 ]
then 
	exit 1
fi
./echo Hello World
