package service

import "github.com/klearwave/service-info/pkg/db"

type Service struct {
	Database *db.Database
}

func NewService(db *db.Database) *Service {
	return &Service{
		Database: db,
	}
}
