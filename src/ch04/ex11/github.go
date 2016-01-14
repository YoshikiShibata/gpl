// Copyright © 2016 Yoshiki Shibata

package main

import "time"

const GitHubAPIURL = "https://api.github.com/repos"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type CreateIssue struct {
	Title string
	Body  string
}

type EditIssue struct {
	Title string
	Body  string
	State string // "open" or "close"
}
