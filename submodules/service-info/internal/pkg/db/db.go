package db

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultWaitTimeoutSeconds         = 30
	defaultWaitTimeoutIntervalSeconds = 2

	MissingDatabaseID int = 0
)

// Database represents a connection to a backend database.
type Database struct {
	Lock       *sync.Mutex
	Config     *Config
	Connection *gorm.DB
}

// NewDatabase creates a new instance of a database object.
func NewDatabase(config *Config) (*Database, error) {
	if err := config.Parse(); err != nil {
		return nil, fmt.Errorf("unable to parse database connection; %w", err)
	}

	return &Database{
		Lock:   &sync.Mutex{},
		Config: config,
	}, nil
}

// Open opens the connection to the database.
func (database *Database) Open() error {
	db, err := gorm.Open(postgres.Open(database.Config.String))
	if err != nil {
		return err
	}

	database.Connection = db

	// wait up to 30 seconds for the connection to open
	err = database.Wait(defaultWaitTimeoutSeconds)
	if err != nil {
		return err
	}

	return err
}

// Close closes the connection to the database.
func (database *Database) Close() error {
	connection, err := database.Connection.DB()
	if err != nil {
		return err
	}

	return connection.Close()
}

// Read is a generic function to read a model
// when no special logic is required.
func (database *Database) Read(id int, model any) error {
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
func (database *Database) FindBy(field string, value, model any) (*gorm.DB, error) {
	result := database.Connection.Where(map[string]any{field: value}).First(model)
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
func (database *Database) Delete(id int, model any) error {
	result := database.Connection.Delete(model, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Wait waits for a database connection to be established.
func (database *Database) Wait(timeoutSeconds int64) error {
	db, err := database.Connection.DB()
	if err != nil {
		return fmt.Errorf("unable to create sql database connection; %w", err)
	}

	ticker := time.NewTicker(defaultWaitTimeoutIntervalSeconds * time.Second)
	defer ticker.Stop()

	timeoutChan := time.After(time.Duration(timeoutSeconds) * time.Second)

	for {
		select {
		case <-ticker.C:
			if err := db.Ping(); err == nil {
				return nil
			} else {
				continue
			}
		case <-timeoutChan:
			return fmt.Errorf("database connection timeout after %d seconds", timeoutSeconds)
		}
	}
}
