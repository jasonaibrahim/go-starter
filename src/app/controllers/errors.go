package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StandardErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if len(c.Errors) > 0 {
			err = c.Errors[0]
		}

		logError(err)

		switch c.Writer.Status() {
		case 400:
			c.HTML(RenderErrorTemplate(BadRequestError, err))
		case 404:
			c.HTML(RenderErrorTemplate(NotFoundError, err))
		case 500:
			c.HTML(RenderErrorTemplate(UnknownError, err))
		default:
			c.Next()
		}
	}
}

func RenderErrorTemplate(error interface{}, err error) (int, string, interface{}) {
	var appError *AppError

	switch error.(type) {
	case *AppError:
		appError = error.(*AppError)
	default:
		appError = UnknownError
	}

	status := error.(*AppError).Status
	title := error.(*AppError).Title
	message := error.(*AppError).Details
	href := error.(*AppError).Href
	code := error.(*AppError).Code

	logError(error)

	return appError.Status, "errors/error", gin.H{
		"status":  status,
		"title":   title,
		"message": message,
		"href":    href,
		"code":    code,
		"error":   err,
		"meta":    gin.H{},
	}
}

func logError(err interface{}) {
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
}

type AppError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Href    string `json:"href"`
}

func constructError(status int, code string, title string, details string, href string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Title:   title,
		Details: details,
		Href:    href,
	}
}

var (
	InvalidStateError = constructError(
		500,
		"invalid_state",
		"Invalid State",
		"Invalid state parameter",
		"",
	)
	IdTokenError = constructError(
		500,
		"invalid_id_token",
		"Invalid ID Token",
		"No id_token field in oauth2 token",
		"",
	)
	UnknownError = constructError(
		500,
		"unknown",
		"Internal Server Error",
		"An unexpected error has occurred.",
		"",
	)
	NotFoundError = constructError(
		404,
		"not_found",
		"Page Not Found",
		"We can't find that page :(",
		"",
	)
	BadRequestError = constructError(
		400,
		"bad_request",
		"Bad Request",
		"There was a problem with the request",
		"",
	)
)
