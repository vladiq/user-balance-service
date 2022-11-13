package mapper

import (
	"testing"
	"time"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMakeTransferEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Transaction{
		FromID: validUUID,
		ToID:   validUUID,
		Amount: moneyAmount,
	}
	req := request.MakeTransfer{
		FromID: validUUID,
		ToID:   validUUID,
		Amount: moneyAmount,
	}
	require.Equal(t, expected, Transfer{}.MakeTransferEntity(req))
}

func TestUserMonthlyReport(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Transfer{
		AccountID: validUUID,
		CreatedAt: time.Date(0, time.Month(0), 1, 0, 0, 0, 0, time.UTC),
	}
	req := request.UserMonthlyReport{
		AccountID: validUUID,
		Year:      0,
		Month:     0,
	}
	require.Equal(t, expected, Transfer{}.UserMonthlyReport(req))
}

func TestEntityToReportEntry(t *testing.T) {
	expected := &response.GetUserMonthlyReport{
		Timestamp: time.Time{},
		IsAccrual: false,
		Info:      "",
		Amount:    moneyAmount,
	}
	entity := domain.Transfer{
		CreatedAt: time.Time{},
		IsAccrual: false,
		Info:      "",
		Amount:    moneyAmount,
	}
	require.Equal(t, expected, Transfer{}.EntityToReportEntry(entity))
}
