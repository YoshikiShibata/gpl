// Copyright © 2016 Yoshiki Shibata

package main

import "fmt"

// You need to get terminal package by the following command:
//  go get golang.org/x/crypto/ssh

// Usage
// 1. Create an issue
// 		issue REPOSITORY [-create -title TITLE -body BODY]
//		if -body is omitted, then an editor will be invoked.
//
// 2. Delete an issue
//		issue REPOSITORY [-delete -issue ISSUE_NO]
//
// 3. Print an issue
//		issue REPOSITORY [-print -issue ISSUE_NO]
//
// 4. Edit an issue
// 		issue REPOSITORY [-edit -issue ISSUE_NO -title TITLE -body BODY -state [open | close]]

func main() {
	var user credentials
	user.Query()

	issue := CreateIssue{
		"issue create test",
		"This is a bug"}

	iss, err := createIssue("YoshikiShibata/gpltest", &issue, &user)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(iss)
}
