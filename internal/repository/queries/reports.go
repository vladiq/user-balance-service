package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
)

const createReportEntryQuery = `
	INSERT INTO reports(service_id, amount) VALUES ($1, $2)
`

func CreateReportEntry(ctx context.Context, tx *sql.Tx, serviceID uuid.UUID, amount float64) error {
	if _, err := tx.ExecContext(ctx, createReportEntryQuery, serviceID, amount); err != nil {
		return fmt.Errorf("executing report entry creating query: %w", err)
	}
	return nil
}

const getReportEntriesQuery = `
	SELECT r.service_id, SUM(r.amount) 
	FROM reports r
	WHERE EXTRACT(MONTH FROM r.created_at) = $1
	GROUP BY r.service_id
`

func GetReportEntries(ctx context.Context, tx *sql.Tx, month int) ([]*domain.Report, error) {
	rows, err := tx.QueryContext(ctx, getReportEntriesQuery, month)
	if err != nil {
		return nil, fmt.Errorf("executing query to get service report entries: %w", err)
	}

	var entries []*domain.Report
	for rows.Next() {
		var entry domain.Report
		err := rows.Scan(&entry.ServiceID, &entry.Amount)
		if err != nil {
			return nil, fmt.Errorf("scanning service report entry: %w", err)
		}
		entries = append(entries, &entry)
	}

	return entries, nil
}
