package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
)

type Reservation struct {
}

func (m Reservation) MakeCreateReservationEntity(r request.CreateReservation) domain.Reservation {
	userID, _ := uuid.Parse(r.UserID)
	serviceID, _ := uuid.Parse(r.ServiceID)
	orderID, _ := uuid.Parse(r.OrderID)

	return domain.Reservation{
		AccountID: userID,
		ServiceID: serviceID,
		OrderID:   orderID,
		Amount:    r.Amount,
	}
}

func (m Reservation) MakeCancelReservationEntity(r request.CancelReservation) domain.Reservation {
	return domain.Reservation{ID: r.ID}
}

func (m Reservation) MakeConfirmReservationEntity(r request.ConfirmReservation) domain.Reservation {
	return domain.Reservation{ID: r.ID}
}
