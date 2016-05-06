#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o xmlselect
go build gopl.io/ch1/fetch
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div div h2

echo ""
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div h2

echo ""
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div class="toc" h2
