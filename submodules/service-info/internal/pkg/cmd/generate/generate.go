package generate

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/server"
)

const generateExample = `
service generate
`

const (
	defaultFile = "./openapi.yaml"
)

type input struct {
	File string
}

func NewCommand() *cobra.Command {
	commandInput := &input{}

	// create the command
	command := &cobra.Command{
		Use:     "generate",
		Short:   "Generate OpenAPI Spec",
		Long:    `Generate OpenAPI Spec`,
		RunE:    func(_ *cobra.Command, _ []string) error { return generate(commandInput) },
		Example: generateExample,
	}

	command.Flags().StringVarP(&commandInput.File, "file", "f", defaultFile, "Path to OpenAPI file output")

	return command
}

func generate(commandInput *input) error {
	gin.SetMode(gin.ReleaseMode)

	srv, err := server.NewServer()
	if err != nil {
		return err
	}

	srv.RegisterRoutes()

	spec, err := srv.API.OpenAPI().YAML()
	if err != nil {
		return err
	}

	outFile, err := os.Create(commandInput.File)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(spec)
	if err != nil {
		return err
	}

	return nil
}
