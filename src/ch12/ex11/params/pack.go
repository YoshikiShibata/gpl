// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package params

import (
	"fmt"
	"reflect"
	"strings"
)

//+ Exercise 12.11

func Pack(ptr interface{}) string {
	var result []string

	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		result = append(result, toParam(name, v.Field(i)))
	}
	return strings.Join(result, "&")
}

func toParam(name string, v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s=%s", name, v.String())
	case reflect.Int:
		return fmt.Sprintf("%s=%d", name, v.Int())
	case reflect.Bool:
		if v.Bool() {
			return fmt.Sprintf("%s=true", name)
		} else {
			return fmt.Sprintf("%s=false", name)
		}
	case reflect.Array, reflect.Slice:
		var result []string
		for i := 0; i < v.Len(); i++ {
			result = append(result, toParam(name, v.Index(i)))
		}
		return strings.Join(result, "&")

	default:
		panic(fmt.Sprintf("Unspported Kind = %d", v.Kind()))
	}
}

//- Exercise 12.11
