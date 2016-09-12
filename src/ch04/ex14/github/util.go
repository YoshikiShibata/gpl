package github

import (
	"encoding/json"
	"net/http"
	"strings"
)

func parseLink(link string) (next, last string) {
	if link == "" {
		return
	}

	links := strings.Split(link, ",")
	for _, link := range links {
		var p *string = nil
		if strings.Contains(link, `rel="next"`) {
			p = &next
		} else if strings.Contains(link, `rel="last"`) {
			p = &last
		} else {
			continue
		}
		sIndex := strings.Index(link, "<")
		eIndex := strings.Index(link, ">")
		*p = link[sIndex+1 : eIndex]
	}
	return
}

func parseBadRequest(resp *http.Response) error {
	var br BadRequest
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return err
	}
	return &br
}
