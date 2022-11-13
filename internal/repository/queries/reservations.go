package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const createReservationQuery = `
	INSERT INTO reservations(account_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4)
`

func CreateReservation(ctx context.Context, tx *sql.Tx, accountID, serviceID, orderID uuid.UUID, amount float64) error {
	_, err := tx.ExecContext(ctx, createReservationQuery, accountID, serviceID, orderID, amount)
	if err != nil {
		return fmt.Errorf("executing create new reservation query: %w", err)
	}
	return nil
}

const deleteReservationQuery = `
	DELETE FROM reservations r WHERE r.id = $1 RETURNING r.account_id, r.service_id, r.amount;
`

func DeleteReservation(ctx context.Context, tx *sql.Tx, id uuid.UUID) (uuid.UUID, uuid.UUID, float64, error) {
	var (
		accountID string
		serviceID string
		amount    float64
	)
	row := tx.QueryRowContext(ctx, deleteReservationQuery, id)
	if err := row.Scan(&accountID, &serviceID, &amount); err != nil {
		return uuid.UUID{}, uuid.UUID{}, 0, fmt.Errorf("executing query to delete a reservation: %w", err)
	}

	accID, _ := uuid.Parse(accountID)
	srvcID, _ := uuid.Parse(serviceID)

	return accID, srvcID, amount, nil
}
