package domain

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
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
