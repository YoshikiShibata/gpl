package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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

func (ml *MilestonesListResult) PrintAsHTMLTable(w io.Writer) {
	var milestoneList = template.Must(template.New("milestoneList").Parse(`
	<h1>milestones</h1>
	<table>
	<tr style='text-align: left'>
	<th>Title</th>
	<th>State</th>
	</tr>
	{{range .Milestones}}
	<tr>
	  <td><a href='milestone/{{.Title}}'>{{.Title}}</a></td>
	  <td>{{.State}}</td>
	</tr>
	{{end}}
	</table>
	`))

	if err := milestoneList.Execute(w, ml); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}

func (ml *MilestonesListResult) PrintMilestone(w io.Writer, title string) {
	for _, ms := range ml.Milestones {
		if ms.Title == title {
			printMilestone(w, ms)
			return
		}
	}
	fmt.Fprintf(w, "Milestone %s Not Found\n", title)
}

func printMilestone(w io.Writer, ms Milestone) {
	fmt.Fprintf(w, "Title: %s State: %s\n", ms.Title, ms.State)
	fmt.Fprintf(w, "Open Issues: %d Closed Issues: %d\n", ms.OpenIssues, ms.ClosedIssues)
	fmt.Fprintf(w, "Description:\n%s\n", ms.Description)

}
