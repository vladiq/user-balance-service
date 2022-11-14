package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"sort"
)

const pageSize = 4

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
	SELECT id, account_id, is_accrual, amount, info, created_at
	FROM transfers
	WHERE account_id = $1 AND id >= $2
	ORDER BY id
	LIMIT $3
`

func GetTransfers(ctx context.Context, tx *sql.Tx, entity domain.Transfer, pageID uuid.UUID, orderBy string) ([]*domain.Transfer, uuid.UUID, error) {
	rows, err := tx.QueryContext(ctx, getTransfersQuery, entity.AccountID, pageID, pageSize+1)
	if err != nil {
		return nil, uuid.UUID{}, fmt.Errorf("executing query to get user transfers entries: %w", err)
	}

	var transfers []*domain.Transfer
	for rows.Next() {
		var t domain.Transfer
		if err := rows.Scan(&t.ID, &t.AccountID, &t.IsAccrual, &t.Amount, &t.Info, &t.CreatedAt); err != nil {
			return nil, uuid.UUID{}, fmt.Errorf("scanning user transfer entries: %w", err)
		}
		transfers = append(transfers, &t)
	}

	var nextPageID uuid.UUID
	if len(transfers) == pageSize+1 {
		nextPageID = transfers[len(transfers)-1].ID
		transfers = transfers[:pageSize]
	}

	switch orderBy {
	case "amount":
		sort.Slice(transfers, func(i, j int) bool {
			return transfers[i].Amount < transfers[j].Amount
		})
	case "date":
		sort.Slice(transfers, func(i, j int) bool {
			return transfers[i].CreatedAt.Before(transfers[j].CreatedAt)
		})
	}

	return transfers, nextPageID, nil
}
