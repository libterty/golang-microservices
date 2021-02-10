package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApplication() {
	mapUrls()
	if err := router.Run(":5052"); err != nil {
		panic(err)
	}
}
