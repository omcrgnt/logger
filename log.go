package logger

import (
	"context"
	"log/slog"
)

// Debug logs at debug level.
func Debug(ctx context.Context, msg string, args ...any) {
	slog.Default().DebugContext(ctx, msg, args...)
}

// Info logs at info level.
func Info(ctx context.Context, msg string, args ...any) {
	slog.Default().InfoContext(ctx, msg, args...)
}

// Warn logs at warn level.
func Warn(ctx context.Context, msg string, args ...any) {
	slog.Default().WarnContext(ctx, msg, args...)
}

// Error logs at error level.
func Error(ctx context.Context, msg string, args ...any) {
	slog.Default().ErrorContext(ctx, msg, args...)
}
