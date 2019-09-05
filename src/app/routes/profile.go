package routes

import (
	"github.com/gin-gonic/gin"
	app2 "github.com/jasonaibrahim/go-starter/src/app"
)

func Profile(c *gin.Context) {
	session, err := app2.Store.Get(c.Request, "auth-session")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"user": session.Values["profile"],
	})
}
