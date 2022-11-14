package mapper

import (
	"github.com/vladiq/user-balance-service/internal/testdata"
	"testing"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountEntity(t *testing.T) {
	expected := domain.Account{
		Balance: 5.5,
	}
	req := request.CreateAccount{
		Amount: 5.5,
	}
	require.Equal(t, expected, Account{}.CreateAccountEntity(req))
}

func TestGetAccountEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Account{
		ID: validUUID,
	}
	req := request.GetAccount{ID: validUUID}
	require.Equal(t, expected, Account{}.GetAccountEntity(req))
}

func TestGetAccountResponse(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := response.GetAccount{
		ID:      validUUID,
		Balance: 5.5,
	}
	entity := domain.Account{
		ID:      validUUID,
		Balance: 5.5,
	}
	require.Equal(t, expected, Account{}.GetAccountResponse(entity))
}

func TestDepositFundsEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Account{
		ID:      validUUID,
		Balance: 5.5,
	}
	req := request.DepositFunds{
		ID:     validUUID,
		Amount: 5.5,
	}
	require.Equal(t, expected, Account{}.DepositFundsEntity(req))
}

func TestWithdrawFundsEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Account{
		ID:      validUUID,
		Balance: 5.5,
	}
	req := request.WithdrawFunds{
		ID:     validUUID,
		Amount: 5.5,
	}
	require.Equal(t, expected, Account{}.WithdrawFundsEntity(req))
}
