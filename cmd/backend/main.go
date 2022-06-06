package main

import (
	"GithubTools/pkg/logger"
	"GithubTools/pkg/router"
	"GithubTools/pkg/signals"
	"net/http"
	"os"
)

func main() {
	signals.ListenForSigterm()

	logger.Info.Println("Starting to listen on port 8080 on PID", os.Getpid())
	err := router.Listen()
	if err != nil && err != http.ErrServerClosed {
		logger.Error.Fatalln("Could not start the router:", err)
	}

	signals.WaitForCleanup()
}
