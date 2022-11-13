package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Reservations struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	ServiceID uuid.UUID `db:"service_id"`
	OrderID   uuid.UUID `db:"order_id"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}

func (dto *Reservations) Entity() *domain.Reservation {
	r := &domain.Reservation{
		ID:        dto.ID,
		AccountID: dto.AccountID,
		ServiceID: dto.ServiceID,
		OrderID:   dto.OrderID,
		Amount:    dto.Amount,
		CreatedAt: dto.CreatedAt,
	}
	return r
}
