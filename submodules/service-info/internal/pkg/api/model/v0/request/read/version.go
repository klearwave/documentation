package read

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

// VersionRequest represents the request when reading a specific version.
type VersionRequest struct {
	model.StringFetcher
}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req VersionRequest) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req VersionRequest) IsValid() (bool, error) {
	return req.StringFetcher.IsValid()
}

// ToReader converts the request to a reader object.  It is used to satisfy the
// Reader interface.
func (req VersionRequest) ToReader() api.Reader {
	return &v0.Version{
		VersionBase: v0.VersionBase{
			ID: &req.ID,
		},
	}
}
