package handlers

import (
	"context"
	"github.com/google/uuid"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type accountsService interface {
	CreateAccount(ctx context.Context, request request.CreateAccount) error
	GetAccount(ctx context.Context, request request.GetAccount) (response.GetAccount, error)
	DepositFunds(ctx context.Context, request request.DepositFunds, accountID uuid.UUID) error
	WithdrawFunds(ctx context.Context, request request.WithdrawFunds, accountID uuid.UUID) error
}

type accounts struct {
	service accountsService
}

func NewAccounts(service accountsService) *accounts {
	return &accounts{service: service}
}

func (h *accounts) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{accountID}", h.getAccount)
	r.Post("/", h.createAccount)
	r.Put("/deposit/{accountID}", h.depositFunds)
	r.Put("/withdraw/{accountID}", h.withdrawFunds)

	return r
}

// createAccount creates a user account with provided amount of money
// @Summary Create a user account with given balance and add an entry to the transfers table
// @Tags    Accounts
// @ID account-create
// @Accept json
// @Param amount body request.CreateAccount true "amount of money on the new account"
// @Success 201 "Created"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router  /accounts [post]
func (h *accounts) createAccount(w http.ResponseWriter, r *http.Request) {
	var req request.CreateAccount

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if err := h.service.CreateAccount(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// getAccount gets a user account by id
// @Summary Get a user account information by its ID
// @Tags    Accounts
// @ID account-get
// @Produce json
// @Param id path string true "account uuid"
// @Success 200 {object} response.GetAccount "Account info"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router  /accounts/{id} [get]
func (h *accounts) getAccount(w http.ResponseWriter, r *http.Request) {
	var req request.GetAccount

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if account, err := h.service.GetAccount(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)
		return
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, account)
	}
}

// depositFunds deposit funds to a given account
// @Summary Deposit funds to an account and add en entry to the transfers table
// @Tags    Accounts
// @ID account-deposit
// @Accept json
// @Param id path string true "account uuid"
// @Param amount body request.DepositFunds true "amount of money"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /accounts/deposit/{id} [put]
func (h *accounts) depositFunds(w http.ResponseWriter, r *http.Request) {
	var req request.DepositFunds

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	accountID, err := uuid.Parse(chi.URLParam(r, "accountID"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
	}

	if err := h.service.DepositFunds(r.Context(), req, accountID); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// depositFunds withdraw funds from a given account
// @Summary Withdraw funds from an account and add en entry to the transfers table
// @Tags    Accounts
// @ID account-withdraw
// @Accept json
// @Param id path string true "account uuid"
// @Param amount body request.WithdrawFunds true "amount of money"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /accounts/withdraw/{id} [put]
func (h *accounts) withdrawFunds(w http.ResponseWriter, r *http.Request) {
	var req request.WithdrawFunds

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	accountID, err := uuid.Parse(chi.URLParam(r, "accountID"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
	}

	if err := h.service.WithdrawFunds(r.Context(), req, accountID); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
