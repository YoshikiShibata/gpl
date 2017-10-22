#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o chat chat.go
go build -o netcat4 netcat.go

echo "Run netcat4 command in another window"
./chat
