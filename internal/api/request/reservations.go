package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CreateReservation struct {
	UserID    uuid.UUID `json:"user_id"`
	ServiceID uuid.UUID `json:"service_id"`
	OrderID   uuid.UUID `json:"order_id"`
	Amount    float64   `json:"amount"`
}

func (r *CreateReservation) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}

	return r.validate()
}

func (r *CreateReservation) validate() error {
	if err := validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, is.UUID),
		validation.Field(&r.ServiceID, validation.Required, is.UUID),
		validation.Field(&r.OrderID, validation.Required, is.UUID),
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

	return r.validate()
}

func (r *CancelReservation) validate() error {
	if err := validation.Validate(r.ID, validation.Required, is.UUID); err != nil {
		return fmt.Errorf("validating ID: %w", err)
	}
	return nil
}

type ConfirmReservation struct {
	ID uuid.UUID `json:"id"`
}

func (r *ConfirmReservation) Bind(req *http.Request) error {
	reservationIDParam := chi.URLParam(req, "reservationID")

	if reservationID, err := uuid.Parse(reservationIDParam); err != nil {
		return fmt.Errorf("binding body: %w", err)
	} else {
		r.ID = reservationID
	}

	return r.validate()
}

func (r *ConfirmReservation) validate() error {
	if err := validation.Validate(r.ID, validation.Required, is.UUID); err != nil {
		return fmt.Errorf("validating ID: %w", err)
	}
	return nil
}
