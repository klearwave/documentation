package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/cmd/generate"
	"github.com/klearwave/service-info/internal/pkg/cmd/health"
	"github.com/klearwave/service-info/internal/pkg/cmd/migrate"
	"github.com/klearwave/service-info/internal/pkg/cmd/run"
	"github.com/klearwave/service-info/internal/pkg/cmd/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(command *cobra.Command) error {
	return command.Execute()
}

// main executes the main program loop.
func main() {
	root := &cobra.Command{
		Use:   "service",
		Short: "Manage service",
		Long:  `Manage service`,
	}

	// add version subcommand for printing version information
	root.AddCommand(version.NewCommand())

	// add other subcommands
	root.AddCommand(generate.NewCommand())
	root.AddCommand(migrate.NewCommand())
	root.AddCommand(run.NewCommand())
	root.AddCommand(health.NewCommand())

	if err := Execute(root); err != nil {
		os.Exit(1)
	}
}
