#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

export GOARCH=amd64 
echo "GOARCH = " $GOARCH
go test

export GOARCH=386 
echo "GOARCH = " $GOARCH
go test
