package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/app/routes"
)

func LoadRoutes(router *gin.Engine) {
	router.GET("/_ah/health", routes.HealthCheck)
	router.Any("/callback", routes.OAuthCallback)
	router.GET("/login", routes.LoginHandler)
	router.GET("/profile", routes.Profile)
	router.GET("/", routes.HomeHandler)
}
