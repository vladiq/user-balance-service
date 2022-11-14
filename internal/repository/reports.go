package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/repository/queries"
)

type reportRepository struct {
	DB *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *reportRepository {
	return &reportRepository{
		DB: db,
	}
}

func (r *reportRepository) GetReport(ctx context.Context, month int) ([]*domain.Report, error) {
	opts := sql.TxOptions{
		ReadOnly:  true,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("beginning transaction: %w", err)
	}

	entries, err := queries.GetReportEntries(ctx, tx, month)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("rolling transaction back: %w", err)
		}
		return nil, fmt.Errorf("getting report entries: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return entries, nil
}
