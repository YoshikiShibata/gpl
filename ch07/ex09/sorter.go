// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

type TrackMultiKeysSorter struct {
	tracks []*Track
	MultiKeysSorter
}

func (t *TrackMultiKeysSorter) Len() int {
	return len(t.tracks)
}

func (t *TrackMultiKeysSorter) Swap(i, j int) {
	t.tracks[i], t.tracks[j] = t.tracks[j], t.tracks[i]
}

func (t *TrackMultiKeysSorter) Less(i, j int) bool {
	return t.LessWithMultiKeys(t.tracks[i], t.tracks[j])
}
