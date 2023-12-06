package models

import (
	"github.com/google/uuid"
)

type Api struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
}
