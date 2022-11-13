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

type reservationRepository struct {
	DB *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) *reservationRepository {
	return &reservationRepository{
		DB: db,
	}
}

func (r *reservationRepository) Create(ctx context.Context, e domain.Reservation) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	if err := queries.CreateReservation(ctx, tx, e.AccountID, e.ServiceID, e.OrderID, e.Amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("creating reservation: %w", err)
	}

	account, err := queries.GetAccount(ctx, tx, e.AccountID)

	if account.Balance-e.Amount < 0 {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("not enough funds to make a reservation: %w", constant.ErrBadRequest)
	}

	if err := queries.WithdrawFunds(ctx, tx, e.AccountID, e.Amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("withdrawing funds: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		e.AccountID,
		false,
		e.Amount,
		"withdrawing money to make a reservation",
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("adding a money transfer record: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}

func (r *reservationRepository) Cancel(ctx context.Context, entity domain.Reservation) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	accountID, _, amount, err := queries.DeleteReservation(ctx, tx, entity.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("deleting reservation entry: %w", err)
	}

	if err := queries.DepositFunds(ctx, tx, accountID, amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("returning funds to account balance: %w", err)
	}

	if err := queries.AddTransferData(
		ctx,
		tx,
		accountID,
		true,
		amount,
		"payback after cancelling reservation",
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("adding a money transfer record: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}

func (r *reservationRepository) Confirm(ctx context.Context, entity domain.Reservation) error {
	opts := sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelSerializable,
	}

	tx, err := r.DB.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}

	_, serviceID, amount, err := queries.DeleteReservation(ctx, tx, entity.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("deleting reservation entry: %w", err)
	}

	if err := queries.CreateReportEntry(ctx, tx, serviceID, amount); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}
		return fmt.Errorf("creating a service report entry: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}
