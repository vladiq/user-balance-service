package account

import (
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"time"
)

type DTO struct {
	ID        uuid.UUID `db:"id"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
}

func (dto *DTO) Entity() *domain.Account {
	acc := &domain.Account{
		ID:        dto.ID,
		Balance:   dto.Balance,
		CreatedAt: dto.CreatedAt,
	}
	return acc
}
