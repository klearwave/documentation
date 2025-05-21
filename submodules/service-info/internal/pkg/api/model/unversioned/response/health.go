package response

import "github.com/klearwave/service-info/internal/pkg/api/model/unversioned"

// Health represents the response for returning information about the service.
type Health struct {
	Body   unversioned.Health
	Status int
}
