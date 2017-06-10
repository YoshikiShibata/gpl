// Copyright © 2016, 2017 Yoshiki Shibata. All rights reserved.

package main

import (
	"ch04/ex14/github"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// This web server caches following information for a specific repository:
//   issues
//   milestones
// To get this information, the "/:owner/:repo" format is used as owner and repository:
// e.g.) localhost:8000/golang/go  the go repository
//
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

type RepoPath struct {
	owner string
	repo  string
}

type RepoInfo struct {
	issues     *github.IssuesListResult
	milestones *github.MilestonesListResult
}

var repoInfoCaches = make(map[RepoPath]RepoInfo)

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL.Path = %q\n", r.URL.Path)

	if r.URL.Path == "/" {
		showUsage(w)
		return
	}
	paths := splitPath(r.URL.Path)

	// paths should be as following:
	//  owner/repo
	//  owner/repo/issue/XXX     where XXX is a issue number
	//  owner/repo/milestone/YYY where YYY is the tile of Milestone
	if len(paths) != 2 && len(paths) != 4 {
		showUsage(w)
	}

	owner := paths[0]
	repo := paths[1]

	repoPath := RepoPath{owner, repo}

	switch len(paths) {
	case 2:
		showTopLevelPage(w, repoPath, repoInfoCaches)
	case 4:
		if paths[2] == "issue" {
			repoInfo, ok := repoInfoCaches[repoPath]
			if !ok {
				fmt.Fprintln(w, "Not Cached\n")
				return
			}
			issueNo, err := strconv.Atoi(paths[3])
			if err != nil {
				fmt.Fprintln(w, "%v\n", err)
				return
			}
			repoInfo.issues.PrintIssue(w, issueNo)
			return
		} else if paths[2] == "milestone" {
			repoInfo, ok := repoInfoCaches[repoPath]
			if !ok {
				fmt.Fprintln(w, "Not Cached\n")
				return
			}
			repoInfo.milestones.PrintMilestone(w, paths[3])
			return
		}
	}
}

func showUsage(w io.Writer) {
	fmt.Fprintln(w, "Please specify the owner and its repository. For example,")
	fmt.Fprintln(w, " localhost:8000/golang/go")
	fmt.Fprintln(w, `where "golang" is the owner and "go" is its repository.`)
}

func splitPath(path string) []string {
	paths := strings.Split(path[1:], "/") // Not that the first / is deleted
	len := len(paths)
	if paths[len-1] == "" {
		return paths[:len-1]
	}
	return paths
}

func showTopLevelPage(w io.Writer, path RepoPath, caches map[RepoPath]RepoInfo) {
	repoInfo, ok := caches[path]
	if !ok {
		var err error
		repoInfo, err = cacheRepoInfo(path, caches)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
			return
		}
	}

	repoInfo.issues.PrintAsHTMLTable(w)
	repoInfo.milestones.PrintAsHTMLTable(w)
	users, err := github.ListUsers()
	if err != nil {
		fmt.Fprintf(w, "Users Could Not Be Obtained: %v\n", err)
	}
	users.PrintAsHTMLTable(w)
}

func cacheRepoInfo(path RepoPath, caches map[RepoPath]RepoInfo) (RepoInfo, error) {
	var repoInfo RepoInfo

	issues, err := github.ListIssues(path.owner, path.repo)
	if err != nil {
		return RepoInfo{}, err
	}
	repoInfo.issues = issues

	milestones, err := github.ListMilestones(path.owner, path.repo)
	if err != nil {
		return RepoInfo{}, err
	}
	repoInfo.milestones = milestones
	caches[path] = repoInfo

	return repoInfo, nil
}
