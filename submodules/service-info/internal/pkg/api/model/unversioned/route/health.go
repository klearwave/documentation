package route

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

const (
	DefaultHealthPath     = "healthz"
	DefaultHealthGroupTag = "Health"
)

// HealthZ defines the routing and subsequent specification for GET to /healthz.
func HealthZ() huma.Operation {
	return huma.Operation{
		OperationID:   "healthz",
		Summary:       "Get health information about the service.",
		Description:   "Get health information about the service.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          "/" + DefaultHealthPath,
		Tags:          []string{DefaultHealthGroupTag},
	}
}
