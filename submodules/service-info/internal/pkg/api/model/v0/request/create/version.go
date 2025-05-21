package create

import (
	"fmt"
	"regexp"

	"github.com/klearwave/service-info/internal/pkg/api"
	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
)

const (
	versionRegex = `^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.-]+))?(?:\+([0-9A-Za-z.-]+))?$`
)

// VersionBody is the base set of inputs required to create a new Version.
type VersionBody struct {
	v0.VersionBase

	ContainerImages []*ContainerImageBody `doc:"Container images associated with this version." gorm:"many2many:version_container_images;" json:"container_images,omitempty"`
}

// VersionRequest represents the request when creating a new version.
type VersionRequest struct {
	api.Authorization
	Body VersionBody
}

// IsAuthorized checks if the user is authorized to create a version.  It is used to satisfy the
// Request interface.
func (req *VersionRequest) IsAuthorized() (bool, error) {
	return req.Authorized()
}

// IsValid checks if the request is valid for a version.  It is used to satisfy the
// Request interface.
func (req *VersionRequest) IsValid() (bool, error) {
	if req == nil {
		return false, apierrors.ErrMissingVersionObject
	}

	if req.Body.ID == nil {
		return false, apierrors.ErrMissingVersionParameterVersionID
	}

	re, err := regexp.Compile(versionRegex)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern [%s]: %w", versionRegex, err)
	}

	if !re.MatchString(*req.Body.ID) {
		return false, fmt.Errorf("input [%s] does not match the required pattern [%s]",
			*req.Body.ID,
			versionRegex,
		)
	}

	return true, nil
}

// ToCreator converts the request to a creator object.  It is used to satisfy the
// Creator interface.
func (req *VersionRequest) ToCreator() api.Creator {
	version := &v0.Version{
		VersionBase: v0.VersionBase{
			ID: req.Body.ID,
		},
	}

	if len(req.Body.ContainerImages) == 0 {
		return version
	}

	// add the container images if we have them
	version.ContainerImages = make([]*v0.ContainerImage, len(req.Body.ContainerImages))

	for i := range req.Body.ContainerImages {
		version.ContainerImages[i] = req.Body.ContainerImages[i].ToObject()
	}

	return version
}
