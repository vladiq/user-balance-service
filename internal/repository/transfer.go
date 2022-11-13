package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/vladiq/user-balance-service/internal/constant"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/repository/queries"

	"github.com/jmoiron/sqlx"
)

type transferRepository struct {
	DB *sqlx.DB
}

func NewTransferRepository(db *sqlx.DB) *transferRepository {
	return &transferRepository{
		DB: db,
	}
}

func (r *transferRepository) MakeTransaction(ctx context.Context, entity domain.Transaction) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	acc, err := queries.GetAccount(ctx, tx, entity.FromID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("getting account: %w", err)
	}

	if acc.Balance-entity.Amount < 0 {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("not enough funds to make a transaction: %w", constant.ErrBadRequest)
	}

	if err := queries.WithdrawFunds(ctx, tx, entity.FromID, entity.Amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("withdrawing funds from an account: %w", err)
	}

	if err := queries.DepositFunds(ctx, tx, entity.ToID, entity.Amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("depositing funds to an account: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		entity.FromID,
		false,
		entity.Amount,
		"transfer money to another user",
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("adding a money transfer record: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		entity.ToID,
		true,
		entity.Amount,
		"money transfer from another user",
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("adding a money transfer record: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
