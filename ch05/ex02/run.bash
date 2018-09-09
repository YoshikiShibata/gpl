#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
go build tagcount.go
./fetch https://golang.org | ./tagcount

# clean
rm fetch tagcount
