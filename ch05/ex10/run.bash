#!/bin/bash +x

# Copyright (C) 2018 Yoshiki Shibata. All rights reserved.

go test -v
if [ $? != 0 ]; then
        exit 1
fi 
go run topo_sort.go
if [ $? != 0 ]; then
        exit 1
fi 
