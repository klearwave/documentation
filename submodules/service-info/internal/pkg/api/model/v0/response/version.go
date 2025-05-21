package response

import v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"

// Version represents the response when interacting with the version API.
type Version struct {
	Body   VersionResponseBody
	Status int
}

// VersionResponseBody represents the body of the response for returning a single version.
type VersionResponseBody struct {
	Items []v0.Version `json:"items"`
}
