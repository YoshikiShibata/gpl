// Copyright © 2015 Alan A. A. Donovan & Brian W. Kernighan.

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

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	wclocks := parseArgs()
	if len(wclocks) == 0 {
		return
	}
	showWorldClock(wclocks[0])
}

func showWorldClock(wc *worldClock) {
	fmt.Println(wc.location)

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
		fmt.Println(string(bytes))
	}

	// mustCopy(os.Stdout, conn)
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

/*
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}
*/
