#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go test -v 
if [ $? != 0 ]; then
        exit 1
fi 

go build -o prettyprint
if [ $? != 0 ]; then
        exit 1
fi 

./prettyprint http://gopl.io

# clean up
rm ./prettyprint
