#!/bin/bash -x

# Copyright (C) 2017 Yoshiki Shibata. All rights reserved.

go build -o issues
./issues repo:golang/go is:open json decoder
