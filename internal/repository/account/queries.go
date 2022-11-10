package account

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Queries {
	return &Queries{DB: db}
}

const getUserBalanceQuery = `
	SELECT a.id, a.balance, a.created_at 
	FROM Accounts a
	WHERE a.id = $1
`

func (q *Queries) GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	dtoAccount := &DTO{}

	row := q.DB.QueryRowxContext(ctx, getUserBalanceQuery, userID)
	err := row.StructScan(dtoAccount)
	if err != nil {
		return nil, err // TODO: handle ErrNoRows
	}

	acc := &domain.Account{
		ID:        dtoAccount.ID,
		Balance:   dtoAccount.Balance,
		CreatedAt: dtoAccount.CreatedAt,
	}

	return acc, nil
}
