// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// JsonMarshal encodes a Go value in JSON form.
func JsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	buf.Grow(2 * 1024)
	if err := jsonEncode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var scratch [64]byte

// jsonEncode writes to buf a JSON representation of v.
func jsonEncode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		b := strconv.AppendInt(scratch[:0], v.Int(), 10)
		buf.Write(b)
		// fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		b := strconv.AppendUint(scratch[:0], v.Uint(), 10)
		buf.Write(b)
		// fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return jsonEncode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := jsonEncode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := jsonEncode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := jsonEncode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Bool: // true or false
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	default: // complex, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}

	return nil
}
