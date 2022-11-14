package handlers

import (
	"context"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocarina/gocsv"
)

type reportsService interface {
	GetServiceReport(ctx context.Context, request request.GetServiceReport) ([]*response.GetServiceReport, error)
}

type reports struct {
	service reportsService
}

func NewReports(service reportsService) *reports {
	return &reports{service: service}
}

func (h *reports) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.servicesMonthReport)

	return r
}

func (h *reports) servicesMonthReport(w http.ResponseWriter, r *http.Request) {
	var req request.GetServiceReport

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	reportEntries, err := h.service.GetServiceReport(r.Context(), req)
	if err != nil {
		utils.RenderError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Add("Content-Disposition", `attachment; filename="user-report.csv"`)
	if err := gocsv.Marshal(reportEntries, w); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}
}
