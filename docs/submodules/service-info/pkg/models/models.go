package models

import "time"

// Model represents the base model that is used for all other models.
type Model struct {
	Id        int       `json:"id,omitempty" gorm:"primarykey" example:"1" doc:"Database ID of the stored object."`
	CreatedAt time.Time `json:"created_at,omitempty" example:"YYYY-MM-DDTHH:MM:SSZ" doc:"Object creation time in RFC 3339 format."`
}

// DeleteResponse represents the generic delete response message for all models.
type DeleteResponse struct {
	Body struct {
		Message string
	}
}
