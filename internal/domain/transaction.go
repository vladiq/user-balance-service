package domain

import (
	"github.com/google/uuid"
)

type Transaction struct {
	FromID uuid.UUID
	ToID   uuid.UUID
	Amount float64
}
