package main

import (
	"GithubTools/pkg/check"
	"GithubTools/pkg/logger"
	"strings"
)

func main() {
	user := "Gebes"
	logger.Info.Printf("Running follow back check for \"%s\"", user)

	users, err := check.NoFollowBacks(user)
	if err != nil {
		logger.Error.Fatalln("Could not run follow back check:", err)
	}
	var names []string
	for _, g := range users {
		names = append(names, *g.Login)
	}
	logger.Info.Printf("Got %d users as a result: %s", len(users), strings.Join(names, ", "))
}
