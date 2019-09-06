package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	app2 "github.com/jasonaibrahim/go-starter/src/app"
	auth2 "github.com/jasonaibrahim/go-starter/src/config/auth"
	"net/http"
	"os"
)

func LoginHandler(c *gin.Context) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	state := base64.StdEncoding.EncodeToString(b)
	sessionName := os.Getenv("SESSION_NAME")
	session, err := app2.Store.Get(c.Request, sessionName)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	session.Values["state"] = state
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	authenticator, err := auth2.NewAuthenticator()
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, authenticator.Config.AuthCodeURL(state))
}
