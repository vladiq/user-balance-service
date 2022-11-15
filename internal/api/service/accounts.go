package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type accountsRepo interface {
	CreateAccount(ctx context.Context, entity domain.Account) error
	GetAccount(ctx context.Context, entity domain.Account) (*domain.Account, error)
	DepositFunds(ctx context.Context, entity domain.Account) error
	WithdrawFunds(ctx context.Context, entity domain.Account) error
}

type accounts struct {
	repo   accountsRepo
	mapper mapper.Account
}

func NewAccounts(repo accountsRepo) *accounts {
	return &accounts{repo: repo, mapper: mapper.Account{}}
}

func (a *accounts) CreateAccount(ctx context.Context, r request.CreateAccount) error {
	return a.repo.CreateAccount(ctx, a.mapper.CreateAccountEntity(r))
}

func (a *accounts) GetAccount(ctx context.Context, r request.GetAccount) (response.GetAccount, error) {
	if account, err := a.repo.GetAccount(ctx, a.mapper.GetAccountEntity(r)); err != nil {
		return response.GetAccount{}, err
	} else {
		return a.mapper.GetAccountResponse(*account), nil
	}
}

func (a *accounts) DepositFunds(ctx context.Context, r request.DepositFunds, accountID uuid.UUID) error {
	return a.repo.DepositFunds(ctx, a.mapper.DepositFundsEntity(r, accountID))
}

func (a *accounts) WithdrawFunds(ctx context.Context, r request.WithdrawFunds, accountID uuid.UUID) error {
	return a.repo.WithdrawFunds(ctx, a.mapper.WithdrawFundsEntity(r, accountID))
}
