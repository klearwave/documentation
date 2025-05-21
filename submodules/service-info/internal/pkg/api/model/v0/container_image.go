package v0

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"gorm.io/gorm"

	"github.com/klearwave/service-info/internal/pkg/api"
	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
	"github.com/klearwave/service-info/internal/pkg/api/model"
	"github.com/klearwave/service-info/internal/pkg/db"
	"github.com/klearwave/service-info/internal/pkg/utils/pointers"
)

const (
	containerImageRegex = `^((?P<registry>[^/]+?)/)?(?P<image>[^:@]+(?:/[^:@]+)*)(?::(?P<tag>[^@]+))?(?:@sha256:(?P<sha>[a-zA-Z0-9]+))?$`
)

// ContainerImageBase is the base set of fields for all ContainerImage objects.
type ContainerImageBase struct {
	Image      *string `doc:"Full container image including the registry, repository and tag." example:"ghcr.io/klearwave/service-info:latest"                            json:"image,omitempty"`
	SHA256Sum  *string `doc:"SHA256 sum of the container image."                               example:"2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d" json:"sha256sum,omitempty"`
	CommitHash *string `doc:"Commit hash related to the image."                                example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d"                         json:"commit_hash,omitempty"`
}

// ContainerImage represents the full database schema for a ContainerImage.  The full schema is also used in responses.
type ContainerImage struct {
	model.WithID

	ContainerImageBase

	ImageRegistry *string `doc:"Container image registry without the image name, tag or sha256sum information." example:"ghcr.io"                gorm:"default:docker.io"    json:"image_registry,omitempty"`
	ImageName     *string `doc:"Container image name without the registry, tag or sha256sum information."       example:"klearwave/service-info" json:"image_name,omitempty"`
	ImageTag      *string `doc:"Container image tag without the registry, image name or sha256 information."    example:"v0.1.2"                 gorm:"default:latest"       json:"image_tag,omitempty"`

	Versions []*Version `doc:"Versions associated with this container image." gorm:"many2many:version_container_images;" json:"versions,omitempty"`
}

// BeforeCreate defines the before create logic for a specific container image.  The BeforeCreate
// is used in conjunction with GORM as a trigger function that is called before
// inserting a record into the database.
func (containerImage *ContainerImage) BeforeCreate(tx *gorm.DB) error {
	if err := containerImage.Parse(); err != nil {
		return fmt.Errorf("%s; %w", apierrors.ErrInvalidContainerImage, err)
	}

	existing := &ContainerImage{}

	err := tx.Where("image = ?", containerImage.Image).First(existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("unable to lookup existing container image by name; %w", err)
	}

	containerImage.ID = existing.ID

	equal, err := containerImage.EqualTo(existing)
	if err != nil {
		return err
	}

	if !equal {
		return apierrors.ErrMismatchParameter
	}

	return nil
}

// Create handles the create request for a container image model.
func (containerImage *ContainerImage) Create(database *db.Database) *api.Result {
	return api.Create(database, containerImage)
}

// Read handles the read request for a container image model.
func (containerImage *ContainerImage) Read(database *db.Database) *api.Result {
	apiResult := &api.Result{}

	defer func() {
		apiResult.Object = containerImage
	}()

	databaseResult := database.Connection.
		Where(map[string]any{"id": containerImage.ID}).
		First(containerImage)

	if databaseResult.Error != nil {
		if errors.Is(databaseResult.Error, gorm.ErrRecordNotFound) {
			apiResult.NotFoundError(nil, containerImage.ID, containerImage)

			return apiResult
		}

		apiResult.UnknownError(databaseResult.Error)

		return apiResult
	}

	if containerImage.ID < 1 {
		apiResult.NotFoundError(nil, containerImage.ID, containerImage)

		return apiResult
	}

	return apiResult
}

// ContainerImages represents a list of container images.
type ContainerImages []ContainerImage

// List handles the list request for a set of versions.
func (containerImages *ContainerImages) List(database *db.Database) *api.Result {
	apiResult := &api.Result{}

	defer func() {
		apiResult.Object = containerImages
	}()

	databaseResult := database.Connection.Find(containerImages)
	if databaseResult.Error != nil {
		apiResult.UnknownError(databaseResult.Error)

		return apiResult
	}

	return apiResult
}

// Delete handles the delete request for a container image model.
func (containerImage *ContainerImage) Delete(database *db.Database) *api.Result {
	id := strconv.Itoa(containerImage.ID)

	return api.Delete(database, id, containerImage)
}

// EqualTo compares another container image for equality and returns
// a bool if they are not equal.
func (containerImage *ContainerImage) EqualTo(compared *ContainerImage) (bool, error) {
	for field, values := range map[string][]*string{
		"sha256sum":      {containerImage.SHA256Sum, compared.SHA256Sum},
		"commit_hash":    {containerImage.CommitHash, compared.CommitHash},
		"image":          {containerImage.Image, compared.Image},
		"image_registry": {containerImage.ImageRegistry, compared.ImageRegistry},
		"image_name":     {containerImage.ImageName, compared.ImageName},
		"image_tag":      {containerImage.ImageTag, compared.ImageTag},
	} {
		if !pointers.EqualString(values[0], values[1]) {
			return false, fmt.Errorf(
				"field [%s] not equal requested [%s] != existing [%s]; %w",
				field,
				pointers.ToString(values[0]),
				pointers.ToString(values[1]),
				apierrors.ErrMismatchParameter,
			)
		}
	}

	return true, nil
}

// Parse sets specific container image metadata such as the image path and the
// image tag.
//
//nolint:cyclop // TODO: refactor
func (containerImage *ContainerImage) Parse() error {
	if containerImage == nil {
		return apierrors.ErrMissingContainerImageObject
	}

	if containerImage.Image == nil || *containerImage.Image == "" {
		return apierrors.ErrMissingContainerImageParameterImage
	}

	image := *containerImage.Image

	// regular expression for splitting container images
	re, err := regexp.Compile(containerImageRegex)
	if err != nil {
		return fmt.Errorf("invalid container image pattern [%s]: %w", containerImageRegex, err)
	}

	matches := re.FindStringSubmatch(image)
	if matches == nil || len(matches) != 6 {
		return fmt.Errorf("invalid container image [%s] for pattern [%s]", image, containerImageRegex)
	}

	// validate that we do not have a mismatch between the full image sha256 (if provided) and the
	// value from the input and set the appropriate value.
	if containerImage.SHA256Sum == nil {
		containerImage.SHA256Sum = pointers.FromString(matches[5])
	} else {
		if *containerImage.SHA256Sum == "" {
			containerImage.SHA256Sum = pointers.FromString(matches[5])
		} else if matches[5] != "" && matches[5] != *containerImage.SHA256Sum {
			return fmt.Errorf("found sha256sum value [%s] but parsed [%s]; %w",
				*containerImage.SHA256Sum,
				matches[5],
				apierrors.ErrMismatchParameter,
			)
		}
	}

	// set the other parsed values.  note that no comparison between what is set and
	// what is parsed needs to happen as we do not provide these values via the request.
	containerImage.ImageRegistry = pointers.FromString(matches[2])
	containerImage.ImageName = pointers.FromString(matches[3])
	containerImage.ImageTag = pointers.FromString(matches[4])

	return nil
}
