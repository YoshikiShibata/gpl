// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
)

type client struct {
	who string
	out chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

const clientTimeOut = time.Minute * 5

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.out <- msg
			}

		case cli := <-entering:
			for client := range clients {
				cli.out <- client.who + " is in now"
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)

	var firstMsg string
	who := conn.RemoteAddr().String()
	if input.Scan() {
		firstMsg = input.Text()
		if user, ok := extractUser(firstMsg); ok {
			who = user
			firstMsg = ""
		}
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}

	if len(firstMsg) > 0 {
		messages <- who + ": " + input.Text()
	}

	timer := time.AfterFunc(clientTimeOut, func() {
		conn.Close()
	})
	for input.Scan() {
		timer.Stop()
		messages <- who + ": " + input.Text()
		timer = time.AfterFunc(clientTimeOut, func() {
			conn.Close()
		})
	}
	timer.Stop()

	// NOTE: ignoring potential errors from input.Err()
	leaving <- client{who, ch}
	messages <- who + " has left"
	conn.Close()
}

var userPattern = regexp.MustCompile("<user>(.+)</user>")

func extractUser(msg string) (string, bool) {
	matches := userPattern.FindAllStringSubmatch(msg, -1)
	if matches == nil {
		return "", false
	}
	return matches[0][1], true
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
