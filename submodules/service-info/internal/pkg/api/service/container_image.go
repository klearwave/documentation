package service

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/internal/pkg/api"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
	createv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/create"
	deletev0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/delete"
	listv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/list"
	readv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/read"
	responsev0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/response"
)

// CreateContainerImageV0 defines the service for creating container images.
func (service *Service) CreateContainerImageV0(ctx context.Context, req *createv0.ContainerImageRequest) (*responsev0.ContainerImage, error) {
	res := service.Create(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getContainerImageV0ResponseFromResult(res)
}

// GetContainerImageV0 defines the service for retrieving specific container images.
func (service *Service) GetContainerImageV0(ctx context.Context, req *readv0.ContainerImageRequest) (*responsev0.ContainerImage, error) {
	res := service.Read(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getContainerImageV0ResponseFromResult(res)
}

// ListContainerImagesV0 defines the service for listing versions.
func (service *Service) ListContainerImagesV0(ctx context.Context, req *listv0.ContainerImageRequest) (*responsev0.ContainerImage, error) {
	res := service.List(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getContainerImagesV0ResponseFromResult(res)
}

// ListContainerImageVersionsV0 defines the service for listing all versions which have a specific container image.
func (service *Service) ListContainerImageVersionsV0(_ context.Context, req *readv0.ContainerImageRequest) (*responsev0.Version, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	// run pre-request validation
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	apiResponse := &responsev0.Version{}

	err := service.Database.Connection.
		Model(req.ToReader()).
		Association("Versions").
		Find(&apiResponse.Body.Items)
	if err != nil {
		apiResponse.Status = http.StatusInternalServerError

		return nil, huma.Error500InternalServerError("unable to fetch versions from container image", err)
	}

	return apiResponse, nil
}

// DeleteContainerImageV0 defines the service for deleting container images.
func (service *Service) DeleteContainerImageV0(ctx context.Context, req *deletev0.ContainerImageRequest) (*responsev0.ContainerImage, error) {
	res := service.Delete(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getContainerImageV0ResponseFromResult(res)
}

// getContainerImageV0ResponseFromResult converts the result to a version object.
func getContainerImageV0ResponseFromResult(res *api.Result) (*responsev0.ContainerImage, error) {
	model, ok := res.Object.(*v0.ContainerImage)
	if !ok {
		return nil, huma.Error500InternalServerError("failed to convert body to proper type")
	}

	return &responsev0.ContainerImage{
		Body: responsev0.ContainerImageResponseBody{
			Items: []v0.ContainerImage{*model},
		},
		Status: http.StatusOK,
	}, nil
}

// getContainerImagesV0ResponseFromResult converts the result to a version object.
func getContainerImagesV0ResponseFromResult(res *api.Result) (*responsev0.ContainerImage, error) {
	model, ok := res.Object.(*v0.ContainerImages)
	if !ok {
		return nil, huma.Error500InternalServerError("failed to convert body to proper type")
	}

	return &responsev0.ContainerImage{
		Body: responsev0.ContainerImageResponseBody{
			Items: *model,
		},
		Status: http.StatusOK,
	}, nil
}
