// Copyright © 2016 Yoshiki Shibata

package github

import (
	"fmt"
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
