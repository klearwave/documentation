package errors

import "errors"

var (
	ErrInvalidVersion   = errors.New("invalid version")
	ErrMissingVersionId = errors.New("missing version_id parameter")
)
