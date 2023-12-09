package models

import (
	"github.com/google/uuid"
)

type ApiPlan struct {
	ID           uuid.UUID `json:"id"`
	ApiID        uuid.UUID `json:"api_id"`
	Name         string    `json:"name"`
	ValueCents   int64     `json:"value_cents"`
	RequestLimit int64     `json:"request_limit"`
}
