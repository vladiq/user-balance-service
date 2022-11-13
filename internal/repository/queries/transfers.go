package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/vladiq/user-balance-service/internal/domain"

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

const getTransfersQuery = `
	SELECT t.is_accrual, t.amount, t.info, t.created_at
	FROM transfers t 
	WHERE 
	    t.account_id = $1
	    AND EXTRACT(YEAR FROM t.created_at) = $2
	  	AND EXTRACT(MONTH FROM t.created_at) = $3
`

func GetTransfers(ctx context.Context, tx *sql.Tx, entity domain.Transfer) ([]*domain.Transfer, error) {
	rows, err := tx.QueryContext(ctx, getTransfersQuery, entity.AccountID, entity.CreatedAt.Year(), entity.CreatedAt.Month())
	if err != nil {
		return nil, fmt.Errorf("executing query to get user transfers entries: %w", err)
	}

	var transfers []*domain.Transfer

	for rows.Next() {
		var t domain.Transfer
		if err := rows.Scan(&t.IsAccrual, &t.Amount, &t.Info, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("scanning user transfer entries: %w", err)
		}
		transfers = append(transfers, &t)
	}

	return transfers, nil
}
