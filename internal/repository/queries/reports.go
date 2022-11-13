package queries

import (
	"context"
	"database/sql"
	"fmt"
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
