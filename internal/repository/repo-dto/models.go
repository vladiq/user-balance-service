package repo_dto

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID `db:"id"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
}

type Reservation struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	ServiceID uuid.UUID `db:"service_id"`
	OrderID   uuid.UUID `db:"order_id"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	ClosedAt  time.Time `db:"closed_at"`
}

type Transaction struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	IsAccrual bool      `db:"is_accrual"`
	Amount    float64   `db:"amount"`
	Info      string    `db:"info"`
	CreatedAt time.Time `db:"created_at"`
}
