// Copyright © 2016 Yoshiki Shibata

package main

import "crypto/sha256"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func totalPopCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(pc[b])
	}
	return total
}

func popCountDiff(sum1, sum2 [sha256.Size]byte) int {
	xorBytes := make([]byte, 0, sha256.Size)
	for i := 0; i < sha256.Size; i++ {
		xorBytes = append(xorBytes, sum1[i]^sum2[i])
	}
	return totalPopCount(xorBytes)
}
