package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vladiq/user-balance-service/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

var serverShutdownTimeout = 5 * time.Second

func RunServer(cfg config.Config, logger zerolog.Logger, router chi.Router) {
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
