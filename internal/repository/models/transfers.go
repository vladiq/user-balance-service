package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Transfers struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	IsAccrual bool      `db:"is_accrual"`
	Amount    float64   `db:"amount"`
	Info      string    `db:"info"`
	CreatedAt time.Time `db:"created_at"`
}

func (dto *Transfers) Entity() *domain.Transfer {
	t := &domain.Transfer{
		ID:        dto.ID,
		AccountID: dto.AccountID,
		IsAccrual: dto.IsAccrual,
		Amount:    dto.Amount,
		Info:      dto.Info,
		CreatedAt: dto.CreatedAt,
	}
	return t
}
