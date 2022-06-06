package client

import (
	"GithubTools/pkg/env"
	"GithubTools/pkg/logger"
	"context"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func init() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: env.GithubAccessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	logger.Info.Println("Created GitHub http client")
	client = github.NewClient(tc)
	logger.Info.Println("Initialied GitHub client")
}
