// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// See page 333.

// Package display provides a means to display structured data.
package display

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

//!+Display

var nestLevel int // Exercise 12.2

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	nestLevel = 0 // Exercise 12.2
	display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)

	//+ Exercise 12.1
	case reflect.Struct:
		var b bytes.Buffer

		b.WriteString(v.Type().String())
		b.WriteRune('{')
		for i := 0; i < v.NumField(); i++ {
			b.WriteString(fmt.Sprintf("%s: %s",
				v.Type().Field(i).Name,
				formatAtom(v.Field(i))))
			if i < (v.NumField() - 1) {
				b.WriteString(", ")
			}
		}
		b.WriteRune('}')
		return b.String()

	case reflect.Array:
		var b bytes.Buffer

		b.WriteString(v.Type().String())
		b.WriteRune('{')
		for i := 0; i < v.Len(); i++ {
			b.WriteString(formatAtom(v.Index(i)))
			if i < (v.Len() - 1) {
				b.WriteString(", ")
			}
		}
		b.WriteRune('}')
		return b.String()

	//- Exercise 12.1
	default: // reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value) {
	//+ Exercise 12.2
	nestLevel++
	if nestLevel > 20 {
		return
	}
	//- Exercise 12.2

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-display
