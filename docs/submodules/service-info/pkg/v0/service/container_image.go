package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/klearwave/service-info/pkg/db"
	"github.com/klearwave/service-info/pkg/models"
	modelsv0 "github.com/klearwave/service-info/pkg/v0/models"
)

var (
	ErrMissingContainerImageId = errors.New("missing id parameter")
)

// CreateContainerImage defines the service for creating a new container image and storing in a database.
func (service *Service) CreateContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestCreate) (*modelsv0.ContainerImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImage := &modelsv0.ContainerImage{}
	containerImage.FromRequest(&request.Body)

	// create the container image
	if err := service.Database.Create(containerImage, "Image", "SHA256Sum", "CommitHash"); err != nil {
		return nil, err
	}

	return &modelsv0.ContainerImageResponse{
		Body: *containerImage,
	}, nil
}

// GetContainerImage defines the service for retrieving a specific container image from the database.
func (service *Service) GetContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestGet) (*modelsv0.ContainerImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	response := &modelsv0.ContainerImageResponse{Body: modelsv0.ContainerImage{}}

	if _, err := service.Database.FindBy("id", request.Id, &response.Body); err != nil {
		return nil, err
	}

	if response.Body.Id == db.MissingDatabaseID {
		response.Status = http.StatusNotFound

		return response, nil
	}

	return response, nil
}

// GetContainerImages defines the service for retrieving all container images from a database.
func (service *Service) GetContainerImages(ctx context.Context, request *modelsv0.ContainerImageRequestList) (*modelsv0.ContainerImagesResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImages := []modelsv0.ContainerImage{}

	result := service.Database.Connection.Find(&containerImages)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.ContainerImagesResponse{
		Body: containerImages,
	}, nil
}

// DeleteContainerImage defines the service for deleting a container image from a database.
func (service *Service) DeleteContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestDelete) (*models.DeleteResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImage := modelsv0.ContainerImage{}

	if request.Id == db.MissingDatabaseID {
		return nil, ErrMissingContainerImageId
	}

	if _, err := service.Database.FindBy("id", request.Id, &containerImage); err != nil {
		return nil, err
	}

	if err := service.Database.Delete(containerImage.Id, containerImage); err != nil {
		return nil, err
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}

// GetContainerImageVersions defines the service for deleting a container image from a database.
func (service *Service) GetContainerImageVersions(ctx context.Context, request *modelsv0.ContainerImageRequestGet) (*modelsv0.VersionsResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	var versions []modelsv0.Version

	response := &modelsv0.VersionsResponse{}

	err := service.Database.Connection.Model(&modelsv0.Version{}).Preload("ContainerImages").Find(&versions).Error

	if err != nil {
		response.Status = http.StatusInternalServerError

		return response, nil
	}

	return response, nil
}
