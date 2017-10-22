// Copyright © 2016 Yoshiki Shibata. All rights reserved

package intset

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// seed to be used. Note that this variable is initialized only once per run
var seed int64 = time.Now().UTC().UnixNano()

func BenchmarkAdd_IntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add_IntSet()
	}
}

func BenchmarkAdd_MapIntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add_MapIntSet()
	}
}

func BenchmarkUnitonWith_IntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionWith_IntSet()
	}
}

func BenchmarkUnionWith_MapIntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unionWith_MapIntSet()
	}
}

func add_IntSet() {
	var set IntSet

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 500; i++ {
		set.Add(rng.Intn(math.MaxInt16))
	}
}

func add_MapIntSet() {
	var set MapIntSet

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 500; i++ {
		set.Add(rng.Intn(math.MaxInt16))
	}
}

func unionWith_IntSet() {
	var x IntSet
	var y IntSet

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 500; i++ {
		x.Add(rng.Intn(math.MaxInt16))
		y.Add(rng.Intn(math.MaxInt16))
	}
	x.UnionWith(&y)
}

func unionWith_MapIntSet() {
	var x MapIntSet
	var y MapIntSet

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 500; i++ {
		x.Add(rng.Intn(math.MaxInt16))
		y.Add(rng.Intn(math.MaxInt16))
	}
	x.UnionWith(&y)
}
