package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Account struct {
}

func (m Account) CreateAccountEntity(request request.CreateAccount) domain.Account {
	return domain.Account{
		Balance: request.Amount,
	}
}

func (m Account) GetAccountEntity(request request.GetAccount) domain.Account {
	return domain.Account{
		ID: request.ID,
	}
}

func (m Account) GetAccountResponse(entity domain.Account) response.GetAccount {
	return response.GetAccount{
		ID:      entity.ID,
		Balance: entity.Balance,
	}
}
