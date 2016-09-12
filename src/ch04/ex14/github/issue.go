package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

// URL for listing issues for a repository.
const ListIssuesURL = "https://api.github.com/repos/%s/%s/issues"

type IssuesListResult struct {
	Issues   []*Issue
	nextLink string
	lastLink string
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // (Markdown format)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func ListIssues(owner, repo string) (*IssuesListResult, error) {
	listURL := fmt.Sprintf(ListIssuesURL, owner, repo)
	fmt.Printf("listURL = %s\n", listURL)
	return listIssues(listURL)
}

func listIssues(listURL string) (*IssuesListResult, error) {
	resp, err := http.Get(listURL)
	if err != nil {
		return nil, err
	}

	result := parseLink(resp.Header.Get("Link"))

	if err := json.NewDecoder(resp.Body).Decode(&(result.Issues)); err != nil {
		resp.Body.Close()
		fmt.Printf("listURL = %q\n", listURL)
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func parseLink(link string) *IssuesListResult {
	var result IssuesListResult

	if link == "" {
		return &result
	}

	links := strings.Split(link, ",")
	for _, link := range links {
		var p *string = nil
		if strings.Contains(link, `rel="next"`) {
			p = &(result.nextLink)
		} else if strings.Contains(link, `rel="last"`) {
			p = &(result.lastLink)
		} else {
			continue
		}
		sIndex := strings.Index(link, "<")
		eIndex := strings.Index(link, ">")
		*p = link[sIndex+1 : eIndex]
	}
	return &result
}

func (il *IssuesListResult) HasNext() bool {
	return il.nextLink != ""
}

func (il *IssuesListResult) Next() (*IssuesListResult, error) {
	if il.nextLink == "" {
		panic("NextLink is not available")
	}

	return listIssues(il.nextLink)
}
