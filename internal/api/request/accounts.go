package request

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

	"github.com/go-ozzo/ozzo-validation"
)

type CreateAccount struct {
	Amount float64 `json:"amount"`
}

func (r *CreateAccount) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}
	return r.validate()
}

func (r *CreateAccount) validate() error {
	if err := validation.Validate(r.Amount, validation.Required, validation.Min(float64(0))); err != nil {
		return fmt.Errorf("validating amount: %w", err)
	}
	return nil
}

type GetAccount struct {
	ID uuid.UUID `json:"id"`
}

func (r *GetAccount) Bind(req *http.Request) error {
	accountIDParam := chi.URLParam(req, "accountID")

	if accountID, err := uuid.Parse(accountIDParam); err != nil {
		return fmt.Errorf("binding body: %w", err)
	} else {
		r.ID = accountID
	}

	return nil
}
