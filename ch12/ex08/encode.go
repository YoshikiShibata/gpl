// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

//+ Exercise 12.7
type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

func (enc *Encoder) Encode(v interface{}) error {
	if err := encode(enc.writer, reflect.ValueOf(v)); err != nil {
		return err
	}
	return nil
}

var byteBuff = make([]byte, 1)

func writeByte(w io.Writer, c byte) {
	byteBuff[0] = c
	w.Write(byteBuff)
}

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		io.WriteString(buf, "nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		writeByte(buf, '(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				writeByte(buf, ' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		writeByte(buf, ')')

	case reflect.Struct: // ((name value) ...)
		writeByte(buf, '(')
		first := true //+ Exercise 12.6
		for i := 0; i < v.NumField(); i++ {
			//+ Exercise 12.6
			if isZeroValue(v.Field(i)) {
				continue
			}
			if !first {
				writeByte(buf, ' ')
			}
			//- Exercise 12.6

			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			writeByte(buf, ')')
			first = false //+ Exercise 12.6
		}
		writeByte(buf, ')')

	case reflect.Map: // ((key value) ...)
		writeByte(buf, '(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				writeByte(buf, ' ')
			}
			writeByte(buf, '(')
			if err := encode(buf, key); err != nil {
				return err
			}
			writeByte(buf, ' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			writeByte(buf, ')')
		}
		writeByte(buf, ')')

	//+ Exercise 12.3
	case reflect.Bool: // t or nil
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, "nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		v := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(v), imag(v))

	case reflect.Interface:
		writeByte(buf, '(')
		t := v.Type()
		if t.Name() == "" { // empty interface
			fmt.Fprintf(buf, "%q ", v.Elem().Type().String())
		} else {
			fmt.Fprintf(buf, `"%s.%s" `, t.PkgPath(), t.Name())
		}

		if err := encode(buf, v.Elem()); err != nil {
			return err
		}
		writeByte(buf, ')')

	//- Exercise 12.3

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

//!-encode
