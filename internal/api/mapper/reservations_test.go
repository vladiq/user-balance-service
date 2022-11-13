package mapper

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
	"testing"
)

func TestMakeCreateReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Reservation{
		AccountID: validUUID,
		ServiceID: validUUID,
		OrderID:   validUUID,
		Amount:    moneyAmount,
	}
	req := request.CreateReservation{
		UserID:    validUUID,
		ServiceID: validUUID,
		OrderID:   validUUID,
		Amount:    moneyAmount,
	}
	require.Equal(t, expected, Reservation{}.MakeCreateReservationEntity(req))
}

func TestMakeCancelReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Reservation{
		ID: validUUID,
	}
	req := request.CancelReservation{
		ID: validUUID,
	}
	require.Equal(t, expected, Reservation{}.MakeCancelReservationEntity(req))
}

func TestMakeConfirmReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(validUUIDString)
	expected := domain.Reservation{
		ID: validUUID,
	}
	req := request.ConfirmReservation{
		ID: validUUID,
	}
	require.Equal(t, expected, Reservation{}.MakeConfirmReservationEntity(req))
}
