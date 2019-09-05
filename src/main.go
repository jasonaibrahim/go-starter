package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/app"
	"github.com/jasonaibrahim/go-starter/src/config"
	"log"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatal("An error occurred during app initialization")
	}

	err = startServer()
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	router := gin.Default()

	err := config.LoadHtmlTemplates(router)
	if err != nil {
		return err
	}

	config.LoadRoutes(router)

	return router.Run()
}
