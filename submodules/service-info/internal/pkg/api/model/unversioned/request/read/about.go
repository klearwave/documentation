package read

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned"
)

// About represents the request for getting information about the service.
type About struct{}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the
// Request interface.
func (req About) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the
// Request interface.
func (req About) IsValid() (bool, error) {
	return true, nil
}

// ToReader converts the request to a reader object.  It is used to satisfy the
// Reader interface.
func (req About) ToReader() api.Reader {
	return &unversioned.About{
		Version:    api.ServerVersion,
		CommitHash: api.CommitHash,
	}
}
