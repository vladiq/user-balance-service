package handlers

import (
	"context"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocarina/gocsv"
)

type transfersService interface {
	MakeTransfer(ctx context.Context, request request.MakeTransfer) error
	UserMonthlyReport(ctx context.Context, request request.UserMonthlyReport) ([]*response.GetUserMonthlyReport, error)
}

type transfers struct {
	service transfersService
}

func NewTransfers(service transfersService) *transfers {
	return &transfers{service: service}
}

func (h *transfers) Routes() *chi.Mux {
	r := chi.NewRouter()

	// /report/dfjfk-2343223-edd324?year=2022&month=03
	r.Get("/reports/{accountID}", h.getUserMonthlyReport)
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

func (h *transfers) getUserMonthlyReport(w http.ResponseWriter, r *http.Request) {
	var req request.UserMonthlyReport

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	reportEntries, err := h.service.UserMonthlyReport(r.Context(), req)
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
