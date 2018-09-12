#/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

rm -fr cached

go run crawl.go http://golang.org
