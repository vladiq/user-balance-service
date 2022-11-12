package transfer

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type transferRepository struct {
	*Queries
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewTransferRepository(logger zerolog.Logger, db *sqlx.DB) *transferRepository {
	return &transferRepository{
		DB:      db,
		Queries: NewQueries(db, logger),
		logger:  logger,
	}
}
