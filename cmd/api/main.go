package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/vladiq/user-balance-service/internal/handlers"
	"github.com/vladiq/user-balance-service/internal/pkg/chilogger"
	"github.com/vladiq/user-balance-service/internal/pkg/logger"
	"github.com/vladiq/user-balance-service/internal/repository"
	"github.com/vladiq/user-balance-service/internal/service"
	"github.com/vladiq/user-balance-service/pkg/config"
	"github.com/vladiq/user-balance-service/pkg/postgres"
	"net/http"
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

	repo := repository.NewRepo(dbConn, l)
	balanceService := service.NewService(repo, l)

	r := chi.NewRouter()
	r.Use(chilogger.NewStructuredLogger(&l))

	balanceHandler := handlers.NewBalance(l, balanceService)
	r.Mount(cfg.Server.BasePath, balanceHandler.Routes())

	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}
