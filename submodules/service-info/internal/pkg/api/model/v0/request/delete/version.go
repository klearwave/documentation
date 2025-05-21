package delete

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

// VersionRequest represents the request when deleting a version.
type VersionRequest struct {
	api.Authorization
	model.StringFetcher
}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req VersionRequest) IsAuthorized() (bool, error) {
	return req.Authorized()
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req VersionRequest) IsValid() (bool, error) {
	return req.StringFetcher.IsValid()
}

// ToDeleter converts the request to a deleter object.  It is used to satisfy the
// Deleter interface.
func (req VersionRequest) ToDeleter() api.Deleter {
	return &v0.Version{
		VersionBase: v0.VersionBase{
			ID: &req.ID,
		},
	}
}
