// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type User struct {
	Login   string
	HTMLURL string
}

const listUsersURL = "https://api.github.com/users"

type UsersListResult struct {
	Users    []*User
	nextLink string
}

func ListUsers() (*UsersListResult, error) {
	return listUsers(listUsersURL)
}

func listUsers(listURL string) (*UsersListResult, error) {
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

	var result UsersListResult

	result.nextLink, _ = parseLink(resp.Header.Get("Link"))

	if err := json.NewDecoder(resp.Body).Decode(&(result.Users)); err != nil {
		fmt.Printf("listURL = %q\n", listURL)
		return nil, err
	}
	return &result, nil
}

func (ul *UsersListResult) HasNext() bool {
	return ul.nextLink != ""
}

func (ul *UsersListResult) Next() (*UsersListResult, error) {
	if ul.nextLink == "" {
		panic("NextLink is not available")
	}

	return listUsers(ul.nextLink)
}

func (ul *UsersListResult) PrintAsHTMLTable(w io.Writer) {
	var userList = template.Must(template.New("userList").Parse(`
	<h1>user</h1>
	<table>
	<tr style='text-align: left'>
	<th>Name</th>
	</tr>
	{{range .Users}}
	<tr>
	  <td>{{.Login}}</td>
	</tr>
	{{end}}
	</table>
	`))

	if err := userList.Execute(w, ul); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
