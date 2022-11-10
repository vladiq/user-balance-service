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
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type handler struct {
	logger  zerolog.Logger
	service Service
}

func NewHandler(logger zerolog.Logger, service Service) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.RedirectSlashes)
	r.Use(chilogger.LoggerMiddleware(&h.logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{userID}", h.getBalance) // получение баланса пользователя (id пользователя)
			//r.Post("/", h.createUser)           // создание аккаунта с данной суммой (id пользователя, сумма)
			//r.Put("/{userID}", h.updateBalance) // зачисление и списание средств (id пользователя, сколько средств зачислить)
		})

		//r.Route("/transactions", func(r chi.Router) {
		//	r.Post("/", h.createTransaction) // перевод от пользователя к пользователю ()
		//})
		//
		//r.Route("/reservations", func(r chi.Router) {
		//	r.Post("/", h.createReservation)            // резервирование на отдельном счёте (id пользователя, id услуги, id заказа, сумма)
		//	r.Put("/{reservationID}", h.endReservation) // признание выручки - списать из резерва деньги, добавить запись в отчёт для бухгалтерии (id пользователя, id услуги, id заказа, сумма)
		//})
		//
		//r.Route("/reports", func(r chi.Router) {
		//	r.Get("/service-report", h.serviceReport)    // предоставить отчёт для всех пользователей (вход: год, месяц)
		//	r.Get("/user-report/{userID}", r.userReport) // получение списка транзакций пользователя
		//})
	})

	return r
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
