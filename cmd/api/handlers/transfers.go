package handlers

import (
	"context"
	"github.com/google/uuid"
	m "github.com/vladiq/user-balance-service/internal/pkg/middleware"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type transfersService interface {
	MakeTransfer(ctx context.Context, request request.MakeTransfer) error
	GetTransfers(ctx context.Context, request request.GetUserTransfers, pageID uuid.UUID) (response.GetUserTransfers, error)
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
	r.With(m.Pagination).Get("/reports/{accountID}", h.listUserTransfers)

	return r
}

// makeTransfer creates a money transfer between two user accounts and adds entries to transfers table
// @Summary Make a money transfer between two user accounts and add balance report entries
// @Tags    Transfers
// @ID transfer-create
// @Accept json
// @Param input body request.MakeTransfer true "transfer info"
// @Success 201 "Created"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router  /transfers [post]
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

// listUserTransfers gets user transfer list, supports pagination and sorting by date and amount
// @Summary Gets user transfer list, supports pagination and sorting by date and amount. Page size is 4.
// @Tags    Transfers
// @ID transfers-report
// @Produce json
// @Param id path string true "account uuid"
// @Param order-by query string false "order by 'date' or 'amount'"
// @Param page-id query string false "pagination uuid key. Don't set it at first and then the next page keys would appear in response body"
// @Success 200 {object} response.GetUserTransfers "next_page_id returned would be the next key for pagination"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /transfers/reports/{id} [get]
func (h *transfers) listUserTransfers(w http.ResponseWriter, r *http.Request) {
	var req request.GetUserTransfers

	if err := req.Bind(r); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	pageID := r.Context().Value(m.PageIDKey)

	userTransfers, err := h.service.GetTransfers(r.Context(), req, pageID.(uuid.UUID))
	if err != nil {
		utils.RenderError(w, r, err)
		return
	}

	render.JSON(w, r, userTransfers)
}
