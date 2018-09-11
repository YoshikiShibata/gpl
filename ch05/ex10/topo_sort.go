// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

type itemSet map[string]bool

func (is itemSet) add(items ...string) {
	for _, item := range items {
		is[item] = true
	}
}

func newItemSet(items ...string) itemSet {
	is := make(itemSet)
	is.add(items...)
	return is
}

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]itemSet{
	"algorithms": newItemSet("data structures"),
	"calculus":   newItemSet("linear algebra"),

	"compilers": newItemSet(
		"data structures",
		"formal languages",
		"computer organization"),

	"data structures":  newItemSet("discrete math"),
	"databases":        newItemSet("data structures"),
	"discrete math":    newItemSet("intro to programming"),
	"formal languages": newItemSet("discrete math"),
	"networks":         newItemSet("operating systems"),
	"operating systems": newItemSet(
		"data structures",
		"computer organization"),
	"programming languages": newItemSet(
		"data structures",
		"computer organization"),
}

func main() {
	ts := topoSort(prereqs)
	for i, course := range ts {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Printf("Topological Orderings: %v\n", isTopologicalOrdered(ts))
}

func topoSort(m map[string]itemSet) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items itemSet)

	visitAll = func(items itemSet) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	items := newItemSet()
	for item := range m {
		items.add(item)
	}

	visitAll(items)
	return order
}

func isTopologicalOrdered(ts []string) bool {
	nodes := make(map[string]int)

	for i, course := range ts {
		nodes[course] = i
	}

	for course, i := range nodes {
		for prereq := range prereqs[course] {
			if i < nodes[prereq] {
				return false
			}
		}
	}
	return true
}
