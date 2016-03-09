// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// Expand replaces each substring "$foo" within s by the text returned by f("foo")
package main

import "regexp"

var pattern = regexp.MustCompile(`(\$\w*)`)

func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s,
		func(sub string) string {
			return f(sub[1:])
		})
}
