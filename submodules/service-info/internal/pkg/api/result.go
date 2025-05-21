package api

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
)

// Result represents the result of a service operation.
type Result struct {
	Error  huma.StatusError
	Object any
	Status int
}

// SetError sets the error for the result.
func (r *Result) SetError(status int, msg string, err error) {
	r.Error = apierrors.APIErrorFor(status, msg, err)
	r.Status = status
}

// NotFoundError sets the error when a record is not found.
func (r *Result) NotFoundError(err error, id, model any) {
	r.SetError(
		http.StatusNotFound,
		fmt.Sprintf("unable to find %T with id: [%v]", model, id),
		err,
	)
}

// UnknownError sets the error when an unknown operation has occurred.
func (r *Result) UnknownError(err error) {
	r.InternalServerError("an unknown error has occurred", err)
}

// InternalServerError sets the error when there is an internal server error.
func (r *Result) InternalServerError(msg string, err error) {
	r.SetError(
		http.StatusInternalServerError,
		msg,
		err,
	)
}
