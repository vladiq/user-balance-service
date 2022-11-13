package request

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

	"github.com/go-ozzo/ozzo-validation"
)

type CreateReservation struct {
	UserID    string  `json:"user_id"`
	ServiceID string  `json:"service_id"`
	OrderID   string  `json:"order_id"`
	Amount    float64 `json:"amount"`
}

func (r *CreateReservation) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}

	return r.validate()
}

func (r *CreateReservation) validate() error {
	if err := validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required), // validation.UID
		validation.Field(&r.ServiceID, validation.Required),
		validation.Field(&r.OrderID, validation.Required),
		validation.Field(&r.Amount, validation.Required, validation.Min(float64(0))),
	); err != nil {
		return fmt.Errorf("validating body: %w", err)
	}
	return nil
}

type CancelReservation struct {
	ID uuid.UUID `json:"id"`
}

func (r *CancelReservation) Bind(req *http.Request) error {
	reservationIDParam := chi.URLParam(req, "reservationID")

	if reservationID, err := uuid.Parse(reservationIDParam); err != nil {
		return fmt.Errorf("binding body: %w", err)
	} else {
		r.ID = reservationID
	}
	return nil
}
