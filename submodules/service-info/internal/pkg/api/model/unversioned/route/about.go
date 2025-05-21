package route

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/internal/pkg/api"
)

const (
	DefaultAboutPath     = "about"
	DefaultAboutGroupTag = "About"
)

// GetAbout defines the routing and subsequent specification for GET to /api/about.
func GetAbout() huma.Operation {
	return huma.Operation{
		OperationID:   "getAbout",
		Summary:       "Get overall information about the service.",
		Description:   "Get overall information about the service.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          api.PathFor(DefaultAboutPath),
		Tags:          []string{DefaultAboutGroupTag},
	}
}
