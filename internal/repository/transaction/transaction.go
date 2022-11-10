package transaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type transactionRepository struct {
	*Queries
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewTransactionRepository(logger zerolog.Logger, db *sqlx.DB) *transactionRepository {
	return &transactionRepository{
		DB:      db,
		Queries: New(db),
		logger:  logger,
	}
}
