#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o listdeps
./listdeps runtime errors unsafe unicode reflect crypto/md5 go/ast
