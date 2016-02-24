#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
go build text.go
./fetch https://golang.org | ./text
