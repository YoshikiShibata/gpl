#/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

rm -fr cached

go run crawl.go http://golang.org
