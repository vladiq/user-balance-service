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

type accountRepository struct {
	DB *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *accountRepository {
	return &accountRepository{
		DB: db,
	}
}

const (
	depositMessage    = "merchant deposit"
	withdrawalMessage = "merchant withdrawal"
)

func (r *accountRepository) GetAccount(ctx context.Context, entity domain.Account) (*domain.Account, error) {
	opts := sql.TxOptions{
		ReadOnly:  true,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("beginning transaction: %w", err)
	}

	acc, err := queries.GetAccount(ctx, tx, entity.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("rolling transaction back: %w", err)
		}
		return nil, fmt.Errorf("getting account: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return acc, nil
}

func (r *accountRepository) CreateAccount(ctx context.Context, entity domain.Account) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	accountID, err := queries.CreateAccount(ctx, tx, entity.Balance)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("executing create account query: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		accountID,
		true,
		entity.Balance,
		"merchant deposit on account creation",
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

func (r *accountRepository) DepositFunds(ctx context.Context, entity domain.Account) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	if err := queries.DepositFunds(ctx, tx, entity.ID, entity.Balance); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("executing create account query: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		entity.ID,
		true,
		entity.Balance,
		depositMessage,
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

func (r *accountRepository) WithdrawFunds(ctx context.Context, entity domain.Account) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	acc, err := queries.GetAccount(ctx, tx, entity.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("getting account: %w", err)
	}

	if acc.Balance-entity.Balance < 0 {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("not enough funds to withdraw: %w", constant.ErrBadRequest)
	}

	if err := queries.WithdrawFunds(ctx, tx, entity.ID, entity.Balance); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("executing create account query: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		entity.ID,
		false,
		entity.Balance,
		withdrawalMessage,
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
