package health

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/request/read"
	"github.com/klearwave/service-info/internal/pkg/api/model/unversioned/route"
	"github.com/klearwave/service-info/internal/pkg/server"
)

const healthExample = `
service health
`

type input struct {
	Host     string
	Port     int
	TLS      bool
	Insecure bool
}

func NewCommand() *cobra.Command {
	commandInput := &input{}

	// create the command
	command := &cobra.Command{
		Use:     "health",
		Short:   "Check if the server is healthy",
		Long:    `Check if the server is healthy`,
		RunE:    func(_ *cobra.Command, _ []string) error { return health(commandInput) },
		Example: healthExample,
	}

	command.Flags().StringVar(&commandInput.Host, "host", "127.0.0.1", "Host address where the server is running")
	command.Flags().IntVar(&commandInput.Port, "port", server.HTTPSPort, "Host port where the server is running")
	command.Flags().BoolVar(&commandInput.TLS, "tls", true, "Whether or not the server is using TLS or not")
	command.Flags().BoolVar(&commandInput.Insecure, "insecure", false, "Skip certificate verification and accept insecure health check")

	return command
}

func health(commandInput *input) error {
	prefix := "http"
	if commandInput.TLS {
		prefix = "https"
	}

	path := fmt.Sprintf(
		"%s://%s:%d/%s",
		prefix,
		commandInput.Host,
		commandInput.Port,
		route.DefaultHealthPath,
	)

	// create the request
	body, err := json.Marshal(&read.Health{})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	var client *http.Client

	//nolint:gosec // need ability to call with insecure for development purposes
	if commandInput.Insecure && commandInput.TLS {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: commandInput.Insecure},
			},
		}
	} else {
		client = &http.Client{}
	}

	// execute the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("service is unhealthy; received code [%d] and error reading body: %v", resp.StatusCode, err)
		}

		return fmt.Errorf("service is unhealthy; received code [%d] and body [%s]", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
