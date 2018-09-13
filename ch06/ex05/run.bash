#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

export GOARCH=amd64 
echo "GOARCH = " $GOARCH
go test -v

export GOARCH=386 
echo "GOARCH = " $GOARCH
go test -v
