package mapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

func TestCreateAccountEntity(t *testing.T) {
	expected := domain.Account{
		Balance: 5.,
	}

	req := request.CreateAccount{
		Amount: 5,
	}

	require.Equal(t, expected, Account{}.CreateAccountEntity(req))
}
