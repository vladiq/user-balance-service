package models

import (
	"time"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
)

type Accounts struct {
	ID        uuid.UUID `db:"id"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
}

func (a *Accounts) Entity() *domain.Account {
	acc := &domain.Account{
		ID:        a.ID,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
	}
	return acc
}
