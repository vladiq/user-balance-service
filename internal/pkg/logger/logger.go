package logger

import (
	"errors"
	"os"

	"github.com/vladiq/user-balance-service/pkg/config"

	"github.com/rs/zerolog"
)

func NewLogger(cfg config.Config) (zerolog.Logger, error) {
	var level zerolog.Level

	switch cfg.Logger.Level {
	case "panic":
		level = zerolog.PanicLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "error":
		level = zerolog.ErrorLevel
	case "warn":
		level = zerolog.WarnLevel
	case "info":
		level = zerolog.InfoLevel
	case "debug":
		level = zerolog.DebugLevel
	case "trace":
		level = zerolog.TraceLevel
	default:
		return zerolog.Logger{}, errors.New("wrong logger level")
	}

	logger := zerolog.Logger{}.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Timestamp().Logger().
		Level(level)

	return logger, nil
}
