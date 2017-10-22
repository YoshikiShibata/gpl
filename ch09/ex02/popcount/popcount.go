// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import "sync"

// pc[i] is the population count of i.
var pc [256]byte

func initPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var initPCOnce sync.Once

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	initPCOnce.Do(initPC)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
