#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata All rights reserved.

go build -o githubserver
if [ $? != 0 ]
then
    exit 1
fi

echo "Please specify a repository with a brower such as"
echo ""
echo "  localhost:8000/golang/go/"
echo ""
echo "Don't forget the last slash!"
echo ""
echo "The Next operation is not implemented yet. So only the first page"
echo "is shown for Issues and Milestones"
echo ""
echo "Enjoy!"

./githubserver
