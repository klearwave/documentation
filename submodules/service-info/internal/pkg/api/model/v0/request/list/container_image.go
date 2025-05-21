package list

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

// ContainerImageRequest represents the request when listing container images.
type ContainerImageRequest struct{}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req ContainerImageRequest) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req ContainerImageRequest) IsValid() (bool, error) {
	return true, nil
}

// ToLister converts the request to a lister object.  It is used to satisfy the
// Lister interface.
func (req ContainerImageRequest) ToLister() api.Lister {
	return &v0.ContainerImages{}
}
