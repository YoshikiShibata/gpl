// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"reflect"
	"testing"
)

func TestInterfaces(t *testing.T) {
	type EmptyInterfaces struct {
		Int        interface{}
		Int8       interface{}
		Int16      interface{}
		Int32      interface{}
		Int64      interface{}
		Uint       interface{}
		Uint8      interface{}
		Uint16     interface{}
		Uint32     interface{}
		Uint64     interface{}
		String     interface{}
		Complex64  interface{}
		Complex128 interface{}
		// Composite types
		IntSlice    interface{}
		StringArray interface{}
		Map         interface{}
	}

	interfacesValue := EmptyInterfaces{
		Int:         int(0),
		Int8:        int8(1),
		Int16:       int16(2),
		Int32:       int32(3),
		Int64:       int64(4),
		Uint:        uint(5),
		Uint8:       uint8(6),
		Uint16:      uint16(7),
		Uint32:      uint32(8),
		Uint64:      uint64(9),
		String:      "こんにちは、世界！",
		Complex64:   complex64(1.5 + 2.5i),
		Complex128:  complex128(3.5 + 4.5i),
		IntSlice:    []int{1, 2, 3, 4, 5},
		StringArray: [...]string{"Hello", "World"},
		Map:         map[string]int{"A": 0, "B": 1},
	}

	// Encode it
	data, err := Marshal(interfacesValue)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var values EmptyInterfaces
	if err := Unmarshal(data, &values); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", values)

	// Check equality.
	if !reflect.DeepEqual(values, interfacesValue) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(interfacesValue)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIndent() = %s\n", data)
}
