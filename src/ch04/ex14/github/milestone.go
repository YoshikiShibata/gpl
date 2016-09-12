package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Milestone struct {
	State        string
	Title        string
	Description  string
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}

const milestonesURL = "https://api.github.com/repos/%s/%s/milestones"

type MilestonesListResult struct {
	Milestones []Milestone
	nextLink   string
	lastLink   string
}

func ListMilestones(owner, repo string) (*MilestonesListResult, error) {
	listURL := fmt.Sprintf(milestonesURL, owner, repo)
	return listMilestones(listURL)
}

func listMilestones(listURL string) (*MilestonesListResult, error) {
	resp, err := http.Get(listURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, parseBadRequest(resp)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status Code is %d", resp.StatusCode)
	}

	var result MilestonesListResult

	result.nextLink, result.lastLink = parseLink(resp.Header.Get("Link"))

	if err := json.NewDecoder(resp.Body).Decode(&(result.Milestones)); err != nil {
		fmt.Printf("listURL = %q\n", listURL)
		return nil, err
	}
	return &result, nil
}

func (ml *MilestonesListResult) HasNext() bool {
	return ml.nextLink != ""
}

func (ml *MilestonesListResult) Next() (*MilestonesListResult, error) {
	if ml.nextLink == "" {
		panic("NextLink is not available")
	}

	return listMilestones(ml.nextLink)
}
