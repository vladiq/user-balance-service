package main

import (
	"context"
	"fmt"
	http2 "github.com/vladiq/user-balance-service/internal/handlers/httphandler"
	"net/http"
	"time"

	"github.com/vladiq/user-balance-service/internal/pkg/logging"
	"github.com/vladiq/user-balance-service/internal/repository/account"
	"github.com/vladiq/user-balance-service/internal/repository/reservation"
	"github.com/vladiq/user-balance-service/internal/repository/transfer"
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

	logger, err := logging.NewLogger(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting logger")
	}

	logger.Info().
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	logger.Trace().Str("DSN", cfg.DB.DSN).Msg("Connecting to database")
	initCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbConn, err := postgres.New(initCtx, &cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer dbConn.Close()

	accountRepo := account.NewAccountRepository(logger, dbConn)
	reservationRepo := reservation.NewReservationRepository(logger, dbConn)
	transactionRepo := transfer.NewTransferRepository(logger, dbConn)

	balanceService := service.NewBalanceService(logger, accountRepo, reservationRepo, transactionRepo)

	r := chi.NewRouter()
	balanceHandler := http2.NewHandler(logger, balanceService)
	r.Mount(cfg.Server.BasePath, balanceHandler.Routes())

	logger.Info().Msgf("Starting service at %s:%d%s", cfg.Server.Host, cfg.Server.Port, cfg.Server.BasePath)
	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}
