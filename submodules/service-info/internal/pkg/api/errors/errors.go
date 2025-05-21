package errors

import (
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

var ErrMissingParameterID = errors.New("missing required id parameter")

// APIErrorFor is a helper function to return an error from a status code.
func APIErrorFor(statusCode int, msg string, err error) huma.StatusError {
	switch statusCode {
	case http.StatusBadRequest:
		return huma.Error400BadRequest(msg, err)
	case http.StatusUnauthorized:
		return huma.Error401Unauthorized(msg, err)
	case http.StatusForbidden:
		return huma.Error403Forbidden(msg, err)
	case http.StatusNotFound:
		return huma.Error404NotFound(msg, err)
	default:
		return huma.Error500InternalServerError(msg, err)
	}
}
