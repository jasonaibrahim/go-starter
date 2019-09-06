package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/app/controllers"
	"log"
	"os"
	"path/filepath"
)

func LoadRoutes(router *gin.Engine) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	// uncomment the following block if you want to serve a single page app using a catchall route. e.g. angular
	//router.NoRoute(func(c *gin.Context) {
	//	c.File(filepath.Join(dir, "public", "index.html"))
	//})

	// catch errors (such as 404) and present a suitable page
	router.Use(controllers.StandardErrorHandler())

	// use the default logger
	router.Use(gin.Logger())

	// serve static files out of the `public` directory
	router.Static("/public", filepath.Join(dir, "public"))

	// define the routes for the application

	// health check
	router.GET("/_ah/health", controllers.HealthCheck)
	// oauth2 routes
	router.Any("/callback", controllers.OAuthCallback)
	router.GET("/login", controllers.LoginHandler)
	router.GET("/profile", controllers.Profile)
	// all other routes
	router.GET("/", controllers.HomeHandler)
	router.GET("/other", controllers.OtherHandler)
}
