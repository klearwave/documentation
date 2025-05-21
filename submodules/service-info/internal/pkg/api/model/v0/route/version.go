package route

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

const (
	DefaultVersionsPath     = "versions"
	DefaultVersionsGroupTag = "Versions"
)

// CreateVersion defines the routing and subsequent specification for POST to /api/v0/versions.
func CreateVersion() huma.Operation {
	return huma.Operation{
		OperationID:   "createVersion",
		Summary:       "Create a new Version.",
		Description:   "Create a new Version.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          versionGroupTags(),
	}
}

// GetVersion defines the routing and subsequent specification for GET to /api/v0/versions/{id}.
func GetVersion() huma.Operation {
	return huma.Operation{
		OperationID:   "getVersion",
		Summary:       "Get specific version information.",
		Description:   "Get specific version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{id}",
		Tags:          versionGroupTags(),
	}
}

// ListVersions defines the routing and subsequent specification for GET to /api/v0/versions.
func ListVersions() huma.Operation {
	return huma.Operation{
		OperationID:   "listVersions",
		Summary:       "List all version information.",
		Description:   "List all version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          versionGroupTags(),
	}
}

// ListVersionContainerImages defines the routing and subsequent specification for GET to /api/v0/versions/{id}/container_images.
func ListVersionContainerImages() huma.Operation {
	return huma.Operation{
		OperationID:   "listVersionContainerImages",
		Summary:       "List container images for a specific version.",
		Description:   "List container images for a specific version.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{id}/container_images",
		Tags:          versionGroupTags(),
	}
}

// DeleteVersion defines the routing and subsequent specification for DELETE to /api/v0/versions/{id}.
func DeleteVersion() huma.Operation {
	return huma.Operation{
		OperationID:   "deleteVersion",
		Summary:       "Delete a specific version.",
		Description:   "Delete a specific version.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{id}",
		Tags:          versionGroupTags(),
	}
}

// versionGroupTags defines the group tags that group similar APIs in the documentation site.
func versionGroupTags() []string {
	return []string{DefaultVersionsGroupTag}
}
