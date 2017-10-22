#!/bin/bash

go test -bench=.

# Results
# BenchmarkPopCount_0-4                       	2000000000	         0.40 ns/op
# BenchmarkPopCountbyShifting_0-4             	20000000	       108 ns/op
# BenchmarkPopCountByClearingBit_0-4          	500000000	         3.10 ns/op
# BenchmarkBitCount_0-4                       	2000000000	         0.40 ns/op
# 
# BenchmarkPopCount_FFFF-4                    	2000000000	         0.44 ns/op
# BenchmarkPopCountbyShifting_FFFF-4          	20000000	        97.5 ns/op
# BenchmarkPopCountByClearingBit_FFFF-4       	100000000	        12.4 ns/op
# BenchmarkBitCount_FFFF-4                    	2000000000	         0.41 ns/op
# 
# BenchmarkPopCount_FFFFFFFF-4                	2000000000	         0.39 ns/op
# BenchmarkPopCountbyShifting_FFFFFFFF-4      	20000000	       108 ns/op
# BenchmarkPopCountByClearingBit_FFFFFFFF-4   	100000000	        24.9 ns/op
# BenchmarkBitCount_FFFFFFFF-4                	2000000000	         0.41 ns/op
# 
# BenchmarkPopCount_AllOnes-4                 	2000000000	         0.44 ns/op
# BenchmarkPopCountbyShifting_AllOnes-4       	20000000	       100 ns/op
# BenchmarkPopCountByClearingBit_AllOnes-4    	30000000	        58.5 ns/op
# BenchmarkBitCount_AllOnes-4                 	2000000000	         0.45 ns/op
