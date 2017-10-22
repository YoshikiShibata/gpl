#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build gopl.io/ch1/fetch
go build outline.go
./fetch https://golang.org | ./outline
