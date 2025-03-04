package routes

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	v0 "github.com/klearwave/service-info/pkg/v0"
	"github.com/klearwave/service-info/pkg/v0/models"
)

const (
	DefaultNodeImagesPath     = "node_images"
	DefaultNodeImagesGroupTag = "NodeImages"
)

// CreateNodeImage defines the routing and subsequent specification for POST to /api/v0/container_images.
func CreateNodeImage(input *models.NodeImage) huma.Operation {
	return huma.Operation{
		OperationID:   "createNodeImage",
		Summary:       "Create a new node image.",
		Description:   "Create a new node image.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultNodeImagesPath),
		Tags:          nodeImagesGroupTags(),
	}
}

// GetNodeImage defines the routing and subsequent specification for GET to /api/v0/container_images/{id}.
func GetNodeImage(input *models.NodeImage) huma.Operation {
	return huma.Operation{
		OperationID:   "getNodeImage",
		Summary:       "Get specific node image information.",
		Description:   "Get specific node image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultNodeImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of node image to get.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: nodeImagesGroupTags(),
	}
}

// GetNodeImages defines the routing and subsequent specification for GET to /api/v0/container_images.
func GetNodeImages(input *models.NodeImage) huma.Operation {
	return huma.Operation{
		OperationID:   "getNodeImages",
		Summary:       "Get all node image information.",
		Description:   "Get all node image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultNodeImagesPath),
		Tags:          nodeImagesGroupTags(),
	}
}

// DeleteNodeImage defines the routing and subsequent specification for DELETE to /api/v0/container_images/{container_image_id}.
func DeleteNodeImage(input *models.NodeImage) huma.Operation {
	return huma.Operation{
		OperationID:   "deleteNodeImage",
		Summary:       "Delete a specific node image.",
		Description:   "Delete a specific node image.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultNodeImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of node image to delete.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: nodeImagesGroupTags(),
	}
}

// nodeImagesGroupTags defines the group tags that group similar APIs in the documentation site.
func nodeImagesGroupTags() []string {
	return []string{
		DefaultNodeImagesGroupTag,
		v0.DefaultGroupTag,
	}
}
