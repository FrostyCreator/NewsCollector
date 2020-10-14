package model

import (
	"github.com/google/uuid"
)

type News struct {
	ID		uuid.UUID 	`json:"id" validate:"required"`
	URL 	string		`json:"url" validate:"required"`
}
