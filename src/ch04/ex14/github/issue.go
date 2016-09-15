package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

const listIssuesURL = "https://api.github.com/repos/%s/%s/issues"

type IssuesListResult struct {
	Issues   []*Issue
	nextLink string
	lastLink string
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // (Markdown format)
}

func ListIssues(owner, repo string) (*IssuesListResult, error) {
	listURL := fmt.Sprintf(listIssuesURL, owner, repo)
	fmt.Printf("listURL = %s\n", listURL)
	return listIssues(listURL)
}

func listIssues(listURL string) (*IssuesListResult, error) {
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

	var result IssuesListResult

	result.nextLink, result.lastLink = parseLink(resp.Header.Get("Link"))

	if err := json.NewDecoder(resp.Body).Decode(&(result.Issues)); err != nil {
		fmt.Printf("listURL = %q\n", listURL)
		return nil, err
	}
	return &result, nil
}

func (il *IssuesListResult) HasNext() bool {
	return il.nextLink != ""
}

func (il *IssuesListResult) Next() (*IssuesListResult, error) {
	if il.nextLink == "" {
		panic("NextLink is not available")
	}

	return listIssues(il.nextLink)
}

func (il *IssuesListResult) PrintAsHTMLTable(w io.Writer) {
	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>issues</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	{{range .Issues}}
	<tr>
	  <td><a href='issue/{{.Number}}'>{{.Number}}</a></td>
	  <td>{{.State}}</td>
	  <td>{{.User.Login}}</td>
	  <td><a href='issue/{{.Number}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	if err := issueList.Execute(w, il); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}

func (il *IssuesListResult) PrintIssue(w io.Writer, issueNo int) {
	for _, issue := range il.Issues {
		if issue.Number == issueNo {
			printIssue(w, issue)
			return
		}
	}
	fmt.Fprintf(w, "Issue %d Not Found\n", issueNo)
}

func printIssue(w io.Writer, issue *Issue) {
	fmt.Fprintf(w, "#%d State: %s User: %s\n",
		issue.Number, issue.State, issue.User.Login)
	fmt.Fprintf(w, "Title: %s\n", issue.Title)
	fmt.Fprintf(w, "Body: %s\n", issue.Body)
}
