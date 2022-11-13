package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/utils"
	"net/http"
)

type transfersService interface {
	MakeTransfer(ctx context.Context, request request.MakeTransfer) error
}

type transfers struct {
	service transfersService
}

func NewTransfers(service transfersService) *transfers {
	return &transfers{service: service}
}

func (h *transfers) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.makeTransfer)

	return r
}

func (h *transfers) makeTransfer(w http.ResponseWriter, r *http.Request) {
	var req request.MakeTransfer

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if err := h.service.MakeTransfer(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
