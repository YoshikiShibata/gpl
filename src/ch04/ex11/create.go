// Copyright © 2016 Yoshiki Shibata

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createIssue(repo string, issue *CreateIssue, user *credentials) (*Issue, error) {
	b, err := json.Marshal(issue)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	fmt.Printf("json = %s\n", string(b))
	restAPIURL := GitHubAPIURL + "/repos/" + repo + "/issues"
	fmt.Printf("Rest API URL = %s\n", restAPIURL)
	req, err := http.NewRequest("POST", restAPIURL, bytes.NewReader(b))
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth(user.username, user.password)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Do failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create failed: %s", resp.Status)
	}

	var createdIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&createdIssue); err != nil {
		return nil, err
	}

	return &createdIssue, nil
}
