package unversioned

import (
	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/db"
)

type About struct {
	Version    string `doc:"Running version of this service in semantic versioning format." example:"v0.1.2"                                   json:"version,omitempty"`
	CommitHash string `doc:"Commit hash of this running version."                           example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" json:"commit_hash,omitempty"`
}

// Read handles the read request for an about model.
func (about *About) Read(_ *db.Database) *api.Result {
	return &api.Result{
		Object: about,
		Error:  nil,
	}
}
