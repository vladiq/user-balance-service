package main

import (
	"context"
	"time"

	"github.com/vladiq/user-balance-service/cmd/api/handlers"
	"github.com/vladiq/user-balance-service/internal/api/service"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"
	"github.com/vladiq/user-balance-service/internal/pkg/httpserver"
	"github.com/vladiq/user-balance-service/internal/pkg/logging"
	"github.com/vladiq/user-balance-service/internal/repository"
	"github.com/vladiq/user-balance-service/pkg/config"
	"github.com/vladiq/user-balance-service/pkg/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

const configYML = "config.yml"

var gatewayTimeout = 30 * time.Second

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
		log.Fatal().Err(err).Msg("error while connecting to database")
	}
	defer dbConn.Close()

	accountRepo := repository.NewAccountRepository(dbConn)
	reservationRepo := repository.NewReservationRepository(dbConn)
	transferRepo := repository.NewTransferRepository(dbConn)

	reservationsService := service.NewReservations(reservationRepo)
	accountsService := service.NewAccounts(accountRepo)
	transfersService := service.NewTransfers(transferRepo)

	reservationsHandler := handlers.NewReservations(reservationsService)
	accountsHandler := handlers.NewAccounts(accountsService)
	transfersHandler := handlers.NewTransfers(transfersService)

	router := chi.NewRouter()
	router.Use(middleware.RedirectSlashes)
	router.Use(chilogger.LoggerMiddleware(&logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(gatewayTimeout))

	router.Route(cfg.Server.BasePath, func(r chi.Router) {
		r.Mount("/reservations", reservationsHandler.Routes())
		r.Mount("/accounts", accountsHandler.Routes())
		r.Mount("/transfers", transfersHandler.Routes())
	})

	httpserver.RunServer(cfg, logger, router)
}
