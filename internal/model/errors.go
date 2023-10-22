package model

import (
	"github.com/doxanocap/pkg/errs"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// default http error responses
var (
	HttpInternalServerError = errs.NewHttp(http.StatusInternalServerError, "internal server error")
	HttpConflictError       = errs.NewHttp(http.StatusConflict, "conflict")
	HttpBadRequest          = errs.NewHttp(http.StatusBadRequest, "bad request")
	HttpUnauthorized        = errs.NewHttp(http.StatusUnauthorized, "unauthorized")
)

// custom errors for special cases
var (
	ErrSuchServiceAlreadyExist = errs.NewHttp(http.StatusConflict, "service with such name already exists")
	ErrServiceIdNotFound       = errs.NewHttp(http.StatusNotFound, "service with such id not found")
)
