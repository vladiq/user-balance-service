package main

import (
	"fmt"
	"net/http"

	"github.com/vladiq/user-balance-service/internal/handlers"
	"github.com/vladiq/user-balance-service/internal/pkg/logger"
	"github.com/vladiq/user-balance-service/internal/repository/account"
	"github.com/vladiq/user-balance-service/internal/repository/reservation"
	"github.com/vladiq/user-balance-service/internal/repository/transaction"
	"github.com/vladiq/user-balance-service/internal/service"
	"github.com/vladiq/user-balance-service/pkg/config"
	"github.com/vladiq/user-balance-service/pkg/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

const configYML = "config.yml"

func main() {
	if err := config.ReadConfigYML(configYML); err != nil {
		log.Fatal().Err(err).Msg("Failed config initialization")

	}
	cfg := config.GetConfigInstance()

	l, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting logger")
	}

	l.Info().
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	l.Trace().Str("DSN", cfg.DB.DSN).Msg("Connecting to database")
	dbConn, err := postgres.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer dbConn.Close()

	accountRepo := account.NewAccountRepository(l, dbConn)
	reservationRepo := reservation.NewReservationRepository(l, dbConn)
	transactionRepo := transaction.NewTransactionRepository(l, dbConn)

	balanceService := service.NewBalanceService(l, accountRepo, reservationRepo, transactionRepo)

	r := chi.NewRouter()
	balanceHandler := handlers.NewHandler(l, balanceService)
	r.Mount(cfg.Server.BasePath, balanceHandler.Routes())

	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}
