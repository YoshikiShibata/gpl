package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const listIssuesURL = "https://api.github.com/repos/%s/%s/issues"

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

func ListIssues(owner, repo string) (*IssuesListResult, error) {
	listURL := fmt.Sprintf(listIssuesURL, owner, repo)
	fmt.Printf("listURL = %s\n", listURL)
	return listIssues(listURL)
}

func listIssues(listURL string) (*IssuesListResult, error) {
	resp, err := http.Get(listURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, parseBadRequest(resp)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status Code is %d", resp.StatusCode)
	}

	var result IssuesListResult

	result.nextLink, result.lastLink = parseLink(resp.Header.Get("Link"))

	if err := json.NewDecoder(resp.Body).Decode(&(result.Issues)); err != nil {
		fmt.Printf("listURL = %q\n", listURL)
		return nil, err
	}
	return &result, nil
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
