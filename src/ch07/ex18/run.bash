#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o a.out
go build gopl.io/ch1/fetch
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 > w3.xml
./a.out < w3.xml > w3.xmltree
