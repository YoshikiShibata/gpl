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
	} {
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("http.Get failed: %v", err)
			continue
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			t.Errorf("[%s] Parse failed: %v", url, err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		result := prettyPrint(doc)
		doc, err = html.Parse(strings.NewReader(result))
		if err != nil {
			t.Errorf("[%s] Parse failed: %v", url, err)
			t.Errorf("The result of prettyPrint: %s\n", result)
		}
	}
}
