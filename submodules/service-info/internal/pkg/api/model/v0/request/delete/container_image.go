package delete

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

// ContainerImageRequest represents the request when reading a specific container image.
type ContainerImageRequest struct {
	api.Authorization
	model.IntegerFetcher
}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req ContainerImageRequest) IsAuthorized() (bool, error) {
	return req.Authorized()
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req ContainerImageRequest) IsValid() (bool, error) {
	return req.IntegerFetcher.IsValid()
}

// ToDeleter converts the request to a deleter object.  It is used to satisfy the
// Deleter interface.
func (req ContainerImageRequest) ToDeleter() api.Deleter {
	return &v0.ContainerImage{
		WithID: model.WithID{
			ID: req.ID,
		},
	}
}
