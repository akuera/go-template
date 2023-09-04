package handlers

import (
	"github.com/akuera/go-template/errors"
	"net/http"
)

var (
	errorToHttpResponse = map[error]int{
		nil:                   http.StatusOK,
		errors.ErrNotFound:    http.StatusNotFound,
		errors.ErrInvalidCode: http.StatusBadRequest,
		errors.ErrInternal:    http.StatusInternalServerError,
		// Add domain errors
	}
)

func httpStatus(err error) int {
	if status, ok := errorToHttpResponse[err]; ok {
		return status
	}
	return http.StatusInternalServerError
}
