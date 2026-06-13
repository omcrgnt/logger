package logger

import (
	"fmt"
	"log/slog"
	"strings"
)

const (
	defaultLevelValue  = "info"
	defaultFormatValue = "json"
)

func mergeDefaults(level, format string) (string, string) {
	if level == "" {
		level = defaultLevelValue
	}
	if format == "" {
		format = defaultFormatValue
	}
	return level, format
}

func applyHandler(level, format string, out Output) error {
	lvl, err := parseLevel(level)
	if err != nil {
		return err
	}
	opts := &slog.HandlerOptions{Level: lvl}
	switch strings.ToLower(format) {
	case "json":
		slog.SetDefault(slog.New(slog.NewJSONHandler(out, opts)))
	case "text":
		slog.SetDefault(slog.New(slog.NewTextHandler(out, opts)))
	default:
		return fmt.Errorf("logger: invalid format %q", format)
	}
	return nil
}

func parseLevel(value string) (slog.Leveler, error) {
	switch strings.ToLower(value) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return nil, fmt.Errorf("logger: invalid level %q", value)
	}
}

func validateLevel(level string) error {
	_, err := parseLevel(level)
	return err
}

func validateFormat(format string) error {
	switch strings.ToLower(format) {
	case "json", "text":
		return nil
	default:
		return fmt.Errorf("logger: invalid format %q", format)
	}
}
