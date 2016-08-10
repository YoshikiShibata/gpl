#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go test -bench=.

# BenchmarkAdd_IntSet-8                  50000         51907 ns/op
# BenchmarkAdd_MapIntSet-8               10000        183672 ns/op
# BenchmarkUnitonWith_IntSet-8           20000         82875 ns/op
# BenchmarkUnionWith_MapIntSet-8          3000        437941 ns/op
