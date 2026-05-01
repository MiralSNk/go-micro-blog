package logger

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

// New создает логгер для микросервиса, котрый принимает параметры: (по умолчанию `prod`)
//
// local -> для внутреннего тестирования на уровне Debug в формате.
// dev -> для тестового сервера на уровне Debug в формате JSON.
// prod -> для основного сервера на уровне Info в формате JSON.
func New(serviceName, level string) *slog.Logger {
	var log *slog.Logger

	switch level {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prod:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	log = log.With(slog.String("service", serviceName))
	return log
}
