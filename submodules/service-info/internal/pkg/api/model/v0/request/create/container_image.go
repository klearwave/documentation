package create

import (
	"fmt"

	"github.com/klearwave/service-info/internal/pkg/api"
	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
	"github.com/klearwave/service-info/internal/pkg/utils/validate"
)

// ContainerImageBody is the base set of inputs required to create a new Container Image.
type ContainerImageBody struct {
	v0.ContainerImageBase
}

// ContainerImageRequest is the base set of inputs required to create a ContainerImageRequest.
type ContainerImageRequest struct {
	api.Authorization
	Body ContainerImageBody
}

// IsAuthorized checks if the user is authorized to create a version.  It is used to satisfy the
// Request interface.
func (req *ContainerImageRequest) IsAuthorized() (bool, error) {
	return req.Authorized()
}

// IsValid checks if the request is valid for a version.  It is used to satisfy the
// Request interface.
func (req *ContainerImageRequest) IsValid() (bool, error) {
	allErrors := []error{}

	// validate sha256sum
	if req.Body.SHA256Sum == nil {
		return false, apierrors.ErrMissingContainerImageParameterSHA256Sum
	}

	if err := validate.SHA256Sum(*req.Body.SHA256Sum); err != nil {
		allErrors = append(
			allErrors,
			fmt.Errorf("%s; %w", err.Error(), apierrors.ErrInvalidContainerImageParameterSHA256Sum),
		)
	}

	// validate commit hash
	if req.Body.CommitHash == nil {
		return false, apierrors.ErrMissingContainerImageParameterCommitHash
	}

	if err := validate.CommitHash(*req.Body.CommitHash); err != nil {
		allErrors = append(
			allErrors,
			fmt.Errorf("%s; %w", err.Error(), apierrors.ErrInvalidContainerImageParameterCommitHash),
		)
	}

	if len(allErrors) > 0 {
		return false, fmt.Errorf("%+v", allErrors)
	}

	return true, nil
}

// ToCreator converts the request to a creator object.  It is used to satisfy the
// Creator interface.
func (req *ContainerImageRequest) ToCreator() api.Creator {
	return req.Body.ToObject()
}

// ToObject converts the request body to an object.
func (body ContainerImageBody) ToObject() *v0.ContainerImage {
	return &v0.ContainerImage{
		ContainerImageBase: v0.ContainerImageBase{
			Image:      body.Image,
			SHA256Sum:  body.SHA256Sum,
			CommitHash: body.CommitHash,
		},
	}
}
