package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/db"
	"github.com/klearwave/service-info/internal/pkg/server"
)

const runExample = `
service run
`

func NewCommand() *cobra.Command {
	// create the command
	command := &cobra.Command{
		Use:     "run",
		Short:   "Run the service",
		Long:    `Run the service`,
		RunE:    run,
		Example: runExample,
	}

	return command
}

func run(_ *cobra.Command, _ []string) error {
	// create the srv
	srv, err := server.NewServer()
	if err != nil {
		return fmt.Errorf("failed setting up server: %w", err)
	}

	srv.RegisterRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// initialize the server
	if err := srv.Init(&db.Config{}); err != nil {
		return err
	}

	// start the server and block waiting for cancel
	if err := srv.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	// gracefully stop the server
	shutdownCtx, cancel := context.WithTimeout(ctx, server.DefaultShutdownTimeoutSeconds*time.Second)
	defer cancel()

	return srv.Stop(shutdownCtx)
}
