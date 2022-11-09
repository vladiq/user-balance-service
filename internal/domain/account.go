package domain

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID
	Balance   float64
	CreatedAt time.Time
}
