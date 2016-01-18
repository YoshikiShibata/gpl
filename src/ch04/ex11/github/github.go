// Copyright © 2016 Yoshiki Shibata

package github

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const GitHubAPIURL = "https://api.github.com"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

func (i *Issue) String() string {
	return fmt.Sprintf("#%-5d %9.9s %.55s",
		i.Number, i.User.Login, i.Title)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type CreateIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type EditIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"` // "open" or "close"
}

func issuesURL(repo string) string {
	return GitHubAPIURL + "/repos/" + repo + "/issues"
}

func newRequest(cmd, url string, body io.Reader, user *Credentials) (*http.Request, error) {
	req, err := http.NewRequest(cmd, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth(user.username, user.password)
	return req, nil
}
