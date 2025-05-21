package response

import v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"

// ContainerImage represents the response when interacting with the version API.
type ContainerImage struct {
	Body   ContainerImageResponseBody
	Status int
}

// ContainerImageResponseBody represents the body of the response for returning a single version.
type ContainerImageResponseBody struct {
	Items []v0.ContainerImage `json:"items"`
}
