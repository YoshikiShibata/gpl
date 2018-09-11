// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package main

import "testing"

func TestToposort(t *testing.T) {
	for i := 0; i < 100; i++ {
		ts := topoSort(prereqs)
		if !isTopologicalOrdered(ts) {
			t.Errorf("Not Topological Ordered: %v", ts)
		}
	}
}
