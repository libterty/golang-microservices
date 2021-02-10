package app

import (
	"../controllers/repositories"
	"../controllers/health"
)

func mapUrls() {
	router.GET("/health", health.HealthCheck)
	router.POST("/repostiories", repositories.CreateRepo)
}
