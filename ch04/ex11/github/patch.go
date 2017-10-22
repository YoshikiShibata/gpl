// Copyright © 2016 Yoshiki Shibata

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func patch(repo string, issue interface{}, issueNo int, user *Credentials) (*Issue, error) {
	jsonBody, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	restAPIURL := issuesURL(repo) + fmt.Sprintf("/%d", issueNo)
	fmt.Printf("Rest API URL = %s\n", restAPIURL)
	req, err := newRequest("PATCH", restAPIURL, bytes.NewReader(jsonBody), user)
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

	var patchedIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&patchedIssue); err != nil {
		return nil, err
	}

	return &patchedIssue, nil
}
