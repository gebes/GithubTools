package client

import (
	"context"
	"github.com/google/go-github/v44/github"
)

func CommitsWithFile(owner, repo string, file string) ([]*github.RepositoryCommit, *github.Response, error) {
	return client.Repositories.ListCommits(context.Background(), owner, repo, &github.CommitsListOptions{
		Path: file,
	})
}
