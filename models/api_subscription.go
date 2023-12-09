package models

import (
	"github.com/google/uuid"
)

type ApiSubscription struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	ApiPlanID    uuid.UUID `json:"api_plan_id"`
}
