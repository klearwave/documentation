package api

import (
	"fmt"

	"github.com/klearwave/service-info/internal/pkg/db"
)

var (
	ServerVersion = "unstable"
	CommitHash    = "unknown"
)

const (
	DefaultPath = "/api"
)

type Reader interface {
	Read(database *db.Database) *Result
}

type Lister interface {
	List(database *db.Database) *Result
}

type Creator interface {
	Create(database *db.Database) *Result
}

type Updater interface {
	Update(database *db.Database) *Result
}

type Deleter interface {
	Delete(database *db.Database) *Result
}

func PathFor(path string) string {
	return fmt.Sprintf("%s/%s", DefaultPath, path)
}
