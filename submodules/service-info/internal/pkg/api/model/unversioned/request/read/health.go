package read

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned"
)

// Health represents the request for getting information about the service.
type Health struct{}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the
// Request interface.
func (req Health) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the
// Request interface.
func (req Health) IsValid() (bool, error) {
	return true, nil
}

// ToReader converts the request to a reader object.  It is used to satisfy the
// Reader interface.
func (req Health) ToReader() api.Reader {
	return &unversioned.Health{}
}
