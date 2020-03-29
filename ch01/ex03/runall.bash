#!/bin/bash

# Copyright (C) 2020 Yoshiki Shibata

go vet github.com/YoshikiShibata/gpl/ch01/ex03/concat
if [ $? != 0 ]
then 
	exit 1
fi

golint github.com/YoshikiShibata/gpl/ch01/ex03/concat
if [ $? != 0 ]
then 
	exit 1
fi

go test -bench=. github.com/YoshikiShibata/gpl/ch01/ex03/concat
