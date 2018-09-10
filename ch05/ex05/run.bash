#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build -o wai
if [ $? != 0 ]; then
        exit 1
fi 

./wai https://golang.org \
		https://www.ricoh.co.jp \
		https://yshibata.blog.so-net.ne.jp\
		https://jp.merpay.com

# clean

rm wai
