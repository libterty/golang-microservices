package app

import (
	"../controllers"
)

func mapUrls()  {
	router.GET("/users/:user_id", controllers.GetUser)
}