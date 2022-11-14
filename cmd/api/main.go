package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vladiq/user-balance-service/cmd/api/handlers"
	_ "github.com/vladiq/user-balance-service/docs"
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
	httpSwagger "github.com/swaggo/http-swagger"
)

const configYML = "config.yml"

// @title       User Balance Microservice
// @version 1.0
// @description A microservice for user balance management, money transfer and report generation.
// @contact.name Vladislav Kosogorov
// @host     localhost:7000
// @BasePath /balance-service
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
	reportRepo := repository.NewReportRepository(dbConn)

	reservationsService := service.NewReservations(reservationRepo)
	accountsService := service.NewAccounts(accountRepo)
	transfersService := service.NewTransfers(transferRepo)
	reportsService := service.NewReports(reportRepo)

	reservationsHandler := handlers.NewReservations(reservationsService)
	accountsHandler := handlers.NewAccounts(accountsService)
	transfersHandler := handlers.NewTransfers(transfersService)
	reportsHandler := handlers.NewReports(reportsService)

	router := chi.NewRouter()
	router.Use(middleware.RedirectSlashes)
	router.Use(chilogger.LoggerMiddleware(&logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	swaggerURL := fmt.Sprintf("http://%s:%d%s/swagger/doc.json", cfg.Server.Host, cfg.Server.Port, cfg.Server.BasePath)
	router.Get(cfg.Server.BasePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))

	router.Route(cfg.Server.BasePath, func(router chi.Router) {
		router.Mount("/reservations", reservationsHandler.Routes())
		router.Mount("/accounts", accountsHandler.Routes())
		router.Mount("/transfers", transfersHandler.Routes())
		router.Mount("/reports", reportsHandler.Routes())
	})

	httpserver.RunServer(cfg, logger, router)
}
