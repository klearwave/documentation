package models

import "github.com/klearwave/service-info/pkg/models"

// VersionBase is the base set of fields for all Version objects.
type VersionBase struct {
	VersionId string `json:"version_id,omitempty" example:"v0.1.2" doc:"Version in semantic versioning format."`
	Latest    bool   `json:"latest,omitempty" example:"true" doc:"Whether or not this is the latest version."`
}

// VersionRequestBody is the base set of inputs required to create a Version.
type VersionRequestBody struct {
	VersionBase

	ContainerImages []*ContainerImageRequestBody `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// TableName defines the table name for the request.
func (body VersionRequestBody) TableName() string {
	return "versions"
}

// Version represents the full database schema for a Version.  The full schema is also used in responses.
type Version struct {
	models.Model
	VersionBase

	XVersion int `json:"x_version,omitempty" example:"0" doc:"The X version of this release (e.g. 0; v0.1.2 == x.y.z)."`
	YVersion int `json:"y_version,omitempty" example:"1" doc:"The Y version of this release (e.g. 1; v0.1.2 == x.y.z)."`
	ZVersion int `json:"z_version,omitempty" example:"2" doc:"The Z version of this release (e.g. 2; v0.1.2 == x.y.z)."`

	ContainerImages []*ContainerImage `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// FromRequest converts a request object to a response object.
func (version *Version) FromRequest(request *VersionRequestBody) {
	version.VersionBase = request.VersionBase

	if len(request.ContainerImages) < 1 {
		return
	}

	version.ContainerImages = make([]*ContainerImage, len(request.ContainerImages))

	for i := range request.ContainerImages {
		version.ContainerImages[i] = &ContainerImage{}
		version.ContainerImages[i].FromRequest(request.ContainerImages[i])
	}
}

// VersionRequestCreate represents the request when creating a version.
type VersionRequestCreate struct {
	Body VersionRequestBody
}

// VersionRequestGet represents the request when getting a version.
type VersionRequestGet struct {
	VersionID string `path:"version_id"`
}

// VersionRequestDelete represents the request when deleting a version.
type VersionRequestDelete struct {
	VersionID string `path:"version_id"`
}

// VersionRequestList represents the request when listing all versions.
type VersionRequestList struct{}

// VersionResponse is the response returned that includes the Version struct as the response body.
type VersionResponse struct {
	Body   Version
	Status int
}

// VersionsResponse is the response returned when multiple versions are requested.
type VersionsResponse struct {
	Body   []Version
	Status int
}
