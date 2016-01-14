// Copyright © 2016 Yoshiki Shibata

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createIssue(repo string, issue *CreateIssue) (*Issue, error) {
	b, err := json.Marshal(issue)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	fmt.Printf("json = %s\n", string(b))
	req, err := http.NewRequest("POST", GitHubAPIURL+"/repos/"+repo, bytes.NewReader(b))
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	// req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create failed: %s", resp.Status)
	}

	var createdIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&createdIssue); err != nil {
		return nil, err
	}
	return &createdIssue, nil
}
