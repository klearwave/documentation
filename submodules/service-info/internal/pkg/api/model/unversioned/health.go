package unversioned

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/db"
)

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
)

type Health struct {
	Status HealthStatus
}

// Read handles the read request for an about model.
func (health *Health) Read(_ *db.Database) *api.Result {
	return &api.Result{
		Object: health,
		Error:  nil,
	}
}
