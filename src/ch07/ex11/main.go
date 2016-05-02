// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{products: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	products map[string]dollars
	sync.RWMutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()

	for item, price := range db.products {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()

	item := req.URL.Query().Get("item")
	if price, ok := db.products[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item, ok := getQueryString(w, req, "item")
	if !ok {
		return
	}
	price, ok := getQueryPositiveInt(w, req, "price")
	if !ok {
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.products[item]; ok {
		db.products[item] = dollars(price)
		fmt.Fprintf(w, "%s: %s\n", item, db.products[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item, ok := getQueryString(w, req, "item")
	if !ok {
		return
	}
	price, ok := getQueryPositiveInt(w, req, "price")
	if !ok {
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.products[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item(%s) has already exist. \n", item)
	} else {
		db.products[item] = dollars(price)
		fmt.Fprintf(w, "%s: %s\n", item, db.products[item])
	}
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	item, ok := getQueryString(w, req, "item")
	if !ok {
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.products[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, db.products[item])
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item(%s) is not found.\n", item)
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item, ok := getQueryString(w, req, "item")
	if !ok {
		return
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.products[item]; ok {
		delete(db.products, item)
		fmt.Fprintf(w, "item(%s) was deleted. \n", item)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s doesn't exist\n", item)
	}
}

// getQuery returns only the first parameter
func getQueryString(w http.ResponseWriter, req *http.Request,
	query string) (string, bool) {
	params, ok := req.URL.Query()[query] // Query returns map[string][]string
	if !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s is not specified\n", query)
		return "", false
	}
	if len(params) == 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "No value is specified for %s\n", query)
		return "", false
	}

	if len(params[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "No value is specified for %s\n", query)
		return "", false
	}

	return params[0], true
}

func getQueryPositiveInt(w http.ResponseWriter, req *http.Request, query string) (int, bool) {
	value, ok := getQueryString(w, req, query)
	if !ok {
		return 0, false
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "value for %s is illegal format(%s)\n", query, value)
		return 0, false
	}

	if intValue < 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "negative %s(%d) is not accepted\n", query, intValue)
		return 0, false
	}

	return int(intValue), true
}
