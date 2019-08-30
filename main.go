package main

import (
	"github.com/jasonaibrahim/ozone-platform/app"
	"github.com/jasonaibrahim/ozone-platform/server"
	"log"
)

func main() {
	err := app.Init()

	if err != nil {
		log.Fatal("An error occurred during app initialization")
	}

	server.Start()
}
