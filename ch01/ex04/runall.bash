#!/bin/bash

# Copyright (C) 2020 Yoshiki Shibata

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

go build -o dup2 dup2.go
if [ $? != 0 ]
then 
	exit 1
fi
cp dup2.go dup2.txt
tail -n 10 dup2.go > tail.txt
head -n 10 dup2.go > head.txt
./dup2 dup2.go dup2.txt tail.txt head.txt
rm dup2 *.txt
