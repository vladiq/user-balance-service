package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/cmd/api/handlers"
	"github.com/vladiq/user-balance-service/internal/api/service"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"
	"github.com/vladiq/user-balance-service/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes)
	router.Use(chilogger.LoggerMiddleware(&logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route(cfg.Server.BasePath, func(r chi.Router) {
		r.Mount("/reservations", reservationsHandler.Routes())
		r.Mount("/accounts", accountsHandler.Routes())
	})

	runServer(cfg, logger, router)
	//http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}

var serverShutdownTimeout = 5 * time.Second

func runServer(cfg config.Config, logger zerolog.Logger, router chi.Router) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Server listen error")
		}
	}()
	logger.Info().Msgf("Started service at %s:%d%s", cfg.Server.Host, cfg.Server.Port, cfg.Server.BasePath)

	<-gracefulShutdown
	logger.Info().Msg("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server shutdown failed")
	}
}
