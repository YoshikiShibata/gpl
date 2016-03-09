#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o findElement
if [ $? != 0 ]; then
        exit 1
fi 
./findElement http://www.gopl.io/ name
./findElement http://www.gopl.io/ system
