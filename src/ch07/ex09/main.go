// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"net/http"
	"sort"
	"strings"
)

var titleKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Title < tq.Title
}

var artistKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Artist < tq.Artist
}

var albumKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Album < tq.Album
}

var yearKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Year < tq.Year
}

var lengthKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Length < tq.Length
}

var sortKeyFuncs = map[string]func(p, q interface{}) bool{
	"TITLE":  titleKey,
	"ARTIST": artistKey,
	"ALBUM":  albumKey,
	"YEAR":   yearKey,
	"LENGTH": lengthKey,
}

func main() {
	sortNormal()
}

func sortNormal() {
	d := make([]*Track, len(tracksData))
	copy(d, tracksData)

	handler := func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("sort")
		f, keys := sortOrders(strings.Split(value, ","))
		if len(f) == 0 {
			printTracksHTML(w, d, keys)
			return
		}

		table := TrackMultiKeysSorter{tracks: d}
		for _, sk := range f {
			table.AddSortKey(sk)
		}
		sort.Sort(&table)
		printTracksHTML(w, d, keys)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func sortOrders(keys []string) ([]func(p, q interface{}) bool, []string) {
	updatedKeys := make([]string, 0, len(keys))
	sortOrderFuncs := make([]func(p, q interface{}) bool, 0, len(keys))

	for _, key := range keys {
		key := strings.TrimSpace(key)
		f, ok := sortKeyFuncs[strings.ToUpper(key)]
		if ok {
			updatedKeys = append(updatedKeys, key)
			sortOrderFuncs = append(sortOrderFuncs, f)
		}
	}
	return sortOrderFuncs, updatedKeys
}
