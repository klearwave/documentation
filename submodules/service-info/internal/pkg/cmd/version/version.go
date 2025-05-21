package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/api"
)

const versionExample = `
service version
`

func NewCommand() *cobra.Command {
	// create the command
	command := &cobra.Command{
		Use:     "version",
		Short:   "Print version information",
		Long:    `Print version information`,
		Run:     func(_ *cobra.Command, _ []string) { run() },
		Example: versionExample,
	}

	return command
}

//nolint:forbidigo // this is just printing out information, so we have no issues here
func run() {
	fmt.Printf("%s\n", api.ServerVersion)
}
