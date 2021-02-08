package main

import (
	"./app"
	"log"
)

// main services
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app.StartApp()
}
