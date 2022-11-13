package mapper

import (
	"testing"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	validUUIDString = "129cec45-9a25-44dc-98c2-d8804c1485bf"
	moneyAmount     = 5.5
)

func TestCreateAccountEntity(t *testing.T) {
	expected := domain.Account{
		Balance: moneyAmount,
	}
	req := request.CreateAccount{
		Amount: moneyAmount,
	}
	require.Equal(t, expected, Account{}.CreateAccountEntity(req))
}

func TestGetAccountEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Account{
		ID: validUUID,
	}
	req := request.GetAccount{ID: validUUID}
	require.Equal(t, expected, Account{}.GetAccountEntity(req))
}

func TestGetAccountResponse(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := response.GetAccount{
		ID:      validUUID,
		Balance: moneyAmount,
	}
	entity := domain.Account{
		ID:      validUUID,
		Balance: moneyAmount,
	}
	require.Equal(t, expected, Account{}.GetAccountResponse(entity))
}

func TestDepositFundsEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Account{
		ID:      validUUID,
		Balance: moneyAmount,
	}
	req := request.DepositFunds{
		ID:     validUUID,
		Amount: moneyAmount,
	}
	require.Equal(t, expected, Account{}.DepositFundsEntity(req))
}

func TestWithdrawFundsEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Account{
		ID:      validUUID,
		Balance: moneyAmount,
	}
	req := request.WithdrawFunds{
		ID:     validUUID,
		Amount: moneyAmount,
	}
	require.Equal(t, expected, Account{}.WithdrawFundsEntity(req))
}
