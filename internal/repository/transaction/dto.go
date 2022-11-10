package transaction

import (
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"time"
)

type DTO struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	IsAccrual bool      `db:"is_accrual"`
	Amount    float64   `db:"amount"`
	Info      string    `db:"info"`
	CreatedAt time.Time `db:"created_at"`
}

func (dto *DTO) Entity() *domain.Transaction {
	t := &domain.Transaction{
		ID:        dto.ID,
		AccountID: dto.AccountID,
		IsAccrual: dto.IsAccrual,
		Amount:    dto.Amount,
		Info:      dto.Info,
		CreatedAt: dto.CreatedAt,
	}
	return t
}
