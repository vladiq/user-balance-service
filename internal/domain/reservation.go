package domain

import (
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    float64
	CreatedAt time.Time
	ClosedAt  time.Time
}
