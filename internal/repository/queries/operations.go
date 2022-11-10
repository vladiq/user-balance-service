package queries

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/repository/repo-dto"
)

// AddFundsToAccount Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
func (q *Queries) AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error {
	panic(errors.New("not implemented"))
	return nil
}

// MakeReservation Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
func (q *Queries) MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error {
	panic(errors.New("not implemented"))
	return nil
}

// AcceptReservation Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
func (q *Queries) AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error {
	panic(errors.New("not implemented"))
	return nil
}

const getUserBalanceQuery = `
	SELECT a.id, a.balance, a.created_at 
	FROM Accounts a
	WHERE a.id = $1
`

func (q *Queries) GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	dtoAccount := &repo_dto.Account{}

	row := q.DB.QueryRowxContext(ctx, getUserBalanceQuery, userID)
	err := row.StructScan(dtoAccount)
	if err != nil {
		return nil, err // handle ErrNoRows?
	}

	account := &domain.Account{
		ID:        dtoAccount.ID,
		Balance:   dtoAccount.Balance,
		CreatedAt: dtoAccount.CreatedAt,
	}

	return account, nil
}
