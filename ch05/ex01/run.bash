#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
go build findlinks1.go
./fetch https://golang.org | ./findlinks1
