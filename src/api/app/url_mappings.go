package app

import (
	"../controllers/repositories"
)

func mapUrls() {
	router.POST("/repostiories", repositories.CreateRepo)
}
