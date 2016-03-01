// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import (
	"errors"
	"strings"
)

// Func is the type of the function to memoize.
// Closing done channel means the cancellation of this func.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// ErrCancelled indicates that Get operation is cancelled by the client
// and no cache is updated.
var ErrCanceled = errors.New("request canceled")

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	done     <-chan struct{} // the cleint wants to cancel when closed
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	if res.err == ErrCanceled {
		select {
		case <-done:
			// client cancelled
			return res.value, res.err
		default:
			// client did not cancelled. Retry again.
			return memo.Get(key, done)
		}
	}
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]

		if e != nil {
			select {
			case <-e.ready:
				if e.res.err == ErrCanceled {
					// Previous one was cancelled
					delete(cache, req.key)
					e = nil
				}
			default:
				// do nothing
			}
		}

		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, req.done) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, done)
	if e.res.err != nil &&
		strings.Contains(e.res.err.Error(), "request canceled") {
		e.res.err = ErrCanceled
	}
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
