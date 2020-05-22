// Copyright © 2016, 2020 Yoshiki Shibata

package main

import "testing"

func TestNil(t *testing.T) {
	if eliminateAdjacentDuplicates(nil) != nil {
		t.Errorf("Passing nil returns %v, want nil",
			eliminateAdjacentDuplicates(nil))
	}
}

func TestEliminateAdjacentDuplicates(t *testing.T) {
	data := []struct {
		s        []string
		expected []string
	}{
		{
			[]string{},
			[]string{}},
		{
			[]string{"Hello"},
			[]string{"Hello"}},
		{
			[]string{"Hello", "World", "World"},
			[]string{"Hello", "World"}},
		{
			[]string{"Hello", "Hello", "World", "World"},
			[]string{"Hello", "World"}},
		{
			[]string{"Hello", "Hello", "World", "World", "World"},
			[]string{"Hello", "World"}},
		{
			[]string{"Hello", "Hello", "Hello", "World", "World", "World"},
			[]string{"Hello", "World"}},
		{
			[]string{"Hello", "Hello", "World", "World", "Hello", "World"},
			[]string{"Hello", "World", "Hello", "World"}},
	}

	duplicateSlice := func(s []string) []string {
		duplicated := make([]string, len(s))
		copy(duplicated, s)
		return duplicated
	}

	t.Run("1", func(t *testing.T) {
		for _, d := range data {
			result := eliminateAdjacentDuplicates(duplicateSlice(d.s))
			if len(result) != len(d.expected) {
				t.Errorf("Result length is %d, want %d",
					len(result), len(d.expected))
			}
			for i := 0; i < len(d.expected); i++ {
				if result[i] != d.expected[i] &&
					d.s[i] != d.expected[i] { // in-place test
					t.Errorf(`result[%d] is "%s", want "%s"`,
						i, result[i], d.expected[i])
				}
			}
		}
	})

	t.Run("2", func(t *testing.T) {
		for _, d := range data {
			result := eliminateAdjacentDuplicates2(duplicateSlice(d.s))
			if len(result) != len(d.expected) {
				t.Errorf("Result length is %d, want %d",
					len(result), len(d.expected))
			}
			for i := 0; i < len(d.expected); i++ {
				if result[i] != d.expected[i] &&
					d.s[i] != d.expected[i] { // in-place test
					t.Errorf(`result[%d] is "%s", want "%s"`,
						i, result[i], d.expected[i])
				}
			}
		}
	})
}
