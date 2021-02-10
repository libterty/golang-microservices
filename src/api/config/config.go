package config

import (
	"os"
)

const (
	apiGithubAccessToken = "Github_Token"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
