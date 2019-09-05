package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/app"
	"github.com/jasonaibrahim/go-starter/src/app/errors"
	"net/http"
	"os"
)

func HomeHandler(c *gin.Context) {
	sessionName := os.Getenv("SESSION_NAME")
	session, err := app.Store.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(errors.ErrorMessage(errors.UnknownError, err))
		return
	}

	c.HTML(http.StatusOK, "home/other", gin.H{
		"profile": session.Values["profile"],
		"title":   "FOO!!!",
	})
}
