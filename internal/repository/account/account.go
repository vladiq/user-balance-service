package account

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type accountRepository struct {
	*Queries
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewAccountRepository(logger zerolog.Logger, db *sqlx.DB) *accountRepository {
	return &accountRepository{
		DB:      db,
		Queries: New(db),
		logger:  logger,
	}
}
