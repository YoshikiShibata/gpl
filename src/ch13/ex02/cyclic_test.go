// Copyright © 2016, 2017 Yoshiki Shibata. All rights reserved.

package main

import "testing"

//+ Exercise 13.2

func TestCyclicLinkList(t *testing.T) {
	type link struct {
		value string
		tail  *link
		prev  *link
	}

	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	d, e, f := &link{value: "d"}, &link{value: "e"}, &link{value: "f"}
	a.tail, b.tail = b, a // cyclic
	c.tail = c            // self cyclic

	d.tail, d.prev, e.tail = e, e, nil // non-cyclic

	for _, test := range []struct {
		l      *link
		cyclic bool
	}{
		{a, true},
		{c, true},
		{d, false},
		{f, false},
	} {
		if IsCyclic(test.l) != test.cyclic {
			t.Errorf("IsCyclic(%s) is %v, but want %v",
				test.l.value, IsCyclic(test.l), test.cyclic)
		}
	}
}

func TestCyclicSlice(t *testing.T) {
	type pointer struct {
		next []*pointer
	}

	a := &pointer{}
	b := &pointer{}
	c := &pointer{}

	a.next = append(a.next, b)
	b.next = append(b.next, c)
	c.next = append(c.next, nil)
	if IsCyclic(a) {
		t.Errorf("IsCyclic(a) is true, but want false")
	}

	c.next = append(c.next, a)
	if !IsCyclic(a) {
		t.Errorf("IsCyclic(a) is false, but want true")
	}
}

func TestCyclicArray(t *testing.T) {
	type pointer struct {
		next [1]*pointer
	}

	a := &pointer{}
	b := &pointer{}
	c := &pointer{}

	a.next[0] = b
	b.next[0] = c
	if IsCyclic(a) {
		t.Errorf("IsCyclic(a) is true, but want false")
	}

	c.next[0] = a
	if !IsCyclic(a) {
		t.Errorf("IsCyclic(a) is false, but want true")
	}
}

func TestCyclicMap(t *testing.T) {
	type pointer struct {
		value map[string]*pointer
	}
	a := &pointer{make(map[string]*pointer)}
	b := &pointer{make(map[string]*pointer)}
	c := &pointer{make(map[string]*pointer)}

	a.value["a"] = b
	b.value["b"] = c

	if IsCyclic(a) {
		t.Errorf("IsCyclic(a) is true, but want false")
	}

	c.value["c"] = a
	if !IsCyclic(a) {
		t.Errorf("IsCyclic(a) is false, but want true")
	}

}

func TestCyclicRecursiveSlice(t *testing.T) {
	type S []S
	var s = make(S, 1)
	s[0] = s

	if !IsCyclic(s) {
		t.Errorf("IsCyclic(s) is false, but want true")
	}
}

//- Exercise 13.2
