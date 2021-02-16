package app

import (
	"../controllers/health"
	"../controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.HealthCheck)
	router.POST("/repostiories", repositories.CreateRepo)
}
