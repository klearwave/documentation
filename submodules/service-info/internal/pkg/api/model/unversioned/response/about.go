package response

import "github.com/klearwave/service-info/internal/pkg/api/model/unversioned"

// About represents the response for returning information about the service.
type About struct {
	Body   unversioned.About
	Status int
}
