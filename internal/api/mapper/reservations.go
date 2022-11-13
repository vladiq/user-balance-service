package mapper

import (
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Reservation struct {
}

func (m Reservation) MakeCreateReservationEntity(request request.CreateReservation) domain.Reservation {
	userID, _ := uuid.Parse(request.UserID)
	serviceID, _ := uuid.Parse(request.ServiceID)
	orderID, _ := uuid.Parse(request.OrderID)

	return domain.Reservation{
		AccountID: userID,
		ServiceID: serviceID,
		OrderID:   orderID,
		Amount:    request.Amount,
	}
}
