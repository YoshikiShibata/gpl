// Copyright © 2016 Yoshiki Shibata

package github

func Close(repo string, issueNo int, user *Credentials) (*Issue, error) {
	issue := CloseIssue{"close"}

	return patch(repo, &issue, issueNo, user)
}
