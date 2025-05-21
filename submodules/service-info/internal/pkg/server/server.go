package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/route"
	routev0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/route"
	"github.com/klearwave/service-info/internal/pkg/api/service"
	"github.com/klearwave/service-info/internal/pkg/db"
)

const (
	HTTPPort  = 8888
	HTTPSPort = 8443

	tlsCertPath = "/tls.crt"
	tlsKeyPath  = "/tls.key"

	DefaultReadHeaderTimeoutSeconds = 15
	DefaultShutdownTimeoutSeconds   = 10
)

// server is a struct that represents a running server instance and parameters
// required to run a server instance.
type server struct {
	Router *gin.Engine
	API    huma.API
	Server *http.Server
	TLS    bool

	// services
	Service *service.Service
}

// NewServer returns a new instance of a server.  It returns an error if the
// server cannot be created.
func NewServer() (*server, error) {
	s := &server{
		Router:  gin.Default(),
		Service: service.NewService(),
	}

	config := huma.DefaultConfig("Information API", api.ServerVersion)
	config.DocsPath = "/api/v0/docs"

	s.API = humagin.New(s.Router, config)

	// determine if we are using TLS and set appropriately
	port := HTTPPort

	if certExists(tlsCertPath) && certExists(tlsKeyPath) {
		s.TLS = true
		port = HTTPSPort
	}

	// create the server
	s.Server = &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		Handler:           s.Router,
		ReadHeaderTimeout: DefaultReadHeaderTimeoutSeconds * time.Second,
	}

	return s, nil
}

// Init initializes the server.
func (s *server) Init(config *db.Config) error {
	// create and validate the database config
	database, err := db.NewDatabase(config)
	if err != nil {
		return err
	}

	// start the web service
	return s.Service.Start(database)
}

// Start starts the server.
func (s *server) Start() error {
	// ensure the server is initialized
	if s.Service.Database == nil {
		return errors.New("service not initialized; missing database configuration")
	}

	var err error

	go func() {
		if s.TLS {
			log.Printf("Starting HTTPS server on %s", s.Server.Addr)
			err = s.Server.ListenAndServeTLS(tlsCertPath, tlsKeyPath)
		} else {
			log.Printf("Starting HTTP server on %s", s.Server.Addr)
			err = s.Server.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			err = fmt.Errorf("server error: %w", err)
		}
	}()

	return nil
}

// Stop stops the server.
func (s *server) Stop(ctx context.Context) error {
	log.Println("Shutting down server...")

	if err := s.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown failed: %w", err)
	}

	return s.Service.Stop()
}

// RegisterRoutes registers all routes associated with a server instance.
func (s *server) RegisterRoutes() {
	// unversioned routes
	huma.Register(s.API, route.HealthZ(), s.Service.HealthZ)
	huma.Register(s.API, route.GetAbout(), s.Service.GetAbout)

	// version routes v0
	huma.Register(s.API, routev0.CreateVersion(), s.Service.CreateVersionV0)
	huma.Register(s.API, routev0.GetVersion(), s.Service.GetVersionV0)
	huma.Register(s.API, routev0.ListVersions(), s.Service.ListVersionsV0)
	huma.Register(s.API, routev0.ListVersionContainerImages(), s.Service.ListVersionContainerImagesV0)
	huma.Register(s.API, routev0.DeleteVersion(), s.Service.DeleteVersionV0)

	// container image routes v0
	huma.Register(s.API, routev0.CreateContainerImage(), s.Service.CreateContainerImageV0)
	huma.Register(s.API, routev0.GetContainerImage(), s.Service.GetContainerImageV0)
	huma.Register(s.API, routev0.ListContainerImages(), s.Service.ListContainerImagesV0)
	huma.Register(s.API, routev0.ListContainerImageVersions(), s.Service.ListContainerImageVersionsV0)
	huma.Register(s.API, routev0.DeleteContainerImage(), s.Service.DeleteContainerImageV0)
}

// certExists determines if a certificate exists at a given path.
func certExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
