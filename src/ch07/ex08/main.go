// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"sort"
)

var title = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Title < tq.Title
}

var year = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Year < tq.Year
}

func main() {
	fmt.Println("\n=== original data ==")
	printTracks(tracksData)

	sortNormal()
	sortStable()
}

func sortNormal() {
	fmt.Println("\n=== sort.Sort ==")

	d := make([]*Track, len(tracksData))
	copy(d, tracksData)
	table := TrackMultiKeysSorter{tracks: d}

	table.AddSortKey(title)
	table.AddSortKey(year)
	sort.Sort(&table)
	printTracks(d)
}

func sortStable() {
	fmt.Println("\n=== sort.Stable ==")

	d := make([]*Track, len(tracksData))
	copy(d, tracksData)
	table := StableTrackMultiKeysSorter{tracks: d}

	table.AddSortKey(title)
	table.AddSortKey(year)

	for table.HasNext() {
		sort.Stable(&table)
	}
	printTracks(d)
}
