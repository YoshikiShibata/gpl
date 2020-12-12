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

// Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz, Linux(Ubuntu 15.0)
// 2016.06.28 Tip version of Go
// Elapsed Time = 1m40.000107401s
// 580596 per second
// 580596 per second

// 1.3GHz Intel Core M, MacOS X
// Elapsed Time = 1m40.000289819s
// 555193 per second
// 555193 per second

// 1.3GHz Intel Core M, MacOS X
// 2018.9.28 Go1.12 tip version
// Elapsed Time = 1m40.000101093s
// 537087 per second
// 537087 per second

// 1.3GHz Intel Core M, MacOS X
// 2018.9.28 Go1.12 tip version
// $ GOMAXPROCS=1 go run pingpong.go
// Elapsed Time = 1m40.056992551s
// 1050995 per second
// 1050995 per second

// Apple M1, macOS Big Sur (Version 11.0.1)
// 2020.12.12 Go 1.16 tip version
// % GOMAXPROCS=1 go run pingpong.go
// Elapsed Time = 1m40.013922708s
// 2451451 per second
// 2451451 per second
