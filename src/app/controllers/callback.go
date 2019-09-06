package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jasonaibrahim/go-starter/src/app"
	"github.com/jasonaibrahim/go-starter/src/config/auth"
	"golang.org/x/oauth2"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
)

func getSession(request *http.Request) (*sessions.Session, error) {
	sessionName := os.Getenv("SESSION_NAME")
	return app.Store.Get(request, sessionName)
}

func isStateInvalid(request *http.Request, session *sessions.Session) bool {
	return request.URL.Query().Get("state") != session.Values["state"]
}

func getAccessToken(request *http.Request, authenticator *auth.Authenticator) (*oauth2.Token, error) {
	code := request.URL.Query().Get("code")
	return authenticator.Config.Exchange(context.TODO(), code)
}

func getProfile(idToken *oidc.IDToken) (map[string]interface{}, error) {
	var profile map[string]interface{}
	var err = idToken.Claims(&profile)
	return profile, err
}

func getIdToken(authenticator auth.Authenticator, idToken string) (*oidc.IDToken, error) {
	clientId := os.Getenv("AUTH0_CLIENT_ID")
	oidcConfig := &oidc.Config{ClientID: clientId}
	return authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), idToken)
}

func OAuthCallback(c *gin.Context) {
	session, err := getSession(c.Request)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	if isStateInvalid(c.Request, session) {
		c.HTML(RenderErrorTemplate(InvalidStateError, nil))
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	token, err := getAccessToken(c.Request, authenticator)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.HTML(RenderErrorTemplate(IdTokenError, nil))
		return
	}

	idToken, err := getIdToken(*authenticator, rawIDToken)
	if err != nil {
		c.HTML(RenderErrorTemplate(IdTokenError, err))
		return
	}

	profile, err := getProfile(idToken)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	err = updateSession(session, rawIDToken, token.AccessToken, profile, c.Request, c.Writer)
	if err != nil {
		c.HTML(RenderErrorTemplate(UnknownError, err))
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func updateSession(
	session *sessions.Session,
	idToken string,
	accessToken string,
	profile map[string]interface{},
	request *http.Request,
	writer http.ResponseWriter,
) error {
	session.Values["id_token"] = idToken
	session.Values["access_token"] = accessToken
	session.Values["profile"] = profile

	return session.Save(request, writer)
}
