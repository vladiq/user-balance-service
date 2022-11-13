package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
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

type DepositFunds struct {
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
}

func (df *DepositFunds) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(df); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}
	return df.validate()
}

func (df *DepositFunds) validate() error {
	if err := validation.Validate(df.Amount, validation.Required, validation.Min(float64(0))); err != nil {
		return fmt.Errorf("validating amount: %w", err)
	}
	// TODO add validation of uuid correctness to all validate() methods!!!
	return nil
}

type WithdrawFunds struct {
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
}

func (wf *WithdrawFunds) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(wf); err != nil {
		return fmt.Errorf("binding body: %w", err)
	}
	return wf.validate()
}

func (wf *WithdrawFunds) validate() error {
	if err := validation.Validate(wf.Amount, validation.Required, validation.Min(float64(0))); err != nil {
		return fmt.Errorf("validating amount: %w", err)
	}
	// TODO: add validation of uuid correctness to all validate() methods!!!
	return nil
}
