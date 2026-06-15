package logger

import (
	"context"
	"io"
)

// Logger is the logging API.
type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}

// Output is a sealed log sink port for SDI.
type Output interface {
	io.Writer
	markerLoggerOutput()
}
