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

// CreateVersionV0 defines the service for retrieving specific version information.
func (service *Service) CreateVersionV0(ctx context.Context, req *createv0.VersionRequest) (*responsev0.Version, error) {
	res := service.Create(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getVersionV0ResponseFromResult(res)
}

// GetVersionV0 defines the service for retrieving specific version information.
func (service *Service) GetVersionV0(ctx context.Context, req *readv0.VersionRequest) (*responsev0.Version, error) {
	res := service.Read(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getVersionV0ResponseFromResult(res)
}

// ListVersionsV0 defines the service for listing versions.
func (service *Service) ListVersionsV0(ctx context.Context, req *listv0.VersionRequest) (*responsev0.Version, error) {
	res := service.List(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getVersionsV0ResponseFromResult(res)
}

// ListVersionContainerImagesV0 defines the service for listing all container images belonging to a specific version.
func (service *Service) ListVersionContainerImagesV0(_ context.Context, req *readv0.VersionRequest) (*responsev0.ContainerImage, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	// run pre-request validation
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	apiResponse := &responsev0.ContainerImage{}

	err := service.Database.Connection.
		Model(req.ToReader()).
		Association("ContainerImages").
		Find(&apiResponse.Body.Items)
	if err != nil {
		apiResponse.Status = http.StatusInternalServerError

		return nil, huma.Error500InternalServerError("unable to fetch container images from version", err)
	}

	return apiResponse, nil
}

// DeleteVersionV0 defines the service for deleting a specific version.
func (service *Service) DeleteVersionV0(ctx context.Context, req *deletev0.VersionRequest) (*responsev0.Version, error) {
	res := service.Delete(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	return getVersionV0ResponseFromResult(res)
}

// getVersionV0ResponseFromResult converts the result to a version object.
func getVersionV0ResponseFromResult(res *api.Result) (*responsev0.Version, error) {
	model, ok := res.Object.(*v0.Version)
	if !ok {
		return nil, huma.Error500InternalServerError("failed to convert body to proper type")
	}

	return &responsev0.Version{
		Body: responsev0.VersionResponseBody{
			Items: []v0.Version{*model},
		},
		Status: http.StatusOK,
	}, nil
}

// getVersionsV0ResponseFromResult converts the result to a version object.
func getVersionsV0ResponseFromResult(res *api.Result) (*responsev0.Version, error) {
	model, ok := res.Object.(*v0.Versions)
	if !ok {
		return nil, huma.Error500InternalServerError("failed to convert body to proper type")
	}

	return &responsev0.Version{
		Body: responsev0.VersionResponseBody{
			Items: *model,
		},
		Status: http.StatusOK,
	}, nil
}
