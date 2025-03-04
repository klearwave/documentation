package routes

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	v0 "github.com/klearwave/service-info/pkg/v0"
	"github.com/klearwave/service-info/pkg/v0/models"
)

const (
	DefaultVersionsPath     = "versions"
	DefaultVersionsGroupTag = "Versions"
)

// CreateVersion defines the routing and subsequent specification for POST to /api/v0/versions.
func CreateVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "createVersion",
		Summary:       "Create a new Version.",
		Description:   "Create a new Version.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          versionInfoTags(),
	}
}

// GetVersion defines the routing and subsequent specification for GET to /api/v0/versions/{version_id}.
func GetVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "getVersion",
		Summary:       "Get specific version information.",
		Description:   "Get specific version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{version_id}",
		Parameters: []*huma.Param{
			{
				Name:        "version_id",
				Description: "Version ID to get.",
				Example:     "v0.1.2",
				In:          "path",
				Required:    true,
			},
		},
		Tags: versionInfoTags(),
	}
}

// GetVersions defines the routing and subsequent specification for GET to /api/v0/versions.
func GetVersions(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "getVersions",
		Summary:       "Get all version information.",
		Description:   "Get all version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          versionInfoTags(),
	}
}

// GetVersionContainerImages defines the routing and subsequent specification for GET to /api/v0/versions/{version_id}/container_images.
func GetVersionContainerImages(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "getVersionContainerImages",
		Summary:       "Get container images for a specific version.",
		Description:   "Get container images for a specific version.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{version_id}/container_images",
		Parameters: []*huma.Param{
			{
				Name:        "version_id",
				Description: "Version ID to get container images for.",
				Example:     "v0.1.2",
				In:          "path",
				Required:    true,
			},
		},
		Tags: versionInfoTags(),
	}
}

// DeleteVersion defines the routing and subsequent specification for DELETE to /api/v0/versions/{version_id}.
func DeleteVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "deleteVersion",
		Summary:       "Delete a specific version.",
		Description:   "Delete a specific version.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{version_id}",
		Parameters: []*huma.Param{
			{
				Name:        "version_id",
				Description: "Version ID to delete.",
				Example:     "v0.1.2",
				In:          "path",
				Required:    true,
			},
		},
		Tags: versionInfoTags(),
	}
}

// versionGroupTags defines the group tags that group similar APIs in the documentation site.
func versionInfoTags() []string {
	return []string{
		DefaultVersionsGroupTag,
		v0.DefaultGroupTag,
	}
}
