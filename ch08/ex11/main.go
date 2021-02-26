// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016, 2021 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var done = make(chan struct{})
var wg sync.WaitGroup

func cancelled(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch) // start a goroutine
	}
	fmt.Println(<-ch) // receive from channel ch
	close(done)       // cancel others
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait()
}

func fetch(url string, ch chan<- string) {
	defer wg.Done()

	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if cancelled(done) {
		return
	}

	cancelChan := make(chan struct{})
	req.Cancel = cancelChan

	go func() {
		select {
		case <-done:
			close(cancelChan)
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case ch <- fmt.Sprint(err): // send to channel ch
		case <-done:
		}
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	select {
	case ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url):
	case <-done:
	}
}
