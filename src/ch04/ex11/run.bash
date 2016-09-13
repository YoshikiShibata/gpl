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

echo ""
echo "Create Issue"
./issue -create -title "Test Issue" -body "This is a test issue" YoshikiShibata/gpltest 
if [ $? != 0 ]
then 
    exit 1
fi
ISSUE_NO=`cat issue_no.txt`

echo ""
echo "Edit Issue No.$ISSUE_NO"
./issue -edit -issue $ISSUE_NO -title "Updated Test Issue" -body "Updated Test Issue" YoshikiShibata/gpltest
if [ $? != 0 ]
then 
    exit 1
fi

echo ""
echo "Print Issue No.$ISSUE_NO"
./issue -print -issue $ISSUE_NO YoshikiShibata/gpltest


echo ""
echo "Close Issue No.$ISSUE_NO"
./issue -close -issue $ISSUE_NO YoshikiShibata/gpltest
