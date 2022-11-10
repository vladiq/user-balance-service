package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service interface {
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
	MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error
	AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type handler struct {
	logger  zerolog.Logger
	service Service
}

func NewBalanceService(logger zerolog.Logger, service Service) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Recoverer)
	router.Use(chilogger.LoggerMiddleware(&h.logger))
	router.Use(middleware.Timeout(5 * time.Second))

	router.Route("/", func(r chi.Router) {
		router.Get("/{userID}", h.getBalance)
	})

	return router
}

func (h *handler) getBalance(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		h.logger.Fatal().Err(err).Msg("failed uuid parsing")
	}
	account, err := h.service.GetUserBalance(r.Context(), userID)

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		h.logger.Fatal().Err(err).Msg("encoding error")
	}
}
