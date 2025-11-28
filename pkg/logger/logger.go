package logger

import (
	"log/slog"
	"os"
)

func Setup(env string) *slog.Logger {
	var handler slog.Handler

	if env == "local" {
		// Текстовый формат для удобства чтения глазами
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	} else {
		// JSON формат для серверов
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(handler)
}
