package env

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	GithubAccessToken string
	Key               string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	GithubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	Key = os.Getenv("KEY")
}
