#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go list golang.org/x/crypto/ssh
if [ $? != 0 ]
then 
	echo "Installing golang.org/x/crypto/ssh ..."
	go get golang.org/x/crypto/ssh
	echo "Done"
fi

go build -o issue
if [ $? != 0 ]
then 
	exit 1
fi

echo "This script is not implemented Yet because of a missing file"
