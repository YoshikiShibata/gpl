// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	constraints := make(map[string]string) // Exercise 12.12
	v := reflect.ValueOf(ptr).Elem()       // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		//+ Exercise 12.12
		tags := strings.Split(name, ",")
		name = tags[0]
		fields[name] = v.Field(i)
		if len(tags) == 2 {
			constraints[name] = tags[1]
		}
		//- Exercise 12.12
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			//+ Exercise 12.12
			if constraints[name] != "" {
				if !validateValue(value, constraints[name]) {
					return fmt.Errorf("%s invalid for %s", value, name)
				}
			}
			//- Exercise 12.12
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

// !+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate

// + Exercise 12.12
var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9\-.]+@[a-zA-Z0-9\-.]+$`)
var creditPattern = regexp.MustCompile(`^[0-9]{10}$`)
var zipPattern = regexp.MustCompile(`^[0-9]{7}$`)

func validateValue(value, constraint string) bool {
	switch constraint {
	case "email":
		return emailPattern.MatchString(value)
	case "credit":
		return creditPattern.MatchString(value)
	case "zip":
		return zipPattern.MatchString(value)
	default:
		return false
	}
}

//- Exercise 12.12
