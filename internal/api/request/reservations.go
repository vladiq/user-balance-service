package request

import (
	"encoding/json"
	"fmt"
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
