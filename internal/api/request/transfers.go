package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

	if err := validation.Validate(r.ToID, validation.Required, is.UUID); err != nil {
		return fmt.Errorf("validating ToID: %w", err)
	}

	if err := validation.Validate(r.FromID, validation.Required, is.UUID); err != nil {
		return fmt.Errorf("validating FromID: %w", err)
	}

	return nil
}

type GetUserTransfers struct {
	AccountID uuid.UUID
	OrderBy   string
}

func (r *GetUserTransfers) Bind(req *http.Request) error {
	accIDParam := chi.URLParam(req, "accountID")
	accID, err := uuid.Parse(accIDParam)
	if err != nil {
		return fmt.Errorf("parsing uuid: %w", err)
	}
	r.AccountID = accID

	orderBy := req.URL.Query().Get("order-by")
	r.OrderBy = orderBy

	return r.validate()
}

func (r *GetUserTransfers) validate() error {
	if err := validation.Validate(r.AccountID, validation.Required, is.UUID); err != nil {
		return fmt.Errorf("validating AccountID: %w", err)
	}

	if r.OrderBy != "" && r.OrderBy != "amount" && r.OrderBy != "date" {
		return errors.New("wrong order-by key format: use date,amount")
	}

	return nil
}
