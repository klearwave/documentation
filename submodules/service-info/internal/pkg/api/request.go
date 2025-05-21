package api

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/klearwave/service-info/internal/pkg/db"
)

// Request is an interface that represents a generic request object to be used by a service.
type Request interface {
	IsValid() (bool, error)
	IsAuthorized() (bool, error)
}

// CreateRequest is an interface that represents a request object to be used by a service create operations.
type CreateRequest interface {
	Request
	ToCreator() Creator
}

// ReadRequest is an interface that represents a request object to be used by a service read operations.
type ReadRequest interface {
	Request
	ToReader() Reader
}

// UpdateRequest is an interface that represents a request object to be used by a service update operations.
type UpdateRequest interface {
	Request
	ToUpdater() Updater
}

// DeleteRequest is an interface that represents a request object to be used by a service delete operations.
type DeleteRequest interface {
	Request
	ToDeleter() Deleter
}

// ListRequest is an interface that represents a request object to be used by a service list operations.
type ListRequest interface {
	Request
	ToLister() Lister
}

// Create runs the logic to execute a Create request against a database and return the result.
func Create(database *db.Database, model any) *Result {
	apiResult := &Result{}

	defer func() {
		apiResult.Object = model
	}()

	databaseResult := database.Connection.Create(model)
	if databaseResult.Error != nil {
		apiResult.UnknownError(databaseResult.Error)

		return apiResult
	}

	return apiResult
}

// Delete runs the logic to execute a Delete request against a database and return the result.
func Delete(database *db.Database, id, model any) *Result {
	apiResult := &Result{}

	defer func() {
		apiResult.Object = model
	}()

	// dynamically retrieve the table name from the go struct using gorm helpers.  we do this
	// because we want to return the complete object back to the user so that they can see
	// what was deleted.
	statement := &gorm.Statement{DB: database.Connection}
	if err := statement.Parse(model); err != nil {
		apiResult.SetError(http.StatusInternalServerError, "unable to determine table name", err)

		return apiResult
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ? RETURNING *", statement.Schema.Table)

	databaseResult := database.Connection.Raw(query, id).Scan(model)

	// return immediately if we did not find a record or we have no errors
	if databaseResult.RowsAffected == 0 {
		apiResult.NotFoundError(nil, id, model)

		return apiResult
	}

	if databaseResult.Error == nil {
		return apiResult
	}

	// handle the error since we know it is not nil at this point
	if errors.Is(databaseResult.Error, gorm.ErrRecordNotFound) {
		apiResult.NotFoundError(databaseResult.Error, id, model)

		return apiResult
	}

	apiResult.UnknownError(databaseResult.Error)

	return apiResult
}
