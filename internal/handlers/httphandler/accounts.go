package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h *handler) getBalance(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	account, err := h.service.GetAccount(r.Context(), userID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
}

func (h *handler) createAccount(w http.ResponseWriter, r *http.Request) {
	reqJSON := struct {
		Amount float64 `json:"amount"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	err = h.service.CreateAccount(r.Context(), reqJSON.Amount)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) updateBalance(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	reqJSON := struct {
		Amount float64 `json:"amount"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	err = h.service.UpdateBalance(r.Context(), userID, reqJSON.Amount)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
