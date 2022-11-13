package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
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

type UserMonthlyReport struct {
	AccountID uuid.UUID
	Year      int
	Month     int
}

func (r *UserMonthlyReport) Bind(req *http.Request) error {
	accIDParam := chi.URLParam(req, "accountID")
	accID, err := uuid.Parse(accIDParam)
	if err != nil {
		return fmt.Errorf("parsing uuid: %w", err)
	}
	r.AccountID = accID

	year, err := strconv.Atoi(req.URL.Query().Get("year"))
	if err != nil {
		return fmt.Errorf("getting year from query: %w", err)
	}
	r.Year = year

	month, err := strconv.Atoi(req.URL.Query().Get("month"))
	if err != nil {
		return fmt.Errorf("getting month from query: %w", err)
	}
	r.Month = month

	return r.validate()
}

func (r *UserMonthlyReport) validate() error {
	if err := validation.Validate(r.Month, validation.Required, validation.Max(12)); err != nil {
		return fmt.Errorf("validating month: %w", err)
	}

	return nil
}
