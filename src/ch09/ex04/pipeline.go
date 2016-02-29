package main

import (
	"fmt"
	"time"
)

const CHAN_CAPACITY = 0

func main() {
	next := make(chan int)
	go pipe(next, 0)
	time.Sleep(time.Hour)
}

func pipe(prev <-chan int, stages int) {
	next := make(chan int, CHAN_CAPACITY)

	stages++
	if stages%10000 == 0 {
		time.Sleep(time.Second)
		fmt.Printf("%d\n", stages)
	}
	go pipe(next, stages)

	for v := range prev {
		next <- v
	}
}

// Linux 16GB Memory
// 14,886,828 bytes free
// 14.15 GB 5,600,000 goroutines
