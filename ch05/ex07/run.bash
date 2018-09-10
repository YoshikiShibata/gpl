#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go test -v 

go build -o prettyprint

./prettyprint http://gopl.io

# clean up
rm ./prettyprint
