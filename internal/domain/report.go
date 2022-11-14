package domain

import (
	"github.com/google/uuid"
	"time"
)

type Report struct {
	ID        uuid.UUID
	ServiceID uuid.UUID
	Amount    float64
	CreatedAt time.Time
}
