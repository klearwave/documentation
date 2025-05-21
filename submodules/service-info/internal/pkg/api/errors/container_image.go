package errors

import "errors"

var (
	ErrInvalidContainerImage = errors.New("invalid container image object")

	ErrMissingContainerImageObject              = errors.New("missing container image object")
	ErrMissingContainerImageParameterImage      = errors.New("missing required image parameter for container image")
	ErrMissingContainerImageParameterSHA256Sum  = errors.New("missing required sha256sum parameter for container image")
	ErrMissingContainerImageParameterCommitHash = errors.New("missing required commit_hash parameter for container image")

	ErrInvalidContainerImageParameterSHA256Sum  = errors.New("invalid sha256sum parameter for container image")
	ErrInvalidContainerImageParameterCommitHash = errors.New("invalid commit_hash parameter for container image")

	ErrMismatchParameter = errors.New("mismatch in container image parameter")
)
