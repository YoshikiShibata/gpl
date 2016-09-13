// Copyright © 2016 Yoshiki Shibata

package main

import (
	"ch04/ex14/github"
	"fmt"
	"os"
	"time"
)

func main() {
	iList, err := github.ListIssues(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	mList, err := github.ListMilestones(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	for i := 0; ; i++ {
		fmt.Printf("%#v\n", mList.Milestones)
		if !mList.HasNext() {
			return
		}
		mList, err = mList.Next()
	}

	for i := 0; ; i++ {
		fmt.Printf("%d : %d issues\n", i, len(iList.Issues))
		if !iList.HasNext() {
			return
		}
		iList, err = iList.Next()
		if err != nil {
			fmt.Printf("%#v\n", err)
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
}
