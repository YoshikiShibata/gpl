// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"reflect"
	"unsafe"
)

//+ Exercise 13.2

// IsCyclic reports whether x is a cyclic structure.
func IsCyclic(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

func isCyclic(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !x.IsValid() {
		return false
	}

	if x.CanAddr() &&
		x.Kind() != reflect.Struct &&
		x.Kind() != reflect.Array {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true
		}
		seen[xptr] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Slice, reflect.Array:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}

	return false
}

//- Exercise 13.2
