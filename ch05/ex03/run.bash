#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
if [ $? != 0 ]; then
        exit 1
fi 
go build text.go
if [ $? != 0 ]; then
        exit 1
fi 
./fetch https://golang.org | ./text

# clean
rm fetch text
