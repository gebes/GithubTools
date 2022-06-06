package signals

import (
	"GithubTools/pkg/logger"
	"GithubTools/pkg/router"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan bool, 1)

func ListenForSigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Debug.Println("Received signal", sig)

		err := router.Shutdown()
		if err != nil {
			logger.Error.Println("Could not stop router:", err)
		}
		logger.Info.Println("Closed router")

		logger.Info.Println("Stopping GithubTools backend")
		done <- true
	}()

}

func WaitForCleanup() {
	<-done
}
