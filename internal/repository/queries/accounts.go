package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/repository/models"
)

const getUserQuery = `
	SELECT a.id, a.balance, a.created_at 
	FROM Accounts a
	WHERE a.id = $1
`

func GetAccount(ctx context.Context, tx *sql.Tx, userID uuid.UUID) (*domain.Account, error) {
	dtoAccount := &models.Accounts{}

	row := tx.QueryRowContext(ctx, getUserQuery, userID)
	err := row.Scan(&dtoAccount.ID, &dtoAccount.Balance, &dtoAccount.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("scanning user: %w", err) // TODO: handle ErrNoRows
	}

	acc := &domain.Account{
		ID:        dtoAccount.ID,
		Balance:   dtoAccount.Balance,
		CreatedAt: dtoAccount.CreatedAt,
	}

	return acc, nil
}

const createAccountQuery = `
	insert into accounts(balance) values ($1)
`

func CreateAccount(ctx context.Context, tx *sql.Tx, amount float64) error {
	_, err := tx.ExecContext(ctx, createAccountQuery, amount)
	if err != nil {
		return fmt.Errorf("executing account creation query: %w", err)
	}

	return nil
}

const depositFundsQuery = `
	UPDATE accounts SET balance = balance + $1 WHERE id = $2
`

func DepositFunds(ctx context.Context, tx *sql.Tx, userID uuid.UUID, amount float64) error {
	_, err := tx.ExecContext(ctx, depositFundsQuery, amount, userID.String())
	if err != nil {
		return fmt.Errorf("executing query to deposit funds to an account: %w", err)
	}

	return nil
}

const withdrawFundsQuery = `
	UPDATE accounts SET balance = balance - $1 WHERE id = $2
`

func WithdrawFunds(ctx context.Context, tx *sql.Tx, userID uuid.UUID, amount float64) error {

	_, err := tx.ExecContext(ctx, withdrawFundsQuery, amount, userID.String())
	if err != nil {
		return fmt.Errorf("executing query to withdraw money from an account: %w", err)
	}

	return nil
}