package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"
	"net/http"
	"time"
)

type BalanceService interface {
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
}

type Balance struct {
	logger  zerolog.Logger
	service BalanceService
}

func NewBalance(logger zerolog.Logger, service BalanceService) *Balance {
	return &Balance{
		logger:  logger,
		service: service,
	}
}

func (b *Balance) Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Recoverer)
	router.Use(chilogger.LoggerMiddleware(&b.logger))
	router.Use(middleware.Timeout(5 * time.Second))

	router.Route("/", func(r chi.Router) {
		router.Get("/{userID}", b.getBalance)
	})

	return router
}

func (b *Balance) getBalance(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		b.logger.Fatal().Err(err).Msg("failed uuid parsing")
	}
	account, err := b.service.GetUserBalance(r.Context(), userID)

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		b.logger.Fatal().Err(err).Msg("encoding error")
	}
}

//func (b *Balance) addFunds(w http.ResponseWriter, r *http.Request) {
//	userIDParam := chi.URLParam(r, "userID")
//	userID, err := uuid.Parse(userIDParam)
//	if err != nil {
//		log.Fatal().Err(err).Msg("TOO BAD!!!!!!!")
//	}
//	err = b.service.AddFundsToAccount(r.Context(), userID, 2134.23)
//	if err != nil {
//		log.Fatal().Err(err).Msg("TOO BAD OMG!!!!!!!")
//	}
//}
