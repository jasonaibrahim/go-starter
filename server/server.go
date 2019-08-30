package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/ozone-platform/app/handlers"
	"log"
)

func Start() {
	r := gin.Default()

	r.GET("/_ah/health", handlers.HealthCheck)

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
