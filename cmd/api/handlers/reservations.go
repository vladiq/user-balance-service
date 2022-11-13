package handlers

import (
	"context"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type reservationsService interface {
	CreateReservation(ctx context.Context, request request.CreateReservation) error
}

type reservations struct {
	service reservationsService
}

func NewReservations(service reservationsService) *reservations {
	return &reservations{service: service}
}

func (h *reservations) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateReservation)

	return r
}

func (h *reservations) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var req request.CreateReservation

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())

		return
	}

	if err := h.service.CreateReservation(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)

		return
	}
}
