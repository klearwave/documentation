package errors

import "errors"

var (
	ErrInvalidVersion = errors.New("invalid version object")

	ErrMissingVersionObject             = errors.New("missing version object")
	ErrMissingVersionParameterVersionID = errors.New("missing required version_id parameter for version")
	ErrMissingVersionParameterID        = errors.New("missing required id parameter for version")
)
