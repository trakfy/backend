package models

import (
	"time"

	"github.com/google/uuid"
)

type ApiKey struct {
	ID          uuid.UUID `json:"id"`
	ApiID       uuid.UUID `json:"api_id"`
	UserID      uuid.UUID `json:"user_id"`
	Key         string    `json:"key"`
	Valid       bool      `json:"validity"`
	QuotaUsed   int       `json:"quota_used"`
	RenewalDate time.Time `json:"renewal_date"`
}
