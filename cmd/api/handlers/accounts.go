package handlers

import (
	"context"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type accountsService interface {
	CreateAccount(ctx context.Context, request request.CreateAccount) error
	GetAccount(ctx context.Context, request request.GetAccount) (response.GetAccount, error)
}

type accounts struct {
	service accountsService
}

func NewAccounts(service accountsService) *accounts {
	return &accounts{service: service}
}

func (h *accounts) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.createAccount)
	r.Get("/{accountID}", h.getAccount)

	return r
}

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
