package handlers

import (
	"context"
	"github.com/go-chi/render"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"

	"github.com/go-chi/chi/v5"
)

type reportsService interface {
	GetServiceReport(ctx context.Context, request request.GetServiceReport) error
}

type reports struct {
	service reportsService
}

func NewReports(service reportsService) *reports {
	return &reports{service: service}
}

func (h *reports) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.getServiceReport)

	return r
}

func (h *reports) getServiceReport(w http.ResponseWriter, r *http.Request) {
	var req request.GetServiceReport

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	//////// csv
}
