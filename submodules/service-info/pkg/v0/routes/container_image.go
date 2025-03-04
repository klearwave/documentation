package routes

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	v0 "github.com/klearwave/service-info/pkg/v0"
	"github.com/klearwave/service-info/pkg/v0/models"
)

const (
	DefaultContainerImagesPath     = "container_images"
	DefaultContainerImagesGroupTag = "ContainerImages"
)

// CreateContainerImage defines the routing and subsequent specification for POST to /api/v0/container_images.
func CreateContainerImage(input *models.ContainerImage) huma.Operation {
	return huma.Operation{
		OperationID:   "createContainerImage",
		Summary:       "Create a new container image.",
		Description:   "Create a new container image.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultContainerImagesPath),
		Tags:          containerImagesGroupTags(),
	}
}

// GetContainerImage defines the routing and subsequent specification for GET to /api/v0/container_images/{id}.
func GetContainerImage(input *models.ContainerImage) huma.Operation {
	return huma.Operation{
		OperationID:   "getContainerImage",
		Summary:       "Get specific container image information.",
		Description:   "Get specific container image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of container image to get.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: containerImagesGroupTags(),
	}
}

// GetContainerImages defines the routing and subsequent specification for GET to /api/v0/container_images.
func GetContainerImages(input *models.ContainerImage) huma.Operation {
	return huma.Operation{
		OperationID:   "getContainerImages",
		Summary:       "Get all container image information.",
		Description:   "Get all container image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath),
		Tags:          containerImagesGroupTags(),
	}
}

// DeleteContainerImage defines the routing and subsequent specification for DELETE to /api/v0/container_images/{container_image_id}.
func DeleteContainerImage(input *models.ContainerImage) huma.Operation {
	return huma.Operation{
		OperationID:   "deleteContainerImage",
		Summary:       "Delete a specific container image.",
		Description:   "Delete a specific container image.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of container image to delete.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: containerImagesGroupTags(),
	}
}

// GetContainerImageVersions defines the routing and subsequent specification for GET to /api/v0/container_images/{id}/versions.
func GetContainerImageVersions(input *models.ContainerImage) huma.Operation {
	return huma.Operation{
		OperationID:   "getContainerImageVersions",
		Summary:       "Get versions for a specific container image.",
		Description:   "Get versions for a specific container image.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}/versions",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Container Image ID to get versions for.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: versionInfoTags(),
	}
}

// containerImagesGroupTags defines the group tags that group similar APIs in the documentation site.
func containerImagesGroupTags() []string {
	return []string{
		DefaultContainerImagesGroupTag,
		v0.DefaultGroupTag,
	}
}
