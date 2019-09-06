package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/app"
	"net/http"
	"os"
)

func HomeHandler(c *gin.Context) {
	sessionName := os.Getenv("SESSION_NAME")
	session, err := app.Store.Get(c.Request, sessionName)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	c.HTML(http.StatusOK, "home", gin.H{
		"meta": gin.H{
			"keywords":    "foo bar baz",
			"author":      "jasonaibrahim",
			"description": "Go Starter Seed",
		},
		"title":   "Go Starter",
		"profile": session.Values["profile"],
	})
}

// TODO: how to apply common meta, title, and session to share across all controllers?

func OtherHandler(c *gin.Context) {
	sessionName := os.Getenv("SESSION_NAME")
	session, err := app.Store.Get(c.Request, sessionName)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	c.HTML(http.StatusOK, "home/other", gin.H{
		"meta": gin.H{
			"keywords":    "foo bar baz",
			"author":      "jasonaibrahim",
			"description": "Go Starter Seed",
		},
		"title":   "Go Starter",
		"profile": session.Values["profile"],
	})
}
