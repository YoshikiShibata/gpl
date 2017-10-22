// Copyright © 2015 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015 Yoshiki Shibata

// Clock is a TCP server that prints the current time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		time.Sleep(time.Second)
	}
}

var port = flag.Int("port", 8000, "port number")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn)
	}
}
