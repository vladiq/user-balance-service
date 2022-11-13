package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

const addTransferDataQuery = `
	INSERT INTO transfers(account_id, is_accrual, amount, info) VALUES ($1, $2, $3, $4)
`

func AddTransferData(ctx context.Context, tx *sql.Tx, accountID uuid.UUID, isAccrual bool, amount float64, info string) error {
	if _, err := tx.ExecContext(ctx, addTransferDataQuery, accountID, isAccrual, amount, info); err != nil {
		return fmt.Errorf("executing query to create a money transfer record: %w", err)
	}

	return nil
}
