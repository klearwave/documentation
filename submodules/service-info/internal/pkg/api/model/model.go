package model

import (
	"time"

	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
)

// Model represents the base model that is used for all other models.
type Model struct {
	CreatedAt *time.Time `doc:"Object creation time in RFC 3339 format."     example:"YYYY-MM-DDTHH:MM:SSZ" json:"created_at,omitempty"`
	UpdatedAt *time.Time `doc:"Object last updated time in RFC 3339 format." example:"YYYY-MM-DDTHH:MM:SSZ" json:"updated_at,omitempty"`
}

// WithID represents the base model that is used for all other models when an ID is used as
// a primary key.  If this model is not used, it is assumed that the underlying model defines
// its own primary key.
type WithID struct {
	Model

	ID int `doc:"Database ID of the stored object." example:"1" gorm:"primarykey" json:"id,omitempty"`
}

// StringFetcher represents a model that uses a string with an 'id' parameter to fetch a model
// from the database.  It is used to inject into other models so that validation may be reused
// across multiple models.
type StringFetcher struct {
	ID string `doc:"Database ID of the stored object." path:"id"`
}

// IntegerFetcher represents a model that uses an integer with an 'id' parameter to fetch a model
// from the database.  It is used to inject into other models so that validation may be reused
// across multiple models.
type IntegerFetcher struct {
	ID int `doc:"Database ID of the stored object." example:"1" path:"id"`
}

// StringFetcher.IsValid validates that a string fetcher is valid.
func (fetcher StringFetcher) IsValid() (bool, error) {
	if fetcher.ID == "" {
		return false, apierrors.ErrMissingParameterID
	}

	return true, nil
}

// IntegerFetcher.IsValid validates that a integer fetcher is valid.  In this case, we will never
// have a database ID of less than 1 (the first entry always has a positive number).
func (fetcher IntegerFetcher) IsValid() (bool, error) {
	if fetcher.ID < 1 {
		return false, apierrors.ErrMissingParameterID
	}

	return true, nil
}
