package logger

import (
	"log/slog"
	"os"
)

const (
	levelDebug string = "debug"
	levelInfo  string = "info"
	levelWarn  string = "warn"
	levelError string = "error"
)

func NewLogger(loggerLevel string) *slog.Logger {
	var log *slog.Logger

	switch loggerLevel {
	case levelDebug:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case levelInfo:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	case levelWarn:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelWarn},
			),
		)
	case levelError:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelError},
			),
		)
	}

	return log
}

func Error(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func String(key string, value string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.StringValue(value),
	}
}

func Int(key string, value int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.IntValue(value),
	}
}

func Any(key string, value any) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(value),
	}
}
