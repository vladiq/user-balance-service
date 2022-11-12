package reservation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Queries struct {
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewQueries(db *sqlx.DB, logger zerolog.Logger) *Queries {
	return &Queries{DB: db, logger: logger}
}

const createReservationQuery = `
	INSERT INTO reservations(account_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error {
	q.logger.Info().Msg("Creating new reservation")

	result, err := q.DB.ExecContext(ctx, createReservationQuery, userID, serviceID, orderID, amount)
	if err != nil {
		return fmt.Errorf("executing create new reservation query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting number of affected rows: %w", err)
	}

	q.logger.Debug().
		Str("query", createReservationQuery).
		Str("userID", userID.String()).
		Str("serviceID", serviceID.String()).
		Str("orderID", orderID.String()).
		Float64("amount", amount).
		Int64("rows affected", rowsAffected).
		Msg("create new reservation")

	return nil
}
