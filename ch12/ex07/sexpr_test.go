// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	//+ Exercise 12.3
	Color         bool
	BlackAndWhite bool
	Duration32    float32 // in minutes
	Duration64    float64 // in minutes
	/* Not supported by JSON
	Complex64     complex64 // just for testing.
	Complex128    complex128
	*/
	//- Exercise 12.3
	Actor  map[string]string
	Oscars []string
	Sequel *string
}

var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,

	//+ Exercise 12.3: This movie was not color, but for testing.
	Color:         false,
	BlackAndWhite: true,
	Duration32:    93.0,
	Duration64:    120.5,
	/*
		Complex64:     1.5 + 2.5i,
		Complex128:    2.5 + 3.5i,
	*/
	//- Exercise 12.3

	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
}

// TestSExpression verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
//	$ go test -v gopl.io/ch12/sexpr
func TestSExpression(t *testing.T) {
	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = \n%s\n", data)
}

// + Exercise 12.7
func TestStreamEncoderDecoder(t *testing.T) {
	// Encode it
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)
	err := encoder.Encode(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	data := buf.Bytes()
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	decoder := NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

//- Exercise 12.7

// + Exercise 12.6
func TestZeroValue(t *testing.T) {
	// Encode it
	var zeroMovie Movie
	data, err := Marshal(zeroMovie)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, zeroMovie) {
		t.Fatal("not equal")
	}
}

//- Exercise 12.6
