// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Copyright © 2016 Yoshiki Shibata

// Issues prints a table of GitHub issues matching the search terms.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var months []*github.Issue
	var years []*github.Issue
	var others []*github.Issue
	now := time.Now()

	for _, item := range result.Items {
		if LessThanMonth(item.CreatedAt, now) {
			months = append(months, item)
			continue
		}

		if LessThanYear(item.CreatedAt, now) {
			years = append(years, item)
			continue
		}

		others = append(others, item)
	}

	showIssues(months, "===== less than a month old =====")
	showIssues(years, "===== less than a year old =====")
	showIssues(others, "===== more than a year old =====")
}

func showIssues(issues []*github.Issue, header string) {
	if len(issues) > 0 {
		fmt.Printf("%s\n", header)
		for _, item := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

/*

$./issue repo:golang/go is:open json decoder
12 issues:
===== less than a year old =====
#11046     kurin encoding/json: Decoder internally buffers full input
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#13558  ajwerner io: MultiReader should be more efficient when chained m
===== more than a year old =====
#5680    eaigner encoding/json: set key converter on en/decoder
#8658  gopherbot encoding/json: use bufio
#5901        rsc encoding/json: allow override type marshaling
#7872  extempora encoding/json: Encoder internally buffers full output
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional

*/
