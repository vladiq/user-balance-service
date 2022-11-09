package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/repository/queries"
)

type repository struct {
	*queries.Queries
	DB *sqlx.DB
	l  zerolog.Logger
}

func NewRepo(db *sqlx.DB, l zerolog.Logger) Repository {
	return &repository{
		DB:      db,
		Queries: queries.New(db),
		l:       l,
	}
}

type Repository interface {
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
	MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error
	AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

//func (r *repository) execTx(ctx context.Context, fn func(queries2 *queries.Queries) error) error {
//	return nil
//}
