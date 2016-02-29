// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"time"
)

const durationInSeconds = 100 // seconds

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	end := make(chan struct{})
	result := make(chan int)

	go pingPong(ch1, ch2, end, result)
	go pingPong(ch2, ch1, end, result)

	start := time.Now()
	ch1 <- struct{}{}
	<-time.Tick(time.Second * durationInSeconds)
	close(end)
	r1 := <-result
	r2 := <-result
	elapsed := time.Now().Sub(start)

	fmt.Printf("Elapsed Time = %v\n", elapsed)
	fmt.Printf("%d per second\n", r1/durationInSeconds)
	fmt.Printf("%d per second\n", r2/durationInSeconds)
}

func pingPong(in <-chan struct{}, out chan<- struct{},
	end <-chan struct{}, result chan<- int) {
	for i := 0; ; i++ {
		select {
		case v := <-in:
			select {
			case out <- v:
			case <-end:
				result <- i
				return
			}
		case <-end:
			result <- i
			return
		}
	}
}

// Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz, Linux(Ubuntu 15.0)
// Elapsed Time = 1m40.000249347s
// 542444 per second
// 542444 per second

// 1.3GHz Interl Core M, MacOS X
// Elapsed Time = 1m40.000289819s
// 555193 per second
// 555193 per second
