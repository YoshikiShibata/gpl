// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package main

import (
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrint(t *testing.T) {
	for _, url := range []string{
		"http://gopl.io",
		"https://golang.org",
		"https://jp.merpay.com",
	} {
		t.Run(url, func(t *testing.T) {
			// prepare
			resp, err := http.Get(url)
			if err != nil {
				t.Logf("http.Get failed: %v", err)
				// skip this case
				return
			}

			doc, err := html.Parse(resp.Body)
			if err != nil {
				t.Errorf("[%s] Parse failed: %v", url, err)
				resp.Body.Close()
				return
			}
			resp.Body.Close()

			// action
			result := prettyPrint(doc)

			// check
			doc, err = html.Parse(strings.NewReader(result))
			if err != nil {
				t.Errorf("[%s] Parse failed: %v", url, err)
				t.Errorf("The result of prettyPrint: %s\n", result)
			}
		})
	}
}
