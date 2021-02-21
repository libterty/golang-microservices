package app

import (
	"../controllers/health"
	"../controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.HealthCheck)
	router.POST("/repostiory", repositories.CreateRepo)
	router.POST("/repostiories", repositories.CreateRepos)
}
