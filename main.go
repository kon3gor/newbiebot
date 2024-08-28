package main

import (
	"fmt"
	"os"

	"github.com/kon3gor/newbiebot/internal/github"
	"github.com/kon3gor/newbiebot/internal/seloutil"
	"github.com/kon3gor/newbiebot/internal/webhook"
	"github.com/kon3gor/newbiebot/internal/ydbrepo"
	"github.com/kon3gor/selo"
)

const (
	tokenVar = "GITHUB_TOKEN"
	envVar   = "ENVIRONMENT"
)

func main() {
	selo.Init()

	c := github.Config{
		Token: os.Getenv(tokenVar),
	}

	selo.
		Unique[*github.GithubClient]().
		SetLazy(true).
		SetFactory(seloutil.Wrap(c, github.NewClient)).
		Record()

	selo.
		Unique[webhook.Repo]().
		SetLazy(true).
		SetFactory(webHookRepoFactory).
		Record()

	selo.
		Unique[*webhook.WebhookManager]().
		SetLazy(true).
		SetFactory(seloutil.Wrap(webhook.Config{}, webhook.NewManager)).
		Record()

	client := selo.Get[*github.GithubClient]()

	issues, err := client.GetGoodFirstIssues("LadybirdBrowser", "ladybird")
	if err != nil {
		panic(err)
	}

	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}

func webHookRepoFactory() webhook.Repo {
	return ydbrepo.NewRepo()
}

