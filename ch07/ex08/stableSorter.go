// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "fmt"

type StableTrackMultiKeysSorter struct {
	tracks    []*Track
	lessFuncs []LessFunc
	lessIndex int
}

func (t *StableTrackMultiKeysSorter) AddSortKey(key LessFunc) {
	t.lessFuncs = append(t.lessFuncs, key)
	t.lessIndex++
}

func (t *StableTrackMultiKeysSorter) Len() int {
	return len(t.tracks)
}

func (t *StableTrackMultiKeysSorter) Swap(i, j int) {
	t.tracks[i], t.tracks[j] = t.tracks[j], t.tracks[i]
}

func (t *StableTrackMultiKeysSorter) Less(i, j int) bool {
	if t.lessIndex < 0 {
		panic(fmt.Errorf("Out of Index: %d", t.lessIndex))
	}
	return t.lessFuncs[t.lessIndex](t.tracks[i], t.tracks[j])
}

func (t *StableTrackMultiKeysSorter) HasNext() bool {
	t.lessIndex--
	return t.lessIndex >= 0
}
