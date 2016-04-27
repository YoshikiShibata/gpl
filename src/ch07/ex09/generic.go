// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

type LessFunc func(p1, p2 interface{}) bool

// MultiKeysSorter sorts data with multi-keys represented by a slice of LessFunc
type MultiKeysSorter struct {
	lessFuncs []LessFunc
}

func (m *MultiKeysSorter) AddSortKey(key LessFunc) {
	m.lessFuncs = append(m.lessFuncs, key)
}

func (m *MultiKeysSorter) LessWithMultiKeys(p, q interface{}) bool {
	if len(m.lessFuncs) == 0 {
		panic("Not Key is added as LessFunc")
	}

	var k int
	for k = 0; k < len(m.lessFuncs)-1; k++ {
		less := m.lessFuncs[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return m.lessFuncs[k](p, q)
}
