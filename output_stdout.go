package logger

import "os"

// OutputStdout writes logs to os.Stdout.
type OutputStdout struct{}

func (OutputStdout) Write(p []byte) (int, error) { return os.Stdout.Write(p) }

func (OutputStdout) markerLoggerOutput() {}
