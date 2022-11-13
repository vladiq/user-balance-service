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
	fmt.Println(accountID)
	_, err := tx.ExecContext(ctx, createReservationQuery, accountID, serviceID, orderID, amount)
	if err != nil {
		return fmt.Errorf("executing create new reservation query: %w", err)
	}
	return nil
}
