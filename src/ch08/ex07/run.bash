#!/bin/bash

go build -o mirror mirror.go
if [ $? != 0 ]
then
    exit 1
fi

rm -fr www.gopl.io
./mirror http://gopl.io


