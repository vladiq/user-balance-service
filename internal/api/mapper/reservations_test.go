package mapper

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/testdata"
	"testing"
)

func TestMakeCreateReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Reservation{
		AccountID: validUUID,
		ServiceID: validUUID,
		OrderID:   validUUID,
		Amount:    5.5,
	}
	req := request.CreateReservation{
		UserID:    validUUID,
		ServiceID: validUUID,
		OrderID:   validUUID,
		Amount:    5.5,
	}
	require.Equal(t, expected, Reservation{}.MakeCreateReservationEntity(req))
}

func TestMakeCancelReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Reservation{
		ID: validUUID,
	}
	req := request.CancelReservation{
		ID: validUUID,
	}
	require.Equal(t, expected, Reservation{}.MakeCancelReservationEntity(req))
}

func TestMakeConfirmReservationEntity(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := domain.Reservation{
		ID: validUUID,
	}
	req := request.ConfirmReservation{
		ID: validUUID,
	}
	require.Equal(t, expected, Reservation{}.MakeConfirmReservationEntity(req))
}
