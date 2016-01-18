// Copyright © 2016 Yoshiki Shibata

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(repo string, issueNo int, user *Credentials) (*Issue, error) {
	restAPIURL := issuesURL(repo) + fmt.Sprintf("/%d", issueNo)

	req, err := newRequest("GET", restAPIURL, nil, user)
	if err != nil {
		return nil, fmt.Errorf("NewReuqest failed: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Do failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
