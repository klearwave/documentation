package db

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MissingDatabaseID int = 0
)

var (
	ErrMissingDatabaseHost     = errors.New("missing DB_HOST environment variable")
	ErrMissingDatabaseName     = errors.New("missing DB_NAME environment variable")
	ErrMissingDatabaseUsername = errors.New("missing DB_USERNAME environment variable")
	ErrMissingDatabasePassword = errors.New("missing DB_PASSWORD environment variable")
)

type Database struct {
	Lock       *sync.Mutex
	Connection *gorm.DB
}

func NewDatabase() (*Database, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, ErrMissingDatabaseHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, ErrMissingDatabaseName
	}

	dbUser := os.Getenv("DB_USERNAME")
	if dbUser == "" {
		return nil, ErrMissingDatabaseUsername
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, ErrMissingDatabasePassword
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword,
	)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}

	return &Database{
		Lock:       &sync.Mutex{},
		Connection: db,
	}, nil
}

// Create is a generic function to create a model
// when no special logic is required.  If fields is specified,
// the the model will only be created with the select fields.
func (database *Database) Create(model interface{}, fields ...string) error {
	if len(fields) == 0 {
		result := database.Connection.Create(model)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}

	result := database.Connection.Select(fields).Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CreateWithFields is a generic function to create
// a model given only a set of specific fields.
func (database *Database) CreateWithFields(model interface{}, fields ...string) error {
	result := database.Connection.Select(fields).Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Read is a generic function to read a model
// when no special logic is required.
func (database *Database) Read(id int, model interface{}) error {
	result := database.Connection.Find(model, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("unable to find id [%d]", id)
	}

	return nil
}

// FindBy is a generic function to read a model by
// which matches a field with a particular value.
func (database *Database) FindBy(field string, value interface{}, model interface{}) (*gorm.DB, error) {
	result := database.Connection.Where(map[string]interface{}{field: value}).First(model)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return result, nil
		}

		return result, result.Error
	}

	return result, nil
}

// Delete is a generic function to delete a model
// when no special logic is required.
func (database *Database) Delete(id int, model interface{}) error {
	result := database.Connection.Delete(model, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
