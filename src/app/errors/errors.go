package errors

import (
	"log"
	"runtime"
)

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Href    string `json:"href"`
}

func newAPIError(status int, code string, title string, details string, href string) *APIError {
	return &APIError{
		Status:  status,
		Code:    code,
		Title:   title,
		Details: details,
		Href:    href,
	}
}

type APIErrors struct {
	Errors []*APIError `json:"errors"`
}

func logErr(err error) {
	if err != nil {
		trace := make([]byte, 1024)
		runtime.Stack(trace, true)
		log.Printf("ERROR: %s\n%s", err, trace)
	}
}

func ErrorMessage(error interface{}, err error) (int, *APIErrors) {
	var apiErrors *APIErrors

	logErr(err)

	switch error.(type) {
	case *APIError:
		apiError := error.(*APIError)
		apiErrors = &APIErrors{
			Errors: []*APIError{apiError},
		}
	case *APIErrors:
		apiErrors = error.(*APIErrors)
	default:
		apiErrors = &APIErrors{
			Errors: []*APIError{UnknownError},
		}
	}
	return apiErrors.Status(), apiErrors
}

func (errors *APIErrors) Status() int {
	return errors.Errors[0].Status
}

var (
	InvalidStateError = newAPIError(
		500,
		"invalid_state",
		"Invalid State",
		"Invalid state parameter",
		"",
	)
	IdTokenError = newAPIError(
		500,
		"invalid_id_token",
		"Invalid ID Token",
		"No id_token field in oauth2 token",
		"",
	)
	UnknownError = newAPIError(
		500,
		"unknown",
		"Internal Server Error",
		"An unexpected error has occurred.",
		"",
	)
)
