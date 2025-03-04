package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"github.com/klearwave/service-info/pkg/db"
	"github.com/klearwave/service-info/pkg/v0/models"
	"github.com/klearwave/service-info/pkg/v0/routes"
	"github.com/klearwave/service-info/pkg/v0/service"
	servicev0 "github.com/klearwave/service-info/pkg/v0/service"
)

var version string = "unstable"

// server is a struct that represents a running server instance and parameters
// required to run a server instance.
type server struct {
	Router   *gin.Engine
	API      huma.API
	Database *db.Database

	// services
	ServiceV0 *servicev0.Service
}

// NewServer returns a new instance of a server.  It returns an error if the
// server cannot be created.
func NewServer() (*server, error) {
	s := &server{Router: gin.Default()}

	config := huma.DefaultConfig("Information API", version)
	config.DocsPath = "/api/v0/docs"

	s.API = humagin.New(s.Router, config)

	db, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	s.Database = db
	s.ServiceV0 = service.NewService(db)

	return s, nil
}

// RegisterRoutes registers all routes associated with a server instance.
func (s *server) RegisterRoutes() {
	// version routes v0
	version := models.Version{}
	huma.Register(s.API, routes.CreateVersion(&version), s.ServiceV0.CreateVersion)
	huma.Register(s.API, routes.GetVersion(&version), s.ServiceV0.GetVersion)
	huma.Register(s.API, routes.GetVersions(&version), s.ServiceV0.GetVersions)
	huma.Register(s.API, routes.DeleteVersion(&version), s.ServiceV0.DeleteVersion)

	// container image routes v0
	containerImage := models.ContainerImage{}
	huma.Register(s.API, routes.CreateContainerImage(&containerImage), s.ServiceV0.CreateContainerImage)
	huma.Register(s.API, routes.GetContainerImage(&containerImage), s.ServiceV0.GetContainerImage)
	huma.Register(s.API, routes.GetContainerImages(&containerImage), s.ServiceV0.GetContainerImages)
	huma.Register(s.API, routes.DeleteContainerImage(&containerImage), s.ServiceV0.DeleteContainerImage)
	huma.Register(s.API, routes.GetContainerImageVersions(&containerImage), s.ServiceV0.GetContainerImageVersions)
}
