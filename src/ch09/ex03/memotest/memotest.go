// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if done != nil {
		cancel := make(chan struct{})
		req.Cancel = cancel
		go func() {
			<-done
			close(cancel)
		}()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
		} else {
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}
	}
}

func SequentialCancel(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		go func() {
			time.Sleep(time.Second / 100) // sleep 50 ms
			close(done)
		}()

		value, err := m.Get(url, done)
		if err != nil {
			fmt.Printf("%s: %v\n", url, err)
		} else {
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
			} else {
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			}
			n.Done()
		}(url)
	}
	n.Wait()
}

func ConcurrentCancel(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			done := make(chan struct{})
			go func() {
				time.Sleep(time.Second / 10) // sleep 100 ms
				close(done)
			}()

			value, err := m.Get(url, done)
			if err != nil {
				fmt.Printf("%s: %v\n", url, err)
			} else {
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			}
			n.Done()
		}(url)
	}
	n.Wait()
}
