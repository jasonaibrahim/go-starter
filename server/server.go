package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ozone/ozone-platform/handlers"
	"log"
	"os"
)

func Start() {
	r := gin.Default()

	r.GET("/_ah/health", handlers.HealthCheck)

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Ozone Platform Server listening on %s", os.Getenv("PORT"))
	}
}
