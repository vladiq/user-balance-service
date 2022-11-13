package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	IsAccrual bool
	Amount    float64
	Info      string
	CreatedAt time.Time
}
