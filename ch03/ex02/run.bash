#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o surface
./surface -type=0 > surface0.html
./surface -type=1 > surface1.html
./surface -type=2 > surface2.html
./surface -type=3 > surface3.html
