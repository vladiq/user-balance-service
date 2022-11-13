package domain

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    float64
	CreatedAt time.Time
}
