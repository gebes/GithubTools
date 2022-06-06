package main

import (
	"GithubTools/pkg/client"
	"GithubTools/pkg/logger"
	"strings"
)

func main() {
	user := "torvalds"
	logger.Info.Printf("Running follower check for \"%s\"", user)

	users, _, err := client.Followers(user)
	if err != nil {
		logger.Error.Fatalln("Could not run follow back check:", err)
	}
	var names []string
	for _, g := range users {
		names = append(names, *g.Login)
	}
	logger.Info.Printf("Got %d followers as a result: %s", len(users), strings.Join(names, ", "))
}
