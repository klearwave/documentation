package models

import (
	"github.com/klearwave/service-info/pkg/models"
)

// NodeImageInput is the set of inputs required to create a node image.
type NodeImageInput struct {
	Image      string `json:"image,omitempty" example:"ghcr.io/klearwave/service-info" doc:"Container image including the registry and repository without the tag."`
	ImageTag   string `json:"image_tag,omitempty" example:"v0.1.2" doc:"Container image tag."`
	SHA256Sum  string `json:"sha256sum,omitempty" example:"2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d" doc:"SHA256 sum of the container image."`
	CommitHash string `json:"commit_hash,omitempty" example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" doc:"Commit hash related to the image."`

	// TODO
	// Versions []string `json:"versions,omitempty" example:"v0.1.2" doc:"Container image tag."`
}

// TableName sets the database table name.
func (input NodeImageInput) TableName() string {
	return "node_images"
}

// NodeImage represents the full database schema for a NodeImage.
type NodeImage struct {
	models.Model

	NodeImageInput
}

// NodeImageRequestCreate represents the request when creating a version.
type NodeImageRequestCreate struct {
	Body NodeImageInput
}

// NodeImageRequestGet represents the request when getting a version.
type NodeImageRequestGet struct {
	Id int `path:"id"`
}

// NodeImageRequestDelete represents the request when deleting a version.
type NodeImageRequestDelete struct {
	Id int `path:"id"`
}

// NodeImageRequestList represents the request when listing all versions.
type NodeImageRequestList struct{}

// NodeImageResponse is the response returned that includes the NodeImage struct as the response body.
type NodeImageResponse struct {
	Body   NodeImage
	Status int
}

// NodeImagesResponse is the response returned when multiple versions are requested.
type NodeImagesResponse struct {
	Body   []NodeImage
	Status int
}
