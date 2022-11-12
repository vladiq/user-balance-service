package reservation

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type reservationRepository struct {
	*Queries
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewReservationRepository(logger zerolog.Logger, db *sqlx.DB) *reservationRepository {
	return &reservationRepository{
		DB:      db,
		Queries: NewQueries(db, logger),
		logger:  logger,
	}
}
