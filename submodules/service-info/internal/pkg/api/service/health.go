package service

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/request/read"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/response"
)

// HealthZ defines the service for ensuring the service is healthy.
func (service *Service) HealthZ(_ context.Context, _ *read.Health) (*response.Health, error) {
	health := &response.Health{
		Body: unversioned.Health{},
	}

	db, err := service.Database.Connection.DB()
	if err != nil {
		health.Status = http.StatusInternalServerError
		health.Body.Status = unversioned.HealthStatusUnhealthy

		return nil, huma.Error500InternalServerError("failed to get database connection")
	}

	if err = db.Ping(); err != nil {
		health.Status = http.StatusInternalServerError
		health.Body.Status = unversioned.HealthStatusUnhealthy

		return nil, huma.Error500InternalServerError("failed to ping database")
	}

	health.Status = http.StatusOK
	health.Body.Status = unversioned.HealthStatusHealthy

	return health, nil
}
