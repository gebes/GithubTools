package client

import (
	"context"
	"github.com/google/go-github/v44/github"
)

func Followers(user string) ([]*github.User, *github.Response, error) {
	return client.Users.ListFollowers(context.Background(), user, nil)
}

func Following(user string) ([]*github.User, *github.Response, error) {
	return client.Users.ListFollowing(context.Background(), user, nil)
}
