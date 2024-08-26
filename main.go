package main

import (
	"fmt"
	"os"

	"github.com/kon3gor/newbiebot/internal/github"
)

const (
	tokenVar = "GITHUB_TOKEN"
	envVar   = "ENVIRONMENT"
)

func main() {
	c := github.Config{
		Token: os.Getenv(tokenVar),
	}

	client := github.NewClient(c)

	issues, err := client.GetGoodFirstIssues("LadybirdBrowser", "ladybird")
	if err != nil {
		panic(err)
	}

	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}
