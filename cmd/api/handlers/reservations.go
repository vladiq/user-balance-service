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
	CancelReservation(ctx context.Context, request request.CancelReservation) error
	ConfirmReservation(ctx context.Context, request request.ConfirmReservation) error
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
	r.Delete("/cancel-reservation/{reservationID}", h.CancelReservation)
	r.Delete("/confirm-reservation/{reservationID}", h.ConfirmReservation)

	return r
}

// CreateReservation creates a reservation for a user on a special reservation account
// @Summary Create a reservation
// @Tags Reservations
// @ID create-reservation
// @Accept json
// @Param input body request.CreateReservation true "Reservation info"
// @Success 201 "Created"
// @Failure 400 {string} constant.ErrBadRequest "Bad request"
// @Failure 500 {string} constant.ErrInternalServerError "Internal server error"
// @Router  /reservations [post]
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

	w.WriteHeader(http.StatusNoContent)
}

// CancelReservation cancels a reservation, deletes an entry from a table and returns money back to user
// @Summary Cancel and delete a reservation returning money back to user
// @Tags Reservations
// @ID cancel-reservation
// @Accept json
// @Param reservation_id path string true "Reservation ID"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router  /reservations/cancel-reservation/{reservation_id} [delete]
func (h *reservations) CancelReservation(w http.ResponseWriter, r *http.Request) {
	var req request.CancelReservation

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if err := h.service.CancelReservation(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ConfirmReservation confirms a reservation, deletes an entry from a table and adds an entry to reports table
// @Summary Confirm and delete a reservation adding a record to reports table
// @Tags Reservations
// @ID confirm-reservation
// @Accept json
// @Param reservation_id path string true "Reservation ID"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router  /reservations/confirm-reservation/{reservation_id} [delete]
func (h *reservations) ConfirmReservation(w http.ResponseWriter, r *http.Request) {
	var req request.ConfirmReservation

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if err := h.service.ConfirmReservation(r.Context(), req); err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
