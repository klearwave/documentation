package list

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

// VersionRequest represents the request when listing versions.
type VersionRequest struct{}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req VersionRequest) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req VersionRequest) IsValid() (bool, error) {
	return true, nil
}

// ToLister converts the request to a lister object.  It is used to satisfy the
// Lister interface.
func (req VersionRequest) ToLister() api.Lister {
	return &v0.Versions{}
}
