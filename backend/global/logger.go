package global

import (
	"flowing/internal/config"
	"log/slog"
	"os"
	"path"
)

func InitLogger(conf *config.Config) {
	writer := os.Stdout
	if conf.Logger.Path != "" {
		_ = os.MkdirAll(path.Dir(conf.Logger.Path), os.ModePerm)
		file, err := os.OpenFile(conf.Logger.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		writer = file
	}
	level := slog.LevelInfo
	switch conf.Logger.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}
	var handler slog.Handler
	switch conf.Logger.Format {
	case "json":
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})
	case "console":
		handler = slog.NewTextHandler(writer, &slog.HandlerOptions{Level: level})
	default:
		panic("invalid logger format")
	}
	slog.SetDefault(slog.New(handler))
}
