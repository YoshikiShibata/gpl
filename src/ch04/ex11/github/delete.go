// Copyright © 2016 Yoshiki Shibata

package github

import (
	"fmt"
	"net/http"
	"os"
)

func Delete(repo string, issueNo int, user *Credentials) error {
	restAPIURL := issuesURL(repo) + fmt.Sprintf("/%d", issueNo)
	req, err := newRequest("DELETE", restAPIURL, nil, user)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("Do failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Sorry, Github doesn't support DELETE for issues\n")
		return fmt.Errorf("delete failed: %s", resp.Status)
	}

	return nil
}
