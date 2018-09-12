// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2016, 2018 Yoshiki Shibata. All rights reserved.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"linear algebra": {"calculus"},
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	ts, err := topoSort(prereqs)

	for i, course := range ts {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}

				for _, orderedItem := range order {
					for _, prereq := range m[orderedItem] {
						if prereq == item {
							return fmt.Errorf("%q and %q are cycled",
								orderedItem, item)
						}
					}
				}

				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	return order, err
}
