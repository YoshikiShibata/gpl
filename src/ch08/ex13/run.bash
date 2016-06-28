#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o chat
go build gopl.io/ch8/netcat3

echo "Run netcat3 command in another window"
./chat
