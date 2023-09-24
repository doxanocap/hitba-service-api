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

	ErrSuchServiceAlreadyExist = errs.NewHttp(http.StatusConflict, "service with such name already exists")
)
