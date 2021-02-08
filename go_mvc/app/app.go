package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func init()  {
	router = gin.Default()
}

func StartApp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mapUrls()

	if err := router.Run(":5050"); err != nil {
		//log.Fatalf("App Stop with: %v", err)
		panic(err)
	}
}
