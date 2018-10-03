// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.

// Netcat is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type worldClock struct {
	location string
	url      string
}

type time struct {
	index int
	time  string
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	out := make(chan time, 3)
	wclocks := parseArgs()
	if len(wclocks) == 0 {
		return
	}
	for i, wc := range wclocks {
		go showWorldClock(wc, out, i)
	}

	for {
		showWorldTimes(out, len(wclocks))
	}
}

func showWorldTimes(out <-chan time, noOfClocks int) {
	times := make([]string, noOfClocks)
	for i := 0; i < noOfClocks; i++ {
		t := <-out
		times[t.index] = t.time
	}
	fmt.Printf("\r%s", strings.Join(times, " "))
}

func showWorldClock(wc *worldClock, out chan<- time, index int) {

	conn, err := net.Dial("tcp", wc.url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		out <- time{index, fmt.Sprintf("%s: %s ", wc.location, string(bytes))}
	}
}

func parseArgs() []*worldClock {
	clocks := []*worldClock{}

	if len(os.Args) == 1 {
		return clocks
	}

	for _, clockSpec := range os.Args[1:] {
		wc, err := parseClockSpec(clockSpec)
		if err != nil {
			fmt.Printf("%s: %v\n", clockSpec, err)
		} else {
			clocks = append(clocks, wc)
		}
	}
	return clocks
}

func parseClockSpec(spec string) (*worldClock, error) {
	components := strings.Split(spec, "=")
	if len(components) != 2 {
		return nil, fmt.Errorf("illegal")
	}
	return &worldClock{components[0], components[1]}, nil
}
