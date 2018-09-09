#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
go build findlinks1.go
./fetch https://golang.org | ./findlinks1

# clean
rm fetch findlinks1
