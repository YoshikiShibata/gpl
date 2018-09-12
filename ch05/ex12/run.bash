#!/bin/bash -x

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build -o outline

if [ $? == 0 ]; then
    ./outline http://gopl.io
    rm outline
fi
