package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vladiq/user-balance-service/cmd/api/handlers"
	"github.com/vladiq/user-balance-service/internal/api/service"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"
	"github.com/vladiq/user-balance-service/internal/repository"
	"net/http"
	"time"

	"github.com/vladiq/user-balance-service/internal/pkg/logging"
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

	accountRepo := repository.NewAccountRepository(dbConn)
	reservationRepo := repository.NewReservationRepository(dbConn)
	//transactionRepo := repository.NewTransferRepository(dbConn)

	reservationsService := service.NewReservations(reservationRepo)
	accountsService := service.NewAccounts(accountRepo)

	reservationsHandler := handlers.NewReservations(reservationsService)
	accountsHandler := handlers.NewAccounts(accountsService)

	r := chi.NewRouter()

	r.Use(middleware.RedirectSlashes)
	r.Use(chilogger.LoggerMiddleware(&logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route(cfg.Server.BasePath, func(r chi.Router) {
		r.Mount("/reservations", reservationsHandler.Routes())
		r.Mount("/accounts", accountsHandler.Routes())
	})
	//r.Mount(cfg.Server.BasePath, reservationsHandler.Routes())
	//r.Mount()

	logger.Info().Msgf("Starting service at %s:%d%s", cfg.Server.Host, cfg.Server.Port, cfg.Server.BasePath)
	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}
