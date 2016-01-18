// Copyright © 2016 Yoshiki Shibata

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Edit(repo, title, body, state string, issueNo int, user *Credentials) (*Issue, error) {
	issue := EditIssue{title, body, state}
	b, err := json.Marshal(&issue)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	restAPIURL := issuesURL(repo) + fmt.Sprintf("/%d", issueNo)
	fmt.Printf("Rest API URL = %s\n", restAPIURL)
	req, err := newRequest("PATCH", restAPIURL, bytes.NewReader(b), user)
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

	var editedIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&editedIssue); err != nil {
		return nil, err
	}

	return &editedIssue, nil
}
