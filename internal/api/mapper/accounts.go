package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Account struct {
}

func (m Account) CreateAccountEntity(r request.CreateAccount) domain.Account {
	return domain.Account{
		Balance: r.Amount,
	}
}

func (m Account) GetAccountEntity(r request.GetAccount) domain.Account {
	return domain.Account{
		ID: r.ID,
	}
}

func (m Account) GetAccountResponse(entity domain.Account) response.GetAccount {
	return response.GetAccount{
		ID:      entity.ID,
		Balance: entity.Balance,
	}
}

func (m Account) DepositFundsEntity(r request.DepositFunds) domain.Account {
	return domain.Account{
		ID:      r.ID,
		Balance: r.Amount,
	}
}

func (m Account) WithdrawFundsEntity(r request.WithdrawFunds) domain.Account {
	return domain.Account{
		ID:      r.ID,
		Balance: r.Amount,
	}
}
