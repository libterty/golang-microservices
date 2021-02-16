package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	apiGithubAccessToken = "Github_Token"
)

func GetGithubAccessToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(apiGithubAccessToken)
}
