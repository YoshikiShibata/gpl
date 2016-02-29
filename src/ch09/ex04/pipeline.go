package main

import (
	"fmt"
	"time"
)

const (
	CHAN_CAPACITY  = 0
	MAX_GOROUTINES = 4600000
	MEASURE_COUNT  = 100
)

func main() {
	next := make(chan int)
	final := make(chan int)
	go pipe(next, 0, final)

	next <- 0    // the initial value
	i := <-final // receiving the final value to wait for the creattion of all goroutines
	if i != 0 {
		panic("i != 0")
	}
	fmt.Printf("%d goroutines are created\n", MAX_GOROUTINES)

	var total int64

	for v := 1; v <= MEASURE_COUNT; v++ {
		start := time.Now()
		next <- v
		r := <-final
		end := time.Now()
		if v != r {
			panic(fmt.Errorf("%d recevied, but want %d", r, v))
		}
		diff := end.Sub(start)
		fmt.Printf("%3d: %v\n", v, diff)
		total += diff.Nanoseconds()
	}

	fmt.Printf("average round trip time = %d nano seconds\n",
		total/MEASURE_COUNT)
}

func pipe(prev <-chan int, stages int, final chan<- int) {
	next := make(chan int, CHAN_CAPACITY)

	stages++
	if stages%10000 == 0 {
		time.Sleep(time.Second)
		fmt.Printf("%d\n", stages)
	}

	if stages == MAX_GOROUTINES {
		for v := range prev {
			final <- v
		}
	} else {
		go pipe(next, stages, final)

		for v := range prev {
			next <- v
		}
	}
}

// Linux 16GB Memory
// 14,886,828 bytes free
// 14.15 GB 5,600,000 goroutines
