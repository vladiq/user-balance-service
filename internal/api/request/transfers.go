package request

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"net/http"
)

type MakeTransfer struct {
	FromID uuid.UUID `json:"from_id"`
	ToID   uuid.UUID `json:"to_id"`
	Amount float64   `json:"amount"`
}

func (r *MakeTransfer) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}
	return r.validate()
}

func (r *MakeTransfer) validate() error {
	if err := validation.Validate(r.Amount, validation.Required, validation.Min(float64(0))); err != nil {
		return fmt.Errorf("validating amount: %w", err)
	}
	return nil
}
