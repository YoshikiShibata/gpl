#!/bin/bash

# Copyright (C) 2016, 2021 Yoshiki Shibata. All rights reserved.

go test -bench=.

# Now macOS doesn't support 32 bits
# echo "32 Bit version"
# GOARCH=386 go test -bench=. bench_test.go intset.go mapintset.go

# BenchmarkAdd_IntSet-8                  50000         51907 ns/op
# BenchmarkAdd_MapIntSet-8               10000        183672 ns/op
# BenchmarkUnitonWith_IntSet-8           20000         82875 ns/op
# BenchmarkUnionWith_MapIntSet-8          3000        437941 ns/op

# 2017.01.27 Go 1.8 (almost)
# BenchmarkAdd_IntSet-4            	   50000	     37298 ns/op
# BenchmarkAdd_MapIntSet-4         	   10000	    133083 ns/op
# BenchmarkUnitonWith_IntSet-4     	   20000	     63415 ns/op
# BenchmarkUnionWith_MapIntSet-4   	    5000	    372974 ns/op
# PASS
# ok  	ch11/ex07	7.423s
# 32 Bit version
# BenchmarkAdd_IntSet-4            	   30000	     40636 ns/op
# BenchmarkAdd_MapIntSet-4         	   10000	    179372 ns/op
# BenchmarkUnitonWith_IntSet-4     	   20000	     68477 ns/op
# BenchmarkUnionWith_MapIntSet-4   	    3000	    462411 ns/op
# PASS
# ok  	command-line-arguments	6.977s
