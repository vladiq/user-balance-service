package service

import (
	"context"
	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type accountsRepo interface {
	CreateAccount(ctx context.Context, entity domain.Account) error
	GetAccount(ctx context.Context, entity domain.Account) (*domain.Account, error)
}

type accounts struct {
	repo   accountsRepo
	mapper mapper.Account
}

func NewAccounts(repo accountsRepo) *accounts {
	return &accounts{repo: repo, mapper: mapper.Account{}}
}

func (a *accounts) CreateAccount(ctx context.Context, request request.CreateAccount) error {
	return a.repo.CreateAccount(ctx, a.mapper.CreateAccountEntity(request))
}

func (a *accounts) GetAccount(ctx context.Context, request request.GetAccount) (response.GetAccount, error) {
	if account, err := a.repo.GetAccount(ctx, a.mapper.GetAccountEntity(request)); err != nil {
		return response.GetAccount{}, err
	} else {
		return a.mapper.GetAccountResponse(*account), nil
	}
}
