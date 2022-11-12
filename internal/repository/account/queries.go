package account

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewQueries(db *sqlx.DB, logger zerolog.Logger) *Queries {
	return &Queries{DB: db, logger: logger}
}

const getUserQuery = `
	SELECT a.id, a.balance, a.created_at 
	FROM Accounts a
	WHERE a.id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	q.logger.Info().Msg("Querying user by id")

	dtoAccount := &DTO{}

	row := q.DB.QueryRowxContext(ctx, getUserQuery, userID)
	err := row.StructScan(dtoAccount)
	if err != nil {
		return nil, fmt.Errorf("scanning user: %w", err) // TODO: handle ErrNoRows
	}

	q.logger.Debug().
		Str("query", createAccountQuery).
		Msg("query user by id")

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

func (q *Queries) CreateAccount(ctx context.Context, amount float64) error {
	q.logger.Info().Msg("Creating new account")

	result, err := q.DB.ExecContext(ctx, createAccountQuery, amount)
	if err != nil {
		return fmt.Errorf("executing account creation query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting number of affected rows: %w", err)
	}

	q.logger.Debug().
		Str("query", createAccountQuery).
		Float64("amount", amount).
		Int64("rows affected", rowsAffected).
		Msg("create new account")

	return nil
}

const depositFundsQuery = `
	UPDATE accounts SET balance = balance + $1 WHERE id = $2
`

func (q *Queries) DepositFunds(ctx context.Context, userID uuid.UUID, amount float64) error {
	q.logger.Info().Msg("Depositing money to account")

	result, err := q.DB.ExecContext(ctx, depositFundsQuery, amount, userID.String())
	if err != nil {
		return fmt.Errorf("executing query to deposit funds to an account: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting number of affected rows: %w", err)
	}

	q.logger.Debug().
		Str("query", depositFundsQuery).
		Float64("amount", amount).
		Int64("rows affected", rowsAffected).
		Msg("deposit money to account")

	return nil
}

const withdrawFundsQuery = `
	UPDATE accounts SET balance = balance - $1 WHERE id = $2
`

func (q *Queries) WithdrawFunds(ctx context.Context, userID uuid.UUID, amount float64) error {
	q.logger.Info().Msg("Withdrawing money from account")

	result, err := q.DB.ExecContext(ctx, withdrawFundsQuery, amount, userID.String())
	if err != nil {
		return fmt.Errorf("executing query to withdraw money from an account: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting number of affected rows: %w", err)
	}

	q.logger.Debug().
		Str("query", withdrawFundsQuery).
		Float64("amount", amount).
		Int64("rows affected", rowsAffected).
		Msg("withdraw money from account")

	return nil
}
