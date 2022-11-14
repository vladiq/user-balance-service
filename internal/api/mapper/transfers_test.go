package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/testdata"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMakeTransferEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Transaction{
		FromID: validUUID,
		ToID:   validUUID,
		Amount: 5.5,
	}
	req := request.MakeTransfer{
		FromID: validUUID,
		ToID:   validUUID,
		Amount: 5.5,
	}
	require.Equal(t, expected, Transfer{}.MakeTransferEntity(req))
}
