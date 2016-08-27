// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	// data, err := json.MarshalIndent(strangelove, "", "    ")
	data, err := json.Marshal(strangelove)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	t.Logf("json.Marshal() = \n%s\n\n", data)

	data, err = JsonMarshal(strangelove)
	if err != nil {
		t.Fatalf("JsonMarshal failed: %v", err)
	}
	t.Logf("jsonMarshal() = \n%s\n", data)

	// Decode it
	var movie Movie
	if err := json.Unmarshal(data, &movie); err != nil {
		t.Fatalf("JSON Unmarshal failed: %v", err)
	}
	t.Logf("json.Unmarshal() = \n%+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func Benchmark_json_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(strangelove)
	}
}

func Benchmark_JsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonMarshal(strangelove)
	}
}
