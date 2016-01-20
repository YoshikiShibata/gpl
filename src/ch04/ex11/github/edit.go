// Copyright © 2016 Yoshiki Shibata

package github

func Edit(repo, title, body string, issueNo int, user *Credentials) (*Issue, error) {
	issue := EditIssue{title, body}

	return patch(repo, &issue, issueNo, user)
}
