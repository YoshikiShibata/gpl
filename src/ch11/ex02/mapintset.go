package intset

import (
	"bytes"
	"fmt"
	"sort"
)

// An MapIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type MapIntSet struct {
	set map[int]bool
}

// Has reports whether the set contains the non-negaive value x.
func (s *MapIntSet) Has(x int) bool {
	if s.set == nil {
		return false
	}
	return s.set[x]
}

// Add adds the non-negative value x to the set.
func (s *MapIntSet) Add(x int) {
	if s.set == nil {
		s.set = make(map[int]bool)
	}
	s.set[x] = true
}

// UnionWith sets s to the union of s and t
func (s *MapIntSet) UnionWith(t *MapIntSet) {
	if t.set == nil {
		return
	}
	if s.set == nil {
		s.set = make(map[int]bool)
	}

	for x, b := range t.set {
		if b {
			s.set[x] = true
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapIntSet) String() string {
	if s.set == nil {
		return "{ }"
	}

	ints := make([]int, 0, len(s.set))
	for x, v := range s.set {
		if v {
			ints = append(ints, x)
		}
	}

	sort.Ints(ints)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, x := range ints {
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", x)
	}
	buf.WriteByte('}')
	return buf.String()
}
