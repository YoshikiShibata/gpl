// Copyright © 2016 Yoshiki Shibata

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

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
	user, password := credentials()

	fmt.Printf("user = %s, password = %s\n", user, password)

	panic("Not Implemented Yet")
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username for 'https://github.com': ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print(fmt.Sprintf("Password for 'https://%s@github.com': ", username))
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nPassword Error: %v\n", err)
	}
	password := string(bytePassword)

	return username, strings.TrimSpace(password)
}
