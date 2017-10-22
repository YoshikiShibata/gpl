// Copyright © 2016 Yoshiki Shibata

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"ch04/ex11/github"
)

// You need to get terminal package by the following command:
//  go get golang.org/x/crypto/ssh

// Usage
// 1. Create an issue
// 		issue [-create -title TITLE -body BODY] REPOSITORY
//		if -body is omitted, then an editor will be invoked.
//
// 2. Close an issue
//		issue [-close -issue ISSUE_NO] REPOSITORY
//
// 3. Print an issue
//		issue [-print -issue ISSUE_NO] REPOSITORY
//
// 4. Edit an issue
// 		issue [-edit -issue ISSUE_NO -title TITLE -body BODY] REPOSITORY

var issueNo = flag.Int("issue", 0, "issue number")
var title = flag.String("title", "", "title for the issue")
var body = flag.String("body", "", "issue body")

var createFlag = flag.Bool("create", false, "create an issue")
var closeFlag = flag.Bool("close", false, "close an issue")
var editFlag = flag.Bool("edit", false, "edit an issue")
var printFlag = flag.Bool("print", false, "print an issue")

var repository string

func main() {
	flag.Parse()
	validateOperationFlags()
	repository = validateAndGetRepository()

	var user github.Credentials
	user.Query()
	fmt.Println()

	switch true {
	case *createFlag:
		b := *body
		if !isFlagSpecified("body") {
			b = invokeEditor()
		}
		issue, err := github.Create(repository, *title, b, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v", issue)
		saveIssueNo(issue.Number)
	case *closeFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		issue, err := github.Close(repository, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v\n", issue)
	case *printFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		issue, err := github.Get(repository, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v\n", issue)
	case *editFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		b := *body
		if !isFlagSpecified("body") {
			b = invokeEditor()
		}
		issue, err := github.Edit(repository, *title, b, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", issue)
	}
}

func saveIssueNo(issueNo int) {
	f, err := os.Create("issue_no.txt")
	if err != nil {
		return
	}
	fmt.Fprintf(f, "%d", issueNo)
	f.Close()
}

func invokeEditor() string {
	f, err := ioutil.TempFile("", "body.")
	if err != nil {
		panic(fmt.Errorf("Cannot crete a temp file: %v", err))
	}

	name := f.Name()
	f.Close() // ignore Error

	cmd := exec.Command("vim", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(fmt.Errorf("Cannot invoke vim: %v", err))
	}

	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Errorf("Cannot read a temp file: %v", err))
	}
	return string(bytes)
}

func isFlagSpecified(name string) (specified bool) {
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			specified = true
		}
	})
	return
}

func validateOperationFlags() {
	flags := []bool{*createFlag, *closeFlag, *editFlag, *printFlag}

	trueCount := 0
	for _, flag := range flags {
		if flag {
			trueCount++
		}
	}

	if trueCount == 0 {
		fmt.Println("Operation flag(-create, -close, -edit, -print) is not specified")
		os.Exit(1)
	}
	if trueCount >= 2 {
		fmt.Printf("Too many operation flags are specified: %#v\n", flags)
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
