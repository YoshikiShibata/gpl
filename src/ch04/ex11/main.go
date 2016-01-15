// Copyright © 2016 Yoshiki Shibata

package main

import (
	"flag"
	"fmt"
	"os"

	"ch04/ex11/github"
)

// You need to get terminal package by the following command:
//  go get golang.org/x/crypto/ssh

// Usage
// 1. Create an issue
// 		issue [-create -title TITLE -body BODY] REPOSITORY
//		if -body is omitted, then an editor will be invoked.
//
// 2. Delete an issue
//		issue [-delete -issue ISSUE_NO] REPOSITORY
//
// 3. Print an issue
//		issue [-print -issue ISSUE_NO] REPOSITORY
//
// 4. Edit an issue
// 		issue [-edit -issue ISSUE_NO -title TITLE -body BODY -state [open | close]] REPOSITORY

var issueNo = flag.Int("issue", 0, "issue number")
var title = flag.String("title", "", "title for the issue")
var body = flag.String("body", "", "issue body")

var createFlag = flag.Bool("create", false, "create an issue")
var deleteFlag = flag.Bool("delete", false, "delete an issue")
var editFlag = flag.Bool("edit", false, "edit an issue")
var printFlag = flag.Bool("print", false, "print an issue")

var repository string

func main() {
	flag.Parse()
	validateOperationFlags()
	repository = validateAndGetRepository()

	var user github.Credentials
	user.Query()

	issue, err := github.Create(repository, *title, *body, &user)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(issue)
}

func validateOperationFlags() {
	flags := []bool{*createFlag, *deleteFlag, *editFlag, *printFlag}

	trueCount := 0
	for _, flag := range flags {
		if flag {
			trueCount++
		}
	}

	if trueCount == 0 {
		fmt.Println("Operation flag(-create, -delete, -edit, -print) is not specified")
		os.Exit(1)
	}
	if trueCount >= 2 {
		fmt.Println("Too many operation flags are specified")
	}
}

func validateAndGetRepository() string {
	if flag.NArg() == 0 {
		fmt.Println("REPOSITORY is not specified")
		os.Exit(1)
	}
	if flag.NArg() != 1 {
		fmt.Println("Too Many argument")
		os.Exit(1)
	}
	return flag.Arg(0)
}
