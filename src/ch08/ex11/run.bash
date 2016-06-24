#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o fetch
./fetch https://golang.org http://gopl.io https://godoc.org http://www.amazon.com
