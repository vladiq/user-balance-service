package httphandler

import (
	"context"
	"time"

	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service interface {
	GetUser(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
	CreateAccount(ctx context.Context, amount float64) error
	UpdateBalance(ctx context.Context, userID uuid.UUID, amount float64) error
	CreateReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
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

	r.Use(middleware.RedirectSlashes)
	r.Use(chilogger.LoggerMiddleware(&h.logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.createAccount)        // создание аккаунта с данной суммой (id пользователя, сумма)
			r.Get("/{userID}", h.getBalance)    // получение баланса пользователя (id пользователя)
			r.Put("/{userID}", h.updateBalance) // зачисление и списание средств (id пользователя, сколько средств зачислить)
		})

		//r.Route("/transactions", func(r chi.Router) {
		//	r.Post("/", h.createTransaction) // перевод от пользователя к пользователю ()
		//})

		r.Route("/reservations", func(r chi.Router) {
			r.Post("/", h.createReservation) // резервирование на отдельном счёте (id пользователя, id услуги, id заказа, сумма)
			//r.Put("/{reservationID}", h.endReservation) // признание выручки - списать из резерва деньги, добавить запись в отчёт для бухгалтерии (id пользователя, id услуги, id заказа, сумма)
		})

		//r.Route("/reports", func(r chi.Router) {
		//	r.Get("/service-report", h.serviceReport)                // предоставить отчёт для всех пользователей (вход: год, месяц)
		//	r.Get("/user-report/{userID}", r.userTransactionsReport) // получение списка транзакций пользователя
		//})

		// r.Get("/docs") // swagger documentation
	})

	return r
}
