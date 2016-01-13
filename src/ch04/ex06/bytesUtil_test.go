// Copyright © 2016 Yoshiki Shibata

package main

import (
	"bytes"
	"testing"
)

func TestNil(t *testing.T) {
	if squashSpaces(nil) != nil {
		t.Errorf("squashSpace(nil) returns %v, but want nil",
			squashSpaces(nil))
	}
}

func TestNoSquash(t *testing.T) {
	s := "Hello World!"
	result := string(squashSpaces([]byte(s)))
	if s != result {
		t.Errorf("Result is %s, but want %s", result, s)
	}
}

func TestSquash(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			"",
			""},
		{
			"Hello  World!",
			"Hello World!"},
		{
			"Hello \n \t \n \v \f \r World!",
			"Hello World!"},
		{
			"Hello　World!", // Japanese white space
			"Hello World!"},
		{
			"Hello   World!     !",
			"Hello World! !"},
	}

	for _, d := range data {
		b := []byte(d.s)
		squashed := squashSpaces(b)
		result := string(squashed)
		if result != d.expected {
			t.Errorf(`Result is "%s", but want "%s"`, result, d.expected)
		}

		// Check if b is actually squashed
		if !bytes.Equal(squashed, b[:len(squashed)]) {
			t.Errorf("Not In place")
		}
	}
}
