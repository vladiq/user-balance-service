package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	IsAccrual bool
	Amount    float64
	Info      string
	CreatedAt time.Time
}
