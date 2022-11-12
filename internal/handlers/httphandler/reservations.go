package httphandler

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

func (h *handler) createReservation(w http.ResponseWriter, r *http.Request) {
	reqJSON := struct {
		UserID    string  `json:"user_id"`
		ServiceID string  `json:"service_id"`
		OrderID   string  `json:"order_id"`
		Amount    float64 `json:"amount"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	userID, err := uuid.Parse(reqJSON.UserID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	serviceID, err := uuid.Parse(reqJSON.ServiceID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	orderID, err := uuid.Parse(reqJSON.OrderID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	err = h.service.CreateReservation(
		r.Context(),
		userID,
		serviceID,
		orderID,
		reqJSON.Amount,
	)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
