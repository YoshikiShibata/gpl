#!/bin/bash

go build gopl.io/ch1/fetch
./fetch http://gopl.io/ch1/helloworld?go-get=1
