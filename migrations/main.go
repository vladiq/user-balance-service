package main

import (
	"context"
	"embed"
	"github.com/vladiq/user-balance-service/pkg/config"
	"github.com/vladiq/user-balance-service/pkg/postgres"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var migrationCmd = "up"

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := postgres.New(context.Background(), &cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("Database connection error")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	err = goose.Run(migrationCmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot run goose")
	}
}
