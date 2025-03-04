package models

import (
	"github.com/klearwave/service-info/pkg/models"
)

// ContainerImageBase is the base set of fields for all ContainerImage objects.
type ContainerImageBase struct {
	Image      string `json:"image,omitempty" example:"ghcr.io/klearwave/service-info:latest" doc:"Full container image including the registry, repository and tag."`
	SHA256Sum  string `json:"sha256sum,omitempty" example:"2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d" doc:"SHA256 sum of the container image."`
	CommitHash string `json:"commit_hash,omitempty" example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" doc:"Commit hash related to the image."`
}

// ContainerImageRequestBody is the base set of inputs required to create a ContainerImage.
type ContainerImageRequestBody struct {
	ContainerImageBase
}

// TableName defines the table name for the request.
func (body ContainerImageRequestBody) TableName() string {
	return "container_images"
}

// ContainerImage represents the full database schema for a ContainerImage.  The full schema is also used in responses.
type ContainerImage struct {
	models.Model

	ContainerImageBase

	ImagePath string `json:"image_path,omitempty" example:"ghcr.io/klearwave/service-info" doc:"Container image path including the registry and repository without the tag."`
	ImageTag  string `json:"image_tag,omitempty" example:"v0.1.2" doc:"Container image tag."`
}

// FromRequest converts a request object to a response object.
func (containerImage *ContainerImage) FromRequest(request *ContainerImageRequestBody) {
	containerImage.ContainerImageBase = request.ContainerImageBase
}

// ContainerImageRequestCreate represents the request when creating a version.
type ContainerImageRequestCreate struct {
	Body ContainerImageRequestBody
}

// ContainerImageRequestGet represents the request when getting a version.
type ContainerImageRequestGet struct {
	Id int `path:"id"`
}

// ContainerImageRequestDelete represents the request when deleting a version.
type ContainerImageRequestDelete struct {
	Id int `path:"id"`
}

// ContainerImageRequestList represents the request when listing all versions.
type ContainerImageRequestList struct{}

// ContainerImageResponse is the response returned that includes the ContainerImage struct as the response body.
type ContainerImageResponse struct {
	Body   ContainerImage
	Status int
}

// ContainerImagesResponse is the response returned when multiple versions are requested.
type ContainerImagesResponse struct {
	Body   []ContainerImage
	Status int
}
