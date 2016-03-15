// Issues print a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ch05/ex11/github"
)

func age(createdAt time.Time) string {
	since := time.Since(createdAt)
	days := since.Hours() / 24.0
	months := int(days / 30.0)
	years := int(days / 365.0)

	if months == 0 {
		return "lm"
	}
	if years == 0 {
		return "ly"
	}
	return "my"
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, age(item.CreatedAt), item.User.Login, item.Title)
	}
}
