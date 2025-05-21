package service

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/request/read"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/response"
)

// GetAbout defines the service for retrieving about information for the service.
func (service *Service) GetAbout(ctx context.Context, req *read.About) (*response.About, error) {
	res := service.Read(ctx, req)
	if res.Error != nil {
		return nil, res.Error
	}

	model, ok := res.Object.(*unversioned.About)
	if !ok {
		return nil, huma.Error500InternalServerError("failed to convert body to proper type")
	}

	return &response.About{
		Body:   *model,
		Status: http.StatusOK,
	}, nil
}
