package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Reservation struct {
}

func (m Reservation) MakeCreateReservationEntity(r request.CreateReservation) domain.Reservation {
	return domain.Reservation{
		AccountID: r.UserID,
		ServiceID: r.ServiceID,
		OrderID:   r.OrderID,
		Amount:    r.Amount,
	}
}

func (m Reservation) MakeCancelReservationEntity(r request.CancelReservation) domain.Reservation {
	return domain.Reservation{ID: r.ID}
}

func (m Reservation) MakeConfirmReservationEntity(r request.ConfirmReservation) domain.Reservation {
	return domain.Reservation{ID: r.ID}
}
